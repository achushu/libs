package out

import "gopkg.in/yaml.v2"

// Config defines settings for Log
type Config struct {
	Enabled   bool          `yaml:"enabled"`
	Filename  string        `yaml:"filename"`
	Async     bool          `yaml:"async"`
	Threshold PriorityLevel `yaml:"threshold"`
	Rotate    *RotateConfig `yaml:"rotate"`
}

// RotateConfig specifies how a rotating log should be rotated
type RotateConfig struct {
	Enabled        bool `yaml:"enabled"`
	MaxSize        int  `yaml:"maxsize"`
	MaxCount       int  `yaml:"maxcount"`
	MaxAge         int  `yaml:"maxage"`
	RotateExisting bool `yaml:"rotate_existing"`
	Compress       bool `yaml:"compress"`
}

// ConfigFromMap returns a Config based on map values.
// Used with spf13/viper.
func ConfigFromMap(m map[string]interface{}) *Config {
	if m == nil {
		return nil
	}
	cfg := new(Config)

	y, err := yaml.Marshal(m)
	if err != nil {
		panic(err)
	}
	if err := yaml.Unmarshal(y, cfg); err != nil {
		panic(err)
	}

	return cfg
}
