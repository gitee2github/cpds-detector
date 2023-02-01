package config

import (
	"testing"

	"github.com/spf13/pflag"
	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
)

func TestInstallFlags(t *testing.T) {
	flags := pflag.NewFlagSet("testing", pflag.ContinueOnError)
	conf := New()
	conf.installFlags(flags)

	err := flags.Parse([]string{
		"--config=/foo/config",
		"--db-address=1.2.3.4",
		"--port=4321",
		"--debug",
		"--log-level=error",
		"--db-port=7890",
		"--db-user=test",
		"--db-password=password1",
		"--bind-address=4.3.2.1",
	})

	assert.Check(t, err)
	assert.Check(t, is.Equal("/foo/config", conf.ConfigPath))
	assert.Check(t, is.Equal("1.2.3.4", conf.DatabaseAddress))
	assert.Check(t, is.Equal(conf.Port, "4321"))
	assert.Equal(t, conf.Debug, true)
	assert.Equal(t, conf.LogLevel, "error")
	assert.Equal(t, conf.DatabasePort, "7890")
	assert.Equal(t, conf.DatabaseUser, "test")
	assert.Equal(t, conf.DatabasePassword, "password1")
	assert.Equal(t, conf.BindAddress, "4.3.2.1")
}
