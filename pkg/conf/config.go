package conf

type DbConfig struct {
	Dialect                      string `yaml:"dialect"`
	DbUrl                        string `yaml:"db_url"`
	MaxIdleConnections           int `yaml:"max_idle_connections"`
	MaxOpenConnections           int `yaml:"max_open_connections"`
	ConnectionMaxLifetimeSeconds int `yaml:"connection_max_lifetime_seconds"`
}
