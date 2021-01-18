package krt

import (
	"github.com/spf13/cobra"

	"github.com/konstellation-io/kli/cmd/factory"
)

// NewCreateCmd command to create a KRT file from a directory.
func NewCreateCmd(f factory.CmdFactory) *cobra.Command {
	log := f.Logger()
	cmd := &cobra.Command{
		Use:   "create",
		Args:  cobra.ExactArgs(2),
		Short: "Create KRT file from a directory",
		RunE: func(cmd *cobra.Command, args []string) error {
			src := args[0]
			target := args[1]

			krt := f.Krt()
			err := krt.Build(src, target)
			if err != nil {
				return err
			}

			log.Success("New KRT file created.")

			return nil
		},
	}

	return cmd
}
