package cmd

import (
	"os"
	"os/exec"

	"github.com/chrisgavin/gh-sops/internal/paths"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

type DecryptCommand struct {
	*RootCommand
}

func registerDecryptCommand(rootCommand *RootCommand) {
	command := &DecryptCommand{
		RootCommand: rootCommand,
	}
	cobraCommand := &cobra.Command{
		Use:           "decrypt",
		Short:         "Decrypt all files in the current directory.",
		SilenceErrors: true,
		SilenceUsage:  true,
		Args:          cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			encryptedPaths, err := paths.FindEncryptedFiles(".")
			if err != nil {
				return errors.Wrap(err, "Error finding encrypted files.")
			}

			for _, encryptedPath := range encryptedPaths {
				decryptedPath := paths.GetDecryptedPath(encryptedPath)
				sopsCommand := exec.CommandContext(cmd.Context(), "sops", "decrypt", encryptedPath, "--output", decryptedPath)
				sopsCommand.Stdout = os.Stdout
				sopsCommand.Stderr = os.Stderr
				err := sopsCommand.Run()
				if err != nil {
					return errors.Wrapf(err, "Error decrypting file %s.", encryptedPath)
				}
			}

			return nil
		},
	}
	command.root.AddCommand(cobraCommand)
}
