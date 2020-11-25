package main

import (
	"github.com/konstellation-io/kli/internal/build"
	"github.com/konstellation-io/kli/pkg/cmd/factory"
	"github.com/konstellation-io/kli/pkg/cmd/root"

	"github.com/guumaster/cligger"
)

func main() {
	buildDate := build.Date
	buildVersion := build.Version

	cmdFactory := factory.New(buildVersion)

	cligger.SetWriter(cmdFactory.IOStreams.Out)

	rootCmd := root.NewRootCmd(cmdFactory, buildVersion, buildDate)

	if err := rootCmd.Execute(); err != nil {
		cligger.Fatal("execution error: %s\n", err)
	}
}
