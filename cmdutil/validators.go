package cmdutil

import (
	"github.com/konstellation-io/kli/internal/errors"
	"github.com/spf13/cobra"
)

// CheckServerFlag checks that server is configured by default or passed as a flag.
func CheckServerFlag(cmd *cobra.Command, _ []string) error {
	s, _ := cmd.Flags().GetString("server")

	if s == "" {
		return errors.ErrNoServerConf
	}

	return nil
}
