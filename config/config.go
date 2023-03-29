package config

import "github.com/spf13/viper"

type Config struct {
	DB_SOURCE string `mapstructure:"DB_SOURCE"`
	API_PORT  string `mapstructure:"API_PORT"`
}

// LoadConfig is a func that returns the configuration fetched from the configuration file
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
