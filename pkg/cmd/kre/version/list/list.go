package list

import (
	"fmt"

	"github.com/konstellation-io/kli/api"
	"github.com/konstellation-io/kli/cmdutil"
	"github.com/konstellation-io/kli/internal/render"
	"github.com/spf13/cobra"
)

// NewListCmd creates a new command to list Versions.
func NewListCmd(f cmdutil.CmdFactory) *cobra.Command {
	log := f.Logger()
	cmd := &cobra.Command{
		Use:     "ls",
		Aliases: []string{"list"},
		Args:    cmdutil.CheckServerFlag,
		Short:   "List all available Versions",
		RunE: func(cmd *cobra.Command, args []string) error {
			s, _ := cmd.Flags().GetString("server")
			c, err := f.ServerClient(s)
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
