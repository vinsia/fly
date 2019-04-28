package fly

import (
	"github.com/spf13/viper"
	"log"
)

type ConfigReader struct {
}

func NewConfigReader() *ConfigReader {
	viper.SetConfigName("servers")
	viper.AddConfigPath("$HOME/.ssh")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	return &ConfigReader{}
}

func (configReader *ConfigReader) GetServer(serverName string) Server{
	server := FromMap(viper.GetStringMapString("config.servers." + serverName))
	defaultServer := configReader.GetDefault()
	server.Merge(&defaultServer)
	return server
}

func (configReader *ConfigReader) GetServerList() []Server {
	serverNameList := viper.GetStringMap("config.servers")
	var serverList []Server
	for serverName := range serverNameList {
		serverList = append(serverList, configReader.GetServer(serverName))
	}
	return serverList
}

func (configReader *ConfigReader) GetDefault() Server {
	server := FromMap(viper.GetStringMapString("config.default"))
	return server
}
