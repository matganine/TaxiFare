package config

import (
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

type Configuration struct {
	DataPath  string `mapstructure:"data_path"`
	RidesFile string `mapstructure:"rides_file"`
}

func New(v *viper.Viper) (*Configuration, error) {
	var c Configuration
	if err := v.Unmarshal(&c); err != nil {
		return nil, fmt.Errorf("viper failed to unmarshal app config: %v", err)
	}
	return &c, nil
}


func SetupViper(v *viper.Viper, filename string) {
	if filename != "" {
		v.SetConfigName(filename)
		v.AddConfigPath(".")
		v.AddConfigPath("/etc/config")
	}

	v.SetDefault("data_path", "/home/ralph/go/src/TaxiFare/data/")
	v.SetDefault("rides_file", "rides.json")

	// Set environment variable support:
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.SetTypeByDefaultValue(true)
	v.SetEnvPrefix("TF")
	v.ReadInConfig()
	v.AutomaticEnv()
}
