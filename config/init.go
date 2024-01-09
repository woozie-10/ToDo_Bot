package config

import "github.com/spf13/viper"

var Config *viper.Viper

func Init() error {
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	Config = viper.GetViper()

	return nil
}
