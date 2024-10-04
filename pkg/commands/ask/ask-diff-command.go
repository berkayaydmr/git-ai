package ask

import (
	"context"
	"fmt"
	"os"

	"github.com/berkayaydmr/git-ai/pkg/clients/gpt"
	"github.com/berkayaydmr/git-ai/pkg/constants"
	"github.com/berkayaydmr/git-ai/pkg/errors"
	"github.com/berkayaydmr/git-ai/pkg/messages"
	"github.com/berkayaydmr/git-ai/pkg/storage"
	"github.com/berkayaydmr/git-ai/pkg/storage/models"
	"github.com/berkayaydmr/git-ai/utils"

	"github.com/urfave/cli/v2"
)

func DiffCommand(storage storage.StorageInterface, gptClient *gpt.Client) *cli.Command {
	return &cli.Command{
		Name:   "diff",
		Action: newAskDiffAction(storage, gptClient),
	}
}

func newAskDiffAction(storage storage.StorageInterface, gptClient *gpt.Client) cli.ActionFunc {
	return func(ctx *cli.Context) error {
		repositoryUrl := ctx.Args().First()
		if repositoryUrl == "" {
			return errors.ErrRepositoryUrlNotProvided
		}

		currDir, err := os.Getwd()
		if err != nil {
			return err
		}

		// git ai diff main init
		// git ai diff path/to/repo main init

		var branch1, branch2 string
		if utils.CheckBranchExist(currDir, repositoryUrl) {
			branch1 = repositoryUrl
			repositoryUrl = currDir
			if !utils.CheckBranchExist(currDir, branch1) {
				return errors.ErrFirstBranchDoesNotExist
			}

			branch2 = ctx.Args().Get(1)
			if !utils.CheckBranchExist(currDir, branch2) {
				return errors.ErrSecondBranchDoesNotExist
			}
		} else {
			branch1 = ctx.Args().Get(1)
			if !utils.CheckBranchExist(repositoryUrl, branch1) {
				return errors.ErrFirstBranchDoesNotExist
			}

			branch2 = ctx.Args().Get(2)
			if !utils.CheckBranchExist(repositoryUrl, branch2) {
				return errors.ErrSecondBranchDoesNotExist
			}
		}

		fmt.Println("Repository URL: ", repositoryUrl)

		if branch1 == branch2 {
			utils.TypeWriterEffect(messages.SameBranchDiffMessage, utils.Normal)
			return nil
		}

		diff, err := utils.BranchDiff(repositoryUrl, branch1, branch2)
		if err != nil {
			return err
		}

		if diff == "" {
			utils.TypeWriterEffect(messages.NoDiffMessage, utils.Normal)
			return nil
		}

		layoutFile, err := os.ReadFile(constants.ReviewLayoutFileLocation)
		if err != nil {
			return err
		}

		apiKeys, err := storage.GetProfiles()
		if err != nil {
			return err
		}

		selectedApiKey, err := selectApiKeyModel(apiKeys)
		if err != nil {
			return err
		}

		fmt.Println("Selected API key: ", selectedApiKey.Name)

		gptContext, cancel := context.WithTimeout(context.Background(), gpt.Timeout)
		defer cancel()
		choies, err := gptClient.Ask(gptContext, selectedApiKey, gpt.MakeModelSoftwareEngineerCharacter, fmt.Sprintf(gpt.RepositoryAndBranchNames, repositoryUrl, branch1, branch2), gpt.AskDiffQuestion, diff, string(layoutFile))
		if err != nil {
			return err
		}

		fmt.Println("\n")
		utils.TypeWriterEffect(utils.MakeChoicesString(choies), utils.Light)

		return nil
	}
}

const temporaryApiKeyName = "tmp"

func selectApiKeyModel(apiKeys []models.Profile) (models.Profile, error) {
	if len(apiKeys) == 1 {
		fmt.Printf(messages.OnlyProfileFound, apiKeys[0].Name)
		return apiKeys[0], nil
	}

	for i, apiKey := range apiKeys {
		utils.TypeWriterEffect(fmt.Sprintf("%d. %s\n", i+1, apiKey.Name), utils.Faster)
	}

	if len(apiKeys) == 0 {
		utils.TypeWriterEffect(messages.CustomProfileOrAdd, utils.Faster)
	}

	fmt.Print("Select an API key: ")
	var selected int
	fmt.Scanln(&selected)

	if selected == 0 {
		fmt.Println()
		apiKey, err := utils.GetGptProfileFromUser(temporaryApiKeyName)
		if err != nil {
			return models.Profile{}, err
		}

		return apiKey, nil
	}

	return apiKeys[selected-1], nil
}
