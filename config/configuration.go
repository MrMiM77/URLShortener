package config

import (
	"github.com/spf13/viper"
	"strings"
	"sync"
)

type Configuration struct {
	SERVER ServerConfiguration
	DATABASE DatabaseConfiguration
}


var instance *Configuration
var once sync.Once

func GetInstance() *Configuration {
	once.Do(func() {
		var err error
		instance, err = initInstance()
		if err != nil {
			panic(err)
		}
	})
	return instance
}


func initInstance() (*Configuration, error){
	viper.SetDefault("SERVER.HOST", "0.0.0.0")
	viper.SetDefault("SERVER.PORT", 8080)
	viper.SetDefault("DATABASE.HOST", "localhost")
	viper.SetDefault("DATABASE.NAME", "UrlShortener")
	viper.SetDefault("DATABASE.USER", "urls")
	viper.SetDefault("DATABASE.PASSWORD", "urlshortener")
	viper.SetDefault("DATABASE.PORT", 5432)

	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		return &Configuration{}, err
	}

	viper.SetEnvPrefix("URLSHORTENER")
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.AutomaticEnv()

	var instance Configuration
	err = viper.Unmarshal(&instance)
	return &instance, err

}