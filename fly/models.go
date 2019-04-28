package fly

import (
	"encoding/json"
	"log"
)

type Server struct {
	Name string
	UserName string
	Host string
	Password string
	Port string
	Tag string
	Category string
}


func (server *Server) ToJson() []byte {
	if data, err:= json.Marshal(server); err == nil {
		return data
	}
	log.Fatal("Can't convert server to json.")
	return nil
}

func (server *Server) Merge(otherServer *Server) {
	if server.Port == "" {
		server.Port = otherServer.Port
	}
	if server.UserName == "" {
		server.UserName = otherServer.UserName
	}
	if server.Password == "" {
		server.Password = otherServer.Password
	}
}

func FromJson(data []byte) Server {
	var server Server
	if err := json.Unmarshal(data, &server); err == nil {
		return server
	}
	log.Fatal("Can't covert json to server.")
	return server
}

func FromMap(data map[string]string) Server {
	server := Server{
		Name: data["name"],
		UserName: data["username"],
		Host: data["host"],
		Password: data["password"],
		Port: data["port"],
		Tag: data["tag"],
		Category: data["category"],
	}
	return server
}

type ServerReader interface {
	GetServer(serverName string) Server
	GetServerList() []Server
	UpdateServer(server Server)
	GetDefault() Server
}
