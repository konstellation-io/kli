package kre

//go:generate mockgen -source=${GOFILE} -destination=$PWD/mocks/${GOFILE} -package=mocks

import (
	"github.com/konstellation-io/kli/api/kre/runtime"
	"github.com/konstellation-io/kli/api/kre/version"
)

// KreInterface first level methods.
type KreInterface interface { // nolint: golint
	Runtime() runtime.RuntimeInterface
	Version() version.VersionInterface
}
