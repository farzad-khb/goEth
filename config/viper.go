package config

import (
	"log"

	"github.com/spf13/viper"
)

// configures viper
func setViper() {
	viper.SetConfigFile("config/.env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln("error in setViper while reading config", err)
	}

}

//get enviroment varibles with viper
func GetEnv(key string) string {
	setViper()
	value, ok := viper.Get(key).(string)
	if !ok {
		log.Fatalln("error in GetEnv while type assertion")
	}
	return value
}
