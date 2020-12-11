package list

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/konstellation-io/kli/api"
	"github.com/konstellation-io/kli/cmdutil"
	"github.com/konstellation-io/kli/internal/render"
)

// NewListCmd creates a new command to list Runtimes.
func NewListCmd(f cmdutil.CmdFactory) *cobra.Command {
	log := f.Logger()
	cmd := &cobra.Command{
		Use:     "ls",
		Aliases: []string{"list"},
		Short:   "List all available runtimes",
		Args:    cmdutil.CheckServerFlag,
		RunE: func(cmd *cobra.Command, args []string) error {
			s, _ := cmd.Flags().GetString("server")
			c, err := f.ServerClient(s)
			if err != nil {
				return err
			}

			list, err := c.ListRuntimes()
			if err != nil {
				return err
			}

			if len(list) == 0 {
				log.Info("No runtimes found.")
				return nil
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
