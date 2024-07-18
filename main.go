package main

import (
	"fmt"
	"os"

	"github.com/berkayaydmr/git-ai/pkg/clients/gpt"
	"github.com/berkayaydmr/git-ai/pkg/commands/ask"
	"github.com/berkayaydmr/git-ai/pkg/commands/remove"
	"github.com/berkayaydmr/git-ai/pkg/commands/set"
	constant "github.com/berkayaydmr/git-ai/pkg/constants"
	"github.com/berkayaydmr/git-ai/pkg/cryptographer"
	"github.com/berkayaydmr/git-ai/pkg/storage"

	"github.com/urfave/cli/v2"
)

func main() {
	gptClient := gpt.New()

	config, err := cryptographer.ParseConfig()
	if err != nil {
		fmt.Println("\n", err, "\n")
		os.Exit(1)
	}

	parameter := cryptographer.Parameter{Config: config}
	cipher := cryptographer.New(parameter)

	storage := storage.New(cipher.Cryptographer, constant.EncryptedSourceFile)

	app := &cli.App{
		Name:  "git-ai",
		Usage: "A CLI tool for comparing git repository branches with GPT",
		Commands: []*cli.Command{
			ask.DiffCommand(storage, gptClient),
			set.Command(storage),
			remove.Command(storage),
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println("\n", err, "\n")
		os.Exit(1)
	}
}
