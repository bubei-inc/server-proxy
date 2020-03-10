package config

import (
	"github.com/DataDog/datadog-agent/pkg/util/log"
	"github.com/spf13/viper"
)

type Config struct {
	ValidatePath string
	ProxyPath    string
	TransferPath string
	ProxyPort    string
	ValidateSite string
}

func ConfigVal(fileName string) *Config {
	viper.SetConfigName(fileName)
	viper.AddConfigPath("./config")
	if err := viper.ReadInConfig(); err != nil {
		log.Info("Format config file error: ", err)
	}
	return &Config{
		ValidatePath: viper.GetString("server.validate-path"),
		ProxyPath:    viper.GetString("server.proxy-path"),
		TransferPath: viper.GetString("server.transfer-path"),
		ProxyPort:    viper.GetString("server.proxy-port"),
		ValidateSite: viper.GetString("server.validate-site"),
	}
}
