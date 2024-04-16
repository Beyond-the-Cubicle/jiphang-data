package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	DatabaseType  string
	DatabaseUrl   string
	SeoulApiKey   string
	GyunggiApiKey string
}

func NewConfig(env string) Config {
	viper.SetConfigName("config_" + env)
	viper.AddConfigPath("./resources")

	setDefaultValues()

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("error during load config file - %w", err))
	}

	databaseConfigViper := viper.Sub("database")
	apiConfigViper := viper.Sub("api.key")

	databaseType := databaseConfigViper.GetString("type")
	databaseUrl := newDatabaseUrl(databaseConfigViper)
	seoulApiKey := apiConfigViper.GetString("seoul")
	gyunggiApiKey := apiConfigViper.GetString("gyunggi")

	return Config{
		DatabaseType:  databaseType,
		DatabaseUrl:   databaseUrl,
		SeoulApiKey:   seoulApiKey,
		GyunggiApiKey: gyunggiApiKey,
	}
}

func setDefaultValues() {
	viper.SetDefault("database.type", "mysql")
	viper.SetDefault("database.url", "127.0.0.1")
	viper.SetDefault("database.port", "3306")
	viper.SetDefault("database.id", "root")
	viper.SetDefault("database.password", "")
	viper.SetDefault("database.database-name", "chulgeun_gil_planner")

	viper.SetDefault("api.key.seoul", "")
	viper.SetDefault("api.key.gyunggi", "")
}

func newDatabaseUrl(viper *viper.Viper) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		viper.GetString("id"),
		viper.GetString("password"),
		viper.GetString("url"),
		viper.GetString("port"),
		viper.GetString("database-name"),
	)
}
