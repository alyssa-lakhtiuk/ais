package config

import (
	"github.com/spf13/viper"
)

var SigningKey string

type (
	Config struct {
		ListenUrl  string     `mapstructure:"ListeningURL"`
		Salt       string     `mapstructure:"Salt"`
		SigningKey string     `mapstructure:"SigningKey"`
		Postgresql Postgresql `mapstructure:"Postgresql"`
	}

	Postgresql struct {
		Password string `mapstructure:"Password"`
		Username string `mapstructure:"User"`
		Host     string `mapstructure:"Host"`
		Port     int    `mapstructure:"Port"`
		DBname   string `mapstructure:"DB"`
	}
)

func New(fileName string) (Config, error) {
	var conf Config
	err := viperSetup(fileName)
	if err != nil {
		return Config{}, err

	}

	err = viper.ReadInConfig()
	if err != nil {
		return Config{}, err

	}
	err = viper.Unmarshal(&conf)
	if err != nil {
		return Config{}, err
	}
	SigningKey = conf.SigningKey
	return conf, nil
}

func viperSetup(filename string) (error error) {
	viper.AddConfigPath(".")
	viper.SetConfigName(filename)
	viper.SetConfigType("json")

	viper.AutomaticEnv()
	return nil
}
