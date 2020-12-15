package publish

import (
	"github.com/konstellation-io/kli/cmdutil"
	"github.com/spf13/cobra"
)

// NewPublishCmd publishes an existing version.
func NewPublishCmd(f cmdutil.CmdFactory) *cobra.Command {
	log := f.Logger()
	cmd := &cobra.Command{
		Use:   "publish",
		Args:  cmdutil.ComposeArgsCheck(cmdutil.CheckServerFlag, cobra.ExactArgs(1)),
		Short: "Publish a version",
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
			err = c.Version().Publish(versionID, comment)
			if err != nil {
				return err
			}

			log.Success("Publishing version '%s'.", versionID)
			return nil
		},
	}

	cmd.Flags().StringP("message", "m", "", "Adds audit message to publish")
	_ = cmd.MarkPersistentFlagRequired("message")

	return cmd
}
