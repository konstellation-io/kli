package add

import (
	"github.com/guumaster/cligger"
	"github.com/spf13/cobra"

	"github.com/konstellation-io/kli/cmdutil"
	"github.com/konstellation-io/kli/internal/config"
	"github.com/konstellation-io/kli/internal/render"
)

func NewAddCmd(f cmdutil.CmdFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "add [name] [url] [token]",
		Aliases: []string{"set"},
		Args:    cobra.ExactArgs(3), //nolint:gomnd
		Short:   "Add a new server config file",
		Example: `
    $ kli server add my-server http://api.local.kre 12345abc
		`,
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg := f.Config()

			newServer := config.ServerConfig{
				Name:     args[0],
				URL:      args[1],
				APIToken: args[2],
			}

			err := cfg.AddServer(newServer)
			if err != nil {
				return err
			}

			r := render.DefaultRenderer(cmd.OutOrStdout())
			cfg.RenderServerList(r)

			cligger.Success("Server '%s' added.", newServer.Name)
			return nil
		},
	}

	return cmd
}
