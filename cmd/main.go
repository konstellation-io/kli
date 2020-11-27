package main

import (
	"github.com/konstellation-io/kli/cmdutil"
	"github.com/konstellation-io/kli/internal/build"
	"github.com/konstellation-io/kli/pkg/cmd/root"

	"github.com/guumaster/cligger"
)

func main() {
	buildDate := build.Date
	buildVersion := build.Version

	cmdFactory := cmdutil.NewCmdFactory(buildVersion)

	cligger.SetWriter(cmdFactory.IOStreams().Out)

	rootCmd := root.NewRootCmd(cmdFactory, buildVersion, buildDate)

	if err := rootCmd.Execute(); err != nil {
		cligger.Fatal("execution error: %s\n", err)
	}
}
