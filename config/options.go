package config

import (
	"github.com/spf13/pflag"
)

func (c *Config) installFlags(flags *pflag.FlagSet) {
	flags.StringVar(&c.ConfigPath, "config", defaultConfigPath, "Configuration file Path")
	flags.BoolVarP(&c.Debug, "debug", "D", false, "Enable debug mode")
	flags.StringVarP(&c.LogLevel, "log-level", "l", "info", `Set the logging level ("debug"|"info"|"warn"|"error"|"fatal")`)

	flags.StringVar(&c.DatabaseAddress, "db-address", "localhost", "Specify database address")
	flags.StringVar(&c.DatabasePort, "db-port", "3306", "Specify database port")
	flags.StringVar(&c.DatabaseUser, "db-user", "admin", "Database username")
	flags.StringVar(&c.DatabasePassword, "db-password", "", "Database password")

	flags.StringVar(&c.BindAddress, "bind-address", "0.0.0.0", "Server bind address")
	flags.StringVarP(&c.Port, "port", "p", "19081", "Port number to listen")

	certPath, keyPath := GetCertPath(), GetKeyPath()
	// TODO: make certificate and key file by openssl instead of using certificate template
	flags.StringVar(&c.CertFile, "cert-file", certPath, "identify HTTPS client using this SSL certificate file")
	flags.StringVar(&c.KeyFile, "key-file", keyPath, "Identify HTTPS client using this SSL key file")
}
