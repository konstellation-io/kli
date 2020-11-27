package kre

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"

	"github.com/konstellation-io/kli/cmdutil"
	kre "github.com/konstellation-io/kli/pkg/cmd/kre/runtime"
)

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

	cmd.PersistentFlags().String("server", f.Config().DefaultServer, "KRE server to use")

	return cmd
}
