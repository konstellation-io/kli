package run

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/guumaster/cligger"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"

	"github.com/konstellation-io/kli/internal/config"
	"github.com/konstellation-io/kli/internal/logger"
	"github.com/konstellation-io/kli/mocks"
	"github.com/konstellation-io/kli/text"
)

type cmd func(f *mocks.MockCmdFactory) *cobra.Command

func NewRunner(t *testing.T, cmd cmd) Runner {
	t.Helper()
	ctrl := gomock.NewController(t)

	dir := setupConfigDir(t)
	defer cleanConfigDir(t, dir)

	cfg := config.NewConfigTest()
	runner := &cmdRunner{t, nil, "", cfg, bytes.NewBufferString("")}

	f := mocks.NewMockCmdFactory(ctrl)
	f.EXPECT().Config().Return(cfg).AnyTimes()

	log := cligger.NewLoggerWithWriter(runner.buf)
	f.EXPECT().Logger().AnyTimes().DoAndReturn(func() logger.Logger {
		return log
	})

	runner.root = cmd(f)

	return runner
}

type Runner interface {
	Run(string) Runner
	RunE(string, error) Runner
	Runf(string, ...interface{}) Runner

	Equal(string) Runner
	Contains(string) Runner
	Containsf(string, ...interface{}) Runner
	Empty() Runner
}

type cmdRunner struct {
	t    *testing.T
	root *cobra.Command
	out  string
	cfg  *config.Config
	buf  *bytes.Buffer
}

func (c *cmdRunner) NewBuffer() *bytes.Buffer {
	return c.buf
}

func (c *cmdRunner) Equal(expected string) Runner {
	expected = text.LinesTrim(expected)
	require.Equal(c.t, expected, c.out)

	return c
}

func (c *cmdRunner) Contains(expected string) Runner {
	require.Contains(c.t, text.LinesTrim(c.out), text.LinesTrim(expected))

	return c
}

func (c *cmdRunner) Containsf(expected string, args ...interface{}) Runner {
	expected = fmt.Sprintf(expected, args...)
	require.Contains(c.t, text.LinesTrim(c.out), text.LinesTrim(expected))

	return c
}

func (c *cmdRunner) Empty() Runner {
	require.Empty(c.t, strings.ReplaceAll(c.out, "\n", ""))

	return c
}

func (c *cmdRunner) Run(cmd string) Runner {
	assert := require.New(c.t)
	b := c.NewBuffer()

	c.out = ""
	c.root.SetOut(b)

	args := strings.Split(cmd, " ")
	args = args[1:]

	c.root.SetArgs(args)

	err := c.root.Execute()
	assert.NoError(err)

	out, err := ioutil.ReadAll(b)
	assert.NoError(err)

	c.out = "\n" + string(out)

	return c
}

func (c *cmdRunner) Runf(format string, args ...interface{}) Runner {
	return c.Run(fmt.Sprintf(format, args...))
}

func (c *cmdRunner) RunE(cmd string, expectedErr error) Runner {
	assert := require.New(c.t)
	b := c.NewBuffer()

	c.out = ""
	c.root.SetOut(b)

	args := strings.Split(cmd, " ")
	args = args[1:]

	c.root.SetArgs(args)

	actualErr := c.root.Execute()
	assert.EqualError(actualErr, expectedErr.Error())

	out, err := ioutil.ReadAll(b)
	assert.NoError(err)

	c.out = "\n" + string(out)

	return c
}
