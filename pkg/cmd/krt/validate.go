package krt

import (
	"github.com/spf13/cobra"

	"github.com/konstellation-io/kli/cmd/factory"
)

// NewValidateCmd validates a KRT yaml file.
func NewValidateCmd(f factory.CmdFactory) *cobra.Command {
	log := f.Logger()
	cmd := &cobra.Command{
		Use:   "validate",
		Args:  cobra.ExactArgs(1),
		Short: "Validate a KRT yml file",
		RunE: func(cmd *cobra.Command, args []string) error {
			krtFile := args[0]

			krt := f.Krt()
			err := krt.Validate(krtFile)
			if err != nil {
				return err
			}

			log.Success("Krt file is valid.")

			return nil
		},
	}

	return cmd
}
