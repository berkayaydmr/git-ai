package set

import (
	"github.com/berkayaydmr/git-ai/pkg/commands/set/subs"
	"github.com/berkayaydmr/git-ai/pkg/storage"
	"github.com/urfave/cli/v2"
)

func Command(storage storage.StorageInterface) *cli.Command {
	return &cli.Command{
		Name:    "set",
		Aliases: []string{"s"},
		Usage:   "Set and save such as api key etc.",
		Subcommands: []*cli.Command{
			subs.SubProfileCommand(storage),
		},
	}
}
