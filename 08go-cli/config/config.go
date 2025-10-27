package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
	Log      LogConfig      `mapstructure:"log"`
}

type ServerConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type DatabaseConfig struct {
	Addr     string `mapstructure:"addr"`     // 格式: "host:port"
	Password string `mapstructure:"password"` // Redis 密码
	DB       int    `mapstructure:"db"`       // 数据库索引 (0-15)
	Protocol int    `mapstructure:"protocol"` // 连接协议
}

type LogConfig struct {
	Level  string `mapstructure:"level"`
	Format string `mapstructure:"format"`
}

func Load() *Config {
	setDefaults()

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		panic(err)
	}

	return &config
}
