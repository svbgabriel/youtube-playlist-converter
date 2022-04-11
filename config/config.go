package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Configurations struct {
	Youtube YoutubeConfigurations
	Spotify SpotifyConfigurations
}

type YoutubeConfigurations struct {
	Key string
}

type SpotifyConfigurations struct {
	ClientId     string
	ClientSecret string
}

func ReadConfig() Configurations {
	// Set the file name of the configurations file
	viper.SetConfigName("config")

	// Set the path to look for the configurations file
	viper.AddConfigPath(".")

	viper.SetConfigType("yml")
	var configuration Configurations

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	return configuration
}
