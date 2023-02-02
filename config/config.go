package config

import (
	"fmt"
	"net"
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	defaultConfigPath = "/etc/cpds/cpds-detector/"
	defaultConfigName = "cpds-detector.json"
	defaultConfigType = "json"
)

type Config struct {
	Debug            bool
	LogLevel         string
	DatabaseAddress  string
	DatabasePort     string
	DatabaseUser     string
	DatabasePassword string
	BindAddress      string
	Port             string
	ConfigPath       string
	CertFile         string
	KeyFile          string
}

func New() *Config {
	return &Config{}
}

func (c *Config) LoadConfig(flags *pflag.FlagSet) {
	logrus.Debug("loading cpds-detector configuration")
	cobra.OnInitialize(func() {
		viper.SetConfigType(defaultConfigType)
		logrus.Debugf("set configuration file type: %s", defaultConfigType)
		viper.SetConfigName(defaultConfigName)
		logrus.Debugf("set configuration file path: %s", defaultConfigPath)
		if c.ConfigPath != defaultConfigPath {
			// Use config file from the flag.
			viper.AddConfigPath(c.ConfigPath)
			logrus.Infof("using configuration path: %s, ignore default configuration path")
		} else {
			viper.AddConfigPath(defaultConfigPath)
			logrus.Infof("using default configuration file path: %s", defaultConfigPath)
		}
		logrus.Debugf("using configuration file path: %s", c.ConfigPath)
		logrus.Debugf("default configuration file path: %s", defaultConfigPath)

		c.parseConfigFile(flags)
	})
	c.installFlags(flags)
}

func (c *Config) parseConfigFile(flags *pflag.FlagSet) {
	viper.BindPFlags(flags)
	logrus.Debug("parse configuration file")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error, Using defalut config file path
			// Using defalut config file path
			logrus.Warnf("config file '%s' not found, using defalut config file path: '%s'", c.ConfigPath, defaultConfigPath)
		} else {

			// Config file was found but another error was produced
			logrus.Errorf("fatal error config file: %s", err)
			panic(fmt.Errorf("fatal error config file: %s", err))
		}
	} else {
		logrus.Infof("using config file: %s", viper.ConfigFileUsed())
	}

	*c = Config{
		Debug:            viper.GetBool("debug"),
		LogLevel:         viper.GetString("log-level"),
		DatabaseAddress:  viper.GetString("db-address"),
		DatabasePort:     viper.GetString("db-port"),
		DatabaseUser:     viper.GetString("db-user"),
		DatabasePassword: viper.GetString("db-password"),
		BindAddress:      viper.GetString("bind-address"),
		Port:             viper.GetString("port"),
		CertFile:         viper.GetString("cert-file"),
		KeyFile:          viper.GetString("key-file"),
	}

	logrus.Debugf("parsed debug mod configuration: %s", viper.GetBool("debug"))
	logrus.Debugf("parsed log-level configuration: %s", viper.GetString("log-level"))
	logrus.Debugf("parsed database address configuration: %s", viper.GetString("db-address"))
	logrus.Debugf("parsed database port configuration: %s", viper.GetString("db-port"))
	logrus.Debugf("parsed database user configuration: %s", viper.GetString("db-user"))
	logrus.Debugf("parsed database password configuration: %s", viper.GetString("db-password"))
	logrus.Debugf("parsed bind address configuration: %s", viper.GetString("bind-address"))
}

func (c *Config) CheckConfig() error {
	if net.ParseIP(c.DatabaseAddress) == nil && c.DatabaseAddress != "localhost" {
		return fmt.Errorf("invalid flag: db-address: %s", c.DatabaseAddress)
	}

	if p, err := strconv.Atoi(c.DatabasePort); err != nil {
		return fmt.Errorf("invalid flag: %s, %s", c.DatabasePort, err)
	} else if p < 0 || p > 65535 {
		return fmt.Errorf("invalid port number range: %s, should be 0 - 65535", c.DatabasePort)
	}

	if net.ParseIP(c.BindAddress) == nil {
		return fmt.Errorf("invalid flag: bind-address: %s", c.BindAddress)
	}

	if p, err := strconv.Atoi(c.Port); err != nil {
		return err
	} else if p < 0 || p > 65535 {
		return fmt.Errorf("invalid port number range: %s, should be 0 - 65535", c.Port)
	}

	if _, err := os.Stat(c.CertFile); err != nil {
		return fmt.Errorf("invalid flag: cert-file: %s, %s", c.CertFile, err)
	}

	if _, err := os.Stat(c.KeyFile); err != nil {
		return fmt.Errorf("invalid flag: key-file: %s, %s", c.KeyFile, err)
	}

	return nil
}
