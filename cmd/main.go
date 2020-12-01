package main

import (
	"github.com/konstellation-io/kli/cmdutil"
	"github.com/konstellation-io/kli/internal/build"
	"github.com/konstellation-io/kli/pkg/cmd/root"
)

func main() {
	buildDate := build.Date
	buildVersion := build.Version

	cmdFactory := cmdutil.NewCmdFactory(buildVersion)

	rootCmd := root.NewRootCmd(cmdFactory, buildVersion, buildDate)

	if err := rootCmd.Execute(); err != nil {
		cmdFactory.Logger().Fatal("execution error: %s\n", err)
	}
}
