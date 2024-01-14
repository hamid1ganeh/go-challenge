package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	ServerPort   string `mapstructure:"SERVER_PORT"`
	DbName       string `mapstructure:"DB_Name"`
	AWSREGION string `mapstructure:"AWS_REGION"`
	AwsAccesskey string `mapstructure:"AWS_Accesskey"`
	AwsSecretKey string `mapstructure:"AWS_SecretKey"`
}

var AppConfig Config

func LoadConfig() error {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
