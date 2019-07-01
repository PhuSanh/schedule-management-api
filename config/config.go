package config

import (
	"fmt"
	Viper "github.com/spf13/viper"
	"os"
)

type Config struct {
	Host 	string
	Name 	string
	Port 	int
	Version string
	ApiVersion string `mapstructure:"api_version"`
	Debug 	bool
	Mysql 	*Mysql
	//Redis 	*Redis
}

var config *Config

// can not declare another name and call it in main -> not work
func init() {

	var folder string
	envValue := os.Getenv("APPLICATION_ENV")

	switch envValue {
	case "production":
		folder = "prod"
	case "dev":
		folder = "dev"
	default:
		folder = "dev" // change to local
	}

	path := fmt.Sprintf("env/%v", folder)

	// Get base env
	config = new(Config)
	fetchDataToConfig(path, "base", config)
	fetchDataToConfig(path, "mysql", &(config.Mysql))
	//fetchDataToConfig(path, "redis", &(config.Redis))
}

func fetchDataToConfig(configPath string, configName string, result interface{}) {
	viper := Viper.New()
	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)

	err := viper.ReadInConfig() // Find and read the env file
	if err == nil {
		err = viper.Unmarshal(result)
		if err != nil { // Handle errors reading the env file
			panic(fmt.Errorf("Fatal error env file: %s", err))
		}
	}
}

func GetConfig() *Config {
	return config
}