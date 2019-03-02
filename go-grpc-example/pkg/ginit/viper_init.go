package ginit

import (
	"github.com/spf13/viper"
)

func InitViper()error{
	viper.SetConfigName("config") // name of config file (without extension)
	viper.AddConfigPath("src/riverside/go-grpc-example/conf/")   // path to look for the config file in
	err := viper.ReadInConfig() // Find and read the config file
	return err
}