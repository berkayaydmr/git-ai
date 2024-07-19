package subs

import (
	"fmt"

	"github.com/berkayaydmr/git-ai/pkg/errors"
	"github.com/berkayaydmr/git-ai/pkg/storage"
	"github.com/berkayaydmr/git-ai/utils"
	"github.com/urfave/cli/v2"
)

func SubProfileCommand(storage storage.StorageInterface) *cli.Command {
	return &cli.Command{
		Name:   "profile",
		Action: newProfileAction(storage),
	}
}

func newProfileAction(storage storage.StorageInterface) cli.ActionFunc {
	return func(ctx *cli.Context) error {
		name := ctx.Args().First()
		if name == "" {
			return errors.ErrNameNotProvided
		}

		apiKeyModel, err := utils.GetGptProfileFromUser(name)
		if err != nil {
			return err
		}

		err = storage.NewProfile(apiKeyModel)
		if err != nil {
			return err
		}

		fmt.Println("Api key created successfully by the name of", name)

		return nil
	}
}
