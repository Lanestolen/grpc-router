package config

import (
	"io/ioutil"
	"os"

	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v2"
)

var (
	defaultConfigFile = "grpc-router/config.yml"
)

type Config struct {
	WireGuardDir  string `yaml:"wireguard-dir"`
	ServiceConfig struct {
		Domain struct {
			Endpoint string `yaml:"endpoint"`
			Port     uint   `yaml:"port"`
		} `yaml:"domain"`
		TLS struct {
			Enabled  bool   `yaml:"enabled"`
			CertFile string `yaml:"certFile"`
			CertKey  string `yaml:"certKey"`
			CAFile   string `yaml:"caFile"`

			Directory string `yaml:"directory"`
		} `yaml:"tls"`
		Auth struct {
			AKey string `yaml:"aKey"`
			SKey string `yaml:"sKey"`
		} `yaml:"auth"`
	} `yaml:"service-config"`
}

func ParseConfig() (Config, error) {
	var conf Config
	configPath := os.Getenv("CONFIG_PATH")

	if configPath == "" {
		confDir, err := os.UserConfigDir()
		if err != nil {
			return Config{}, err
		}

		configPath = confDir + "/" + defaultConfigFile
		log.Debug().Str("default config file", configPath).Msg("CONFIG_PATH env not set using default")
	}

	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return Config{}, err
	}

	if err := yaml.Unmarshal(data, &conf); err != nil {
		return Config{}, err
	}

	return conf, nil
}
