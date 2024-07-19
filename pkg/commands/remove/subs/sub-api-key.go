package subs

import (
	"fmt"

	"github.com/berkayaydmr/git-ai/pkg/errors"
	"github.com/berkayaydmr/git-ai/pkg/storage"
	"github.com/urfave/cli/v2"
)

func SubRemoveProfileCommand(storage storage.StorageInterface) *cli.Command {
	return &cli.Command{
		Name:   "key",
		Action: newRemoveProfileAction(storage),
	}
}

func newRemoveProfileAction(storage storage.StorageInterface) cli.ActionFunc {
	return func(ctx *cli.Context) error {
		name := ctx.Args().First()
		if name == "" {
			return errors.ErrNameNotProvided
		}

		err := storage.RemoveProfile(name)
		if err != nil {
			return err
		}

		fmt.Println("Api key removed successfully by the name of", name)

		return nil
	}
}
