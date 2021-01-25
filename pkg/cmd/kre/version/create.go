package version

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/konstellation-io/kli/cmd/factory"
	"github.com/konstellation-io/kli/pkg/cmd/args"
)

// NewCreateCmd upload a KRT file to runtime and make a new version.
func NewCreateCmd(f factory.CmdFactory) *cobra.Command {
	log := f.Logger()
	cmd := &cobra.Command{
		Use:   "create [krt-file]",
		Args:  args.ComposeArgsCheck(args.CheckServerFlag, cobra.ExactArgs(1)),
		Short: "Upload a KRT and create a new version",
		RunE: func(cmd *cobra.Command, args []string) error {
			s, _ := cmd.Flags().GetString("server")
			c, err := f.KreClient(s)
			if err != nil {
				return err
			}

			krt := args[0]

			versionName, err := c.Version().Create(krt)
			if err != nil {
				return err
			}

			log.Success(fmt.Sprintf("Upload KRT completed, version %s created.", versionName))
			return nil
		},
	}

	return cmd
}
