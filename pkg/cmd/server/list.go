package server

import (
	"github.com/spf13/cobra"

	"github.com/konstellation-io/kli/cmd/factory"
	"github.com/konstellation-io/kli/internal/render"
)

// NewListCmd creates a new command to list servers existing in config file.
func NewListCmd(f factory.CmdFactory) *cobra.Command {
	log := f.Logger()
	cmd := &cobra.Command{
		Use:     "ls",
		Aliases: []string{"list"},
		Short:   "List all available servers",
		Run: func(cmd *cobra.Command, args []string) {
			cfg := f.Config()

			if len(cfg.ServerList) == 0 {
				log.Info("No servers found.")
				return
			}

			r := render.DefaultRenderer(cmd.OutOrStdout())
			cfg.RenderServerList(r)
		},
	}

	return cmd
}
