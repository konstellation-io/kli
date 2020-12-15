package version

import (
	"github.com/konstellation-io/kli/cmdutil"
	"github.com/spf13/cobra"
)

// NewStopCmd stop an existing version.
func NewStopCmd(f cmdutil.CmdFactory) *cobra.Command {
	log := f.Logger()
	cmd := &cobra.Command{
		Use:     "stop",
		Aliases: []string{"stop"},
		Args:    cmdutil.ComposeArgsCheck(cmdutil.CheckServerFlag, cobra.ExactArgs(1)),
		Short:   "Stop a version",
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
			err = c.Version().Stop(versionID, comment)
			if err != nil {
				return err
			}

			log.Success("Stopping version '%s'.", versionID)
			return nil
		},
	}

	cmd.Flags().StringP("message", "m", "", "Adds audit message to stop")
	_ = cmd.MarkPersistentFlagRequired("message")

	return cmd
}
