package version

import (
	"fmt"
	"strings"

	"github.com/konstellation-io/kli/cmdutil"

	"github.com/spf13/cobra"
)

// NewVersionCmd creates a new command to handle 'version' keyword.
func NewVersionCmd(_ cmdutil.CmdFactory, version, buildDate string) *cobra.Command {
	cmd := &cobra.Command{
		Use:    "version",
		Hidden: true,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprint(cmd.OutOrStdout(), Format(version, buildDate))
		},
	}

	return cmd
}

// Format return the version properly formatted.
func Format(version, buildDate string) string {
	version = strings.TrimPrefix(version, "v")

	if buildDate != "" {
		version = fmt.Sprintf("%s (%s)", version, buildDate)
	}

	return fmt.Sprintf("kli version %s\n", version)
}
