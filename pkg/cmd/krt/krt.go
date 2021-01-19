package krt

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"

	"github.com/konstellation-io/kli/cmd/factory"
)

// NewKRTCmd creates a new command to handle 'krt' keyword.
func NewKRTCmd(f factory.CmdFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "krt",
		Short: "Utils for KRT creation and validation",
		Example: heredoc.Doc(`
			$ kli krt validate file.krt
		`),
	}

	cmd.AddCommand(
		NewCreateCmd(f),
		NewValidateCmd(f),
	)

	return cmd
}
