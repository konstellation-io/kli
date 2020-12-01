package kre

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"

	"github.com/konstellation-io/kli/cmdutil"
	kre "github.com/konstellation-io/kli/pkg/cmd/kre/runtime"
)

// NewKRECmd creates a new command to handle 'kre' keyword.
func NewKRECmd(f cmdutil.CmdFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "kre",
		Short: "Manage KRE",
		Example: heredoc.Doc(`
			$ kli kre runtime ls
		`),
	}

	cmd.AddCommand(
		kre.NewRuntimeCmd(f),
	)

	cmd.PersistentFlags().StringP("server", "s", f.Config().DefaultServer, "KRE server to use")

	return cmd
}
