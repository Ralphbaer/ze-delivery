package app

import "github.com/Ralphbaer/ze-delivery/common"

// Config is the top level configuration struct for entire application
type Config struct {
	EnvName               string `env:"ENV_NAME"`
	MongoConnectionString string `env:"MONGO_CONNECTION_STRING"`
	ServerAddress         string `env:"SERVER_ADDRESS"`
	SpecURL               string `env:"SPEC_URL"`
	AdminClientID         string `env:"ADMIN_OAUTH2_CLIENT_ID"`
	AdminClientSecret     string `env:"ADMIN_OAUTH2_CLIENT_SECRET"`
	AdminTokenURL         string `env:"ADMIN_OAUTH2_TOKEN_URL"`
	AdminScopes           string `env:"ADMIN_OAUTH2_SCOPES"`
}

// NewConfig creates a instance of Config
func NewConfig() *Config {
	cfg := &Config{}
	common.SetConfigFromEnvVars(cfg)
	return cfg
}
