package config

import (
	"time"

	"github.com/BurntSushi/toml"

	"web-demo/dao"
	"web-demo/utils/log"
)

// Config Project config
type Config struct {
	// Env service run env, `dev` or `pro`
	Env string `toml:"env"`

	// Timeout Graceful shutdown time after the process receives the exit signal
	// reference: https://fleurer.github.io/2020/01/12/note-about-graceful-shutdown/
	Timeout time.Duration `toml:"timeout"`

	// ServiceAddr service Listen addr, eg: 0.0.0.0:8080
	ServiceAddr string `toml:"service_addr"`

	// LogConfig logger config
	LogConfig *log.Config `toml:"log_config"`

	// DBConfig database config
	DBConfig *dao.Config `toml:"db_config"`
}

// DecodeConfig Parse the configuration file to the internal structure,
// the corresponding source file is config/service.conf
func DecodeConfig(configFile string) (cfg *Config, err error) {
	cfg = new(Config)
	if _, err := toml.DecodeFile(configFile, cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
