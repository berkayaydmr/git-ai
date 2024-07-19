package remove

import (
	"github.com/berkayaydmr/git-ai/pkg/commands/remove/subs"
	"github.com/berkayaydmr/git-ai/pkg/storage"
	"github.com/urfave/cli/v2"
)

func Command(storage storage.StorageInterface) *cli.Command {
	return &cli.Command{
		Name:    "remove",
		Aliases: []string{"rm"},
		Usage:   "Remove something from the storage",
		Subcommands: []*cli.Command{
			subs.SubRemoveProfileCommand(storage),
		},
	}
}
