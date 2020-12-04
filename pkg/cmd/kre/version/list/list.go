package list

import (
	"fmt"

	"github.com/konstellation-io/kli/api"
	"github.com/konstellation-io/kli/cmdutil"
	"github.com/konstellation-io/kli/internal/errors"
	"github.com/konstellation-io/kli/internal/render"
	"github.com/spf13/cobra"
)

func NewListCmd(f cmdutil.CmdFactory) *cobra.Command {
	log := f.Logger()
	cmd := &cobra.Command{
		Use:     "ls",
		Aliases: []string{"list"},
		Short:   "List all available Versions",
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

			runtime, err := cmd.Flags().GetString("runtime")
			if err != nil {
				return err
			}

			list, err := c.ListVersions(runtime)
			if err != nil {
				return err
			}

			if len(list) == 0 {
				log.Info("No versions found.")
				return nil
			}

			r := render.DefaultRenderer(cmd.OutOrStdout())
			listVersions(r, list)

			return nil
		},
	}

	return cmd
}

func listVersions(r render.Renderer, list api.VersionList) {
	r.SetHeader([]string{
		"",
		"Name",
		"Status",
	})

	for i, rn := range list {
		r.Append([]string{
			fmt.Sprint(i + 1),
			rn.Name,
			rn.Status,
		})
	}

	r.Render()
}
