package root

import (
	"github.com/spf13/cobra"

	"github.com/konstellation-io/kli/cmdutil"
	"github.com/konstellation-io/kli/pkg/cmd/kre"
	"github.com/konstellation-io/kli/pkg/cmd/server"
	versionCmd "github.com/konstellation-io/kli/pkg/cmd/version"
)

// NewRootCmd creates the base command where all subcommands are added.
func NewRootCmd(f cmdutil.CmdFactory, version, buildDate string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "kli [command] [subcommand] [flags]",
		Short: "Konstellation CLI",
		Long:  `Use Konstellation API from the command line.`,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			d, err := cmd.Flags().GetBool("debug")
			if d {
				f.Config().Debug = true
			}

			return err
		},
		SilenceErrors: true,
		SilenceUsage:  true,
	}

	cmd.SetOut(f.IOStreams().Out)
	cmd.SetErr(f.IOStreams().ErrOut)

	// Hide help command. only --help
	cmd.SetHelpCommand(&cobra.Command{
		Hidden: true,
	})

	cmd.PersistentFlags().Bool("debug", false, "Set debug mode")

	// Child commands
	cmd.AddCommand(versionCmd.NewVersionCmd(f, version, buildDate))
	cmd.AddCommand(server.NewServerCmd(f))
	cmd.AddCommand(kre.NewKRECmd(f))

	return cmd
}
