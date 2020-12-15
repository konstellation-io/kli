package kre

//go:generate mockgen -source=${GOFILE} -destination=$PWD/mocks/${GOFILE} -package=mocks

import (
	"github.com/konstellation-io/kli/api/kre/runtime"
	"github.com/konstellation-io/kli/api/kre/version"
)

//nolint: golint
type KreInterface interface {
	Runtime() runtime.RuntimeInterface
	Version() version.VersionInterface
}