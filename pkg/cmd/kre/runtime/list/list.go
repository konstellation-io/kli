package list

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/konstellation-io/kli/api"
	"github.com/konstellation-io/kli/cmdutil"
	"github.com/konstellation-io/kli/internal/errors"
	"github.com/konstellation-io/kli/internal/render"
)

func NewListCmd(f cmdutil.CmdFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "ls",
		Aliases: []string{"list"},
		Short:   "List all available runtimes",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg := f.Config()
			s, _ := cmd.Flags().GetString("server")

			if s == "" {
				s = cfg.DefaultServer
			}

			server := cfg.GetByServerName(s)
			if server == nil {
				return errors.ErrUnknownServerName
			}

			c, err := f.ServerClient(server.Name)
			if err != nil {
				return err
			}

			list, err := c.ListRuntimes()
			if err != nil {
				return err
			}

			r := render.DefaultRenderer(cmd.OutOrStdout())
			listRuntimes(r, list)

			return nil
		},
	}

	return cmd
}

func listRuntimes(r render.Renderer, list api.RuntimeList) {
	r.SetHeader([]string{
		"",
		"ID",
		"Name",
	})

	for i, rn := range list {
		r.Append([]string{
			fmt.Sprint(i + 1),
			rn.ID,
			rn.Name,
		})
	}

	r.Render()
}
