package list

import (
	"github.com/spf13/cobra"

	"github.com/konstellation-io/kli/cmdutil"
	"github.com/konstellation-io/kli/internal/render"
)

func NewListCmd(f *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "ls",
		Aliases: []string{"list"},
		Short:   "List all available servers",
		Run: func(cmd *cobra.Command, args []string) {
			cfg := f.Config()

			r := render.DefaultRenderer(cmd.OutOrStdout())
			cfg.RenderServerList(r)
		},
	}

	return cmd
}
