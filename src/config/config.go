package config

import (
	log "agora-vnf-manager/core/log"

	goenv "github.com/evilwire/go-env"
)

type configServer struct {
	Port string `env:"PORT"`
}

var server_defaults configServer = configServer{
	Port: "3000",
}

type configDatabaseServer struct {
	DbUsername               string `env:"USERNAME"`
	DbPassword               string `env:"PASSWORD"`
	DbAddress                string `env:"ADDRESS"`
	DbPort                   string `env:"PORT"`
	DbSchema                 string `env:"SCHEMA"`
	DbName                   string `env:"DBNAME"`
	DbMaxOpenConnections     int    `env:"MAX_OPEN_CONNS"`
	DbMaxIdleConnections     int    `env:"MAX_IDLE_CONNS"`
	DbMaxConnectionsLifetime int    `env:"MAX_CONNS_LIFETIME"`
}

type Config struct {
	Database configDatabaseServer `env:"DATABASE_"`
	Server   configServer         `env:"SERVER_"`
}

var cfg = Config{}

var config_initialized = false

func (cfg *Config) fill_undefined_with_defaults() {
	if cfg.Server.Port == "" {
		cfg.Server.Port = server_defaults.Port
	}
}

func init_config() (config *Config) {
	marshaller := goenv.DefaultEnvMarshaler{
		Environment: goenv.NewOsEnvReader(),
	}
	err := marshaller.Unmarshal(&cfg)
	if err != nil {
		log.Errorf("[init_config]: %s", err.Error())
	} else {
		log.Info("[init_config]: Configuration inicitated successfuly")
	}
	cfg.fill_undefined_with_defaults()
	return &cfg
}

func GetConfig() *Config {
	if !config_initialized {
		cfg := init_config()
		config_initialized = true
		return cfg
	}
	return &cfg
}
