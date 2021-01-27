package krt

import (
	"github.com/spf13/cobra"

	"github.com/konstellation-io/kli/cmd/factory"
)

// NewCreateCmd command to create a KRT file from a directory.
func NewCreateCmd(f factory.CmdFactory) *cobra.Command {
	log := f.Logger()
	cmd := &cobra.Command{
		Use:   "create [source-folder] [dest-file.krt]",
		Args:  cobra.ExactArgs(2), //nolint:gomnd
		Short: "Create KRT file from a directory",
		RunE: func(cmd *cobra.Command, args []string) error {
			version, _ := cmd.Flags().GetString("version")
			src := args[0]
			target := args[1]

			krt := f.Krt()
			err := krt.Build(src, target, version)
			if err != nil {
				return err
			}

			log.Success("New KRT file created.")

			return nil
		},
	}
	cmd.Flags().StringP("version", "v", "", "KRT version name")

	return cmd
}
