package iostreams

import (
	"io"
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/mattn/go-colorable"
	"github.com/mattn/go-isatty"
)

const (
	spinnerDuration = 400 * time.Millisecond
)

type IOStreams struct {
	In     io.ReadCloser
	Out    io.Writer
	ErrOut io.Writer

	progressIndicator *spinner.Spinner

	progressIndicatorEnabled bool

	stdinTTYOverride  bool
	stdinIsTTY        bool
	stdoutTTYOverride bool
	stdoutIsTTY       bool
	stderrTTYOverride bool
	stderrIsTTY       bool
}

func System() *IOStreams {
	stdoutIsTTY := isTerminal(os.Stdout)
	stderrIsTTY := isTerminal(os.Stderr)

	st := &IOStreams{
		In:     os.Stdin,
		Out:    colorable.NewColorable(os.Stdout),
		ErrOut: colorable.NewColorable(os.Stderr),
	}

	if stdoutIsTTY && stderrIsTTY {
		st.progressIndicatorEnabled = true
	}

	// prevent duplicate isTerminal queries now that we know the answer
	st.SetStdoutTTY(stdoutIsTTY)
	st.SetStderrTTY(stderrIsTTY)

	return st
}

func (s *IOStreams) SetStdinTTY(isTTY bool) {
	s.stdinTTYOverride = true
	s.stdinIsTTY = isTTY
}

func (s *IOStreams) IsStdinTTY() bool {
	if s.stdinTTYOverride {
		return s.stdinIsTTY
	}

	if stdin, ok := s.In.(*os.File); ok {
		return isTerminal(stdin)
	}

	return false
}

func (s *IOStreams) SetStdoutTTY(isTTY bool) {
	s.stdoutTTYOverride = true
	s.stdoutIsTTY = isTTY
}

func (s *IOStreams) IsStdoutTTY() bool {
	if s.stdoutTTYOverride {
		return s.stdoutIsTTY
	}

	if stdout, ok := s.Out.(*os.File); ok {
		return isTerminal(stdout)
	}

	return false
}

func (s *IOStreams) SetStderrTTY(isTTY bool) {
	s.stderrTTYOverride = true
	s.stderrIsTTY = isTTY
}

func (s *IOStreams) IsStderrTTY() bool {
	if s.stderrTTYOverride {
		return s.stderrIsTTY
	}

	if stderr, ok := s.ErrOut.(*os.File); ok {
		return isTerminal(stderr)
	}

	return false
}

func (s *IOStreams) StartProgressIndicator() {
	if !s.progressIndicatorEnabled {
		return
	}

	sp := spinner.New(spinner.CharSets[11], spinnerDuration, spinner.WithWriter(s.ErrOut))
	sp.Start()
	s.progressIndicator = sp
}

func (s *IOStreams) StopProgressIndicator() {
	if s.progressIndicator == nil {
		return
	}

	s.progressIndicator.Stop()
	s.progressIndicator = nil
}

func isTerminal(f *os.File) bool {
	return isatty.IsTerminal(f.Fd()) || isatty.IsCygwinTerminal(f.Fd())
}
