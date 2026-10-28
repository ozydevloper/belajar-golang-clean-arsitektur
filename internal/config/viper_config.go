package config

import (
	"github.com/spf13/viper"
)

func NewConfig() *Config {
	viperConfig := viper.New()

	viperConfig.SetConfigName("config")
	viperConfig.SetConfigType("json")
	viperConfig.AddConfigPath(".")

	err := viperConfig.ReadInConfig()
	if err != nil {
		panic("failed read file config")
	}

	config := new(Config)

	config.Log.Level = viperConfig.GetInt("log.level")

	config.DB.Username = viperConfig.GetString("database.username")
	config.DB.Password = viperConfig.GetString("database.password")
	config.DB.Host = viperConfig.GetString("database.host")
	config.DB.Port = viperConfig.GetInt("database.port")
	config.DB.DBName = viperConfig.GetString("database.name")
	config.DB.SSL = viperConfig.GetString("database.ssl")
	config.DB.OpenConn = viperConfig.GetInt("database.pool.max")
	config.DB.IdleConn = viperConfig.GetInt("database.pool.idle")
	config.DB.Lifetime = viperConfig.GetDuration("database.pool.lifetime")
	config.DB.IdleLifetime = viperConfig.GetDuration("database.pool.idleLifetime")

	config.App.Name = viperConfig.GetString("app.name")
	config.App.Addr = viperConfig.GetString("web.addr")
	config.App.IdleTimeout = viperConfig.GetDuration("web.idle")
	config.App.ReadTimeout = viperConfig.GetDuration("web.read")
	config.App.WriteTimeout = viperConfig.GetDuration("web.write")

	return config
}
