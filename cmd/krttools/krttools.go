package krttools

//go:generate mockgen -source=${GOFILE} -destination=$PWD/mocks/${GOFILE} -package=mocks

import (
	"os"

	"github.com/konstellation-io/kre/libs/krt-utils/pkg/builder"
	"github.com/konstellation-io/kre/libs/krt-utils/pkg/validator"
)

// KrtTooler interface with util functions.
type KrtTooler interface {
	Validate(yamlPath string) error
	Build(src, target string) error
}

// KrtTools utils for validating and building krts.
type KrtTools struct {
	validator validator.Validator
	builder   builder.Builder
}

// NewKrtTools factory for KrtTools.
func NewKrtTools() KrtTooler {
	return &KrtTools{
		validator: validator.New(),
		builder:   builder.New(),
	}
}

// Validate validate util for krt command.
func (krt *KrtTools) Validate(yamlPath string) error {
	r, err := os.Open(yamlPath)
	if err != nil {
		return nil
	}

	krtParsed, err := krt.validator.Parse(r)
	if err != nil {
		return nil
	}

	err = krt.validator.Validate(krtParsed)

	return err
}

// BuildKrt build utility for the krt command.
func (krt *KrtTools) Build(src, target string) error {
	err := krt.builder.Build(src, target)
	return err
}
