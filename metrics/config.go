package metrics

import (
	"github.com/achushu/libs/out"
	"gopkg.in/yaml.v2"
)

// Config defines output settings for metrics
type Config struct {
	Enabled    bool        `yaml:"enabled"`
	Runtime    bool        `yaml:"runtime"`
	FileOutput *out.Config `yaml:"file"`
	HTTPOutput *HTTPConfig `yaml:"http"`
}

// HTTPConfig defines the HTTP server to host metrics data on
type HTTPConfig struct {
	Enabled bool `yaml:"enabled"`
	Port    int  `yaml:"port"`
}

// ConfigFromMap returns a Config based on map values.
// Used with spf13/viper.
func ConfigFromMap(config map[string]interface{}) *Config {
	if config == nil {
		return nil
	}
	cfg := new(Config)

	m, err := yaml.Marshal(config)
	if err != nil {
		panic(err)
	}
	if err := yaml.Unmarshal(m, cfg); err != nil {
		panic(err)
	}

	return cfg
}
