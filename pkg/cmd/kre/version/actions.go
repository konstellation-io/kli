package version

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/konstellation-io/kli/cmd/factory"
	"github.com/konstellation-io/kli/pkg/cmd/args"
	"github.com/konstellation-io/kli/pkg/errors"
)

// NewActionCmd manages all the actions of the version command.
func NewActionCmd(f factory.CmdFactory, action string) *cobra.Command {
	log := f.Logger()
	cmd := &cobra.Command{
		Use:   action,
		Args:  args.ComposeArgsCheck(args.CheckServerFlag, cobra.ExactArgs(1)),
		Short: fmt.Sprintf("%s a version", strings.Title(action)),
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
			version := c.Version()
			actionResult := ""
			switch action {
			case "start":
				err = version.Start(versionID, comment)
				actionResult = "Starting"
			case "stop":
				err = version.Stop(versionID, comment)
				actionResult = "Stopping"
			case "publish":
				err = version.Publish(versionID, comment)
				actionResult = "Publishing"
			case "unpublish":
				err = version.Unpublish(versionID, comment)
				actionResult = "Unpublishing"
			default:
				return errors.ErrUnknownVersionAction
			}
			if err != nil {
				return err
			}

			log.Success(fmt.Sprintf("%s version '%s'.", actionResult, versionID))
			return nil
		},
	}

	cmd.Flags().StringP("message", "m", "", fmt.Sprintf("Adds audit message to %s", action))
	_ = cmd.MarkFlagRequired("message")

	return cmd
}
