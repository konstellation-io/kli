package version

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/konstellation-io/kli/api/kre/version"
	"github.com/konstellation-io/kli/cmd/factory"
	"github.com/konstellation-io/kli/internal/render"
	"github.com/konstellation-io/kli/pkg/cmd/args"
)

// NewListCmd creates a new command to list Versions.
func NewListCmd(f factory.CmdFactory) *cobra.Command {
	log := f.Logger()
	cmd := &cobra.Command{
		Use:     "ls",
		Aliases: []string{"list"},
		Args:    args.CheckServerFlag,
		Short:   "List all available Versions",
		RunE: func(cmd *cobra.Command, args []string) error {
			s, _ := cmd.Flags().GetString("server")
			c, err := f.KreClient(s)
			if err != nil {
				return err
			}

			runtime, err := cmd.Flags().GetString("runtime")
			if err != nil {
				return err
			}

			list, err := c.Version().List(runtime)
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
	cmd.Flags().StringP("runtime", "r", "", "Filter for specific runtime")
	_ = cmd.MarkFlagRequired("runtime")

	return cmd
}

func listVersions(r render.Renderer, list version.List) {
	r.SetHeader([]string{
		"",
		"ID",
		"Name",
		"Status",
	})

	for i, rn := range list {
		r.Append([]string{
			fmt.Sprint(i + 1),
			rn.ID,
			rn.Name,
			rn.Status,
		})
	}

	r.Render()
}
