package version

import (
	"github.com/konstellation-io/kli/cmdutil"
	"github.com/spf13/cobra"
)

// NewStartCmd starts an existing version.
func NewStartCmd(f cmdutil.CmdFactory) *cobra.Command {
	log := f.Logger()
	cmd := &cobra.Command{
		Use:     "start",
		Aliases: []string{"start"},
		Args:    cmdutil.ComposeArgsCheck(cmdutil.CheckServerFlag, cobra.ExactArgs(1)),
		Short:   "Start a version",
		RunE: func(cmd *cobra.Command, args []string) error {
			s, _ := cmd.Flags().GetString("server")
			c, err := f.KreClient(s)
			if err != nil {
				return err
			}

			comment, err := cmd.Flags().GetString("message")
			if err != nil {
				return err
			}

			versionID := args[0]
			err = c.Version().Start(versionID, comment)
			if err != nil {
				return err
			}

			log.Success("Starting version '%s'.", versionID)
			return nil
		},
	}

	cmd.Flags().StringP("message", "m", "", "Adds audit message to start")
	_ = cmd.MarkPersistentFlagRequired("message")

	return cmd
}
