package version

import (
	"fmt"
	"strings"

	"github.com/konstellation-io/kli/api/kre/version"
	"github.com/konstellation-io/kli/cmdutil"
	"github.com/konstellation-io/kli/internal/render"
	"github.com/spf13/cobra"
)

func NewConfigCmd(f cmdutil.CmdFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Args:  cmdutil.ComposeArgsCheck(cmdutil.CheckServerFlag, cobra.ExactArgs(1)),
		Short: "Get or set config values",
		RunE: func(cmd *cobra.Command, args []string) error {
			serverName, _ := cmd.Flags().GetString("server")

			versionID := args[0]
			vars, err := cmd.Flags().GetStringSlice("set")
			if err != nil {
				return err
			}

			if len(vars) > 0 {
				return updateConfig(f, serverName, versionID, vars)
			}

			return getConfig(f, cmd, serverName, versionID)

		},
	}
	cmd.Flags().StringSlice("set", []string{}, "Set new key value pair key=value")

	return cmd
}

func getConfig(f cmdutil.CmdFactory, cmd *cobra.Command, serverName, versionID string) error {
	log := f.Logger()
	c, err := f.KreClient(serverName)
	if err != nil {
		return err
	}

	config, err := c.Version().GetConfig(versionID)
	if err != nil {
		return err
	}

	if len(config.Vars) == 0 {
		log.Info("No config found.")
		return nil
	}

	r := render.DefaultRenderer(cmd.OutOrStdout())
	renderVariables(r, config)
	_, _ = fmt.Fprintln(cmd.OutOrStdout())
	if config.Completed {
		log.Success("Version config complete")
	} else {
		log.Warning("Version config incomplete")
	}
	return nil
}

func updateConfig(f cmdutil.CmdFactory, serverName, versionID string, vars []string) error {
	log := f.Logger()
	c, err := f.KreClient(serverName)
	if err != nil {
		return err
	}

	config := []version.ConfigVariableInput{}

	for _, v := range vars {
		arr := strings.Split(v, "=")
		config = append(config, version.ConfigVariableInput{
			Key:   arr[0],
			Value: arr[1],
		})
	}

	completed, err := c.Version().UpdateConfig(versionID, config)
	if err != nil {
		return err
	}

	status := "updated"
	if completed {
		status = "completed"
	}

	log.Success(fmt.Sprintf("Config %s for version '%s'.", status, versionID))

	return nil
}

func renderVariables(r render.Renderer, config *version.Config) {
	r.SetHeader([]string{
		"",
		"TYPE",
		"KEY",
		"VALUE",
	})

	for i, v := range config.Vars {
		r.Append([]string{
			fmt.Sprint(i + 1),
			string(v.Type),
			v.Key,
			v.Value,
		})
	}

	r.Render()
}
