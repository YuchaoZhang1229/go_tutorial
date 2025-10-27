package config

import "github.com/spf13/viper"

func setDefaults() {
	// Server defaults
	viper.SetDefault("server.host", "localhost")
	viper.SetDefault("server.port", 8080)

	// Database defaults
	viper.SetDefault("database.addr", "localhost:6379")
	viper.SetDefault("database.password", "")
	viper.SetDefault("database.dbname", 0)
	viper.SetDefault("database.protocol", 2)

	// Log defaults
	viper.SetDefault("log.level", "info")
	viper.SetDefault("log.format", "json")
}
