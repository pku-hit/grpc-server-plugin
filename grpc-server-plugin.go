package grpc_server_plugin

import (
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-micro/service/grpc"
	register "github.com/pku-hit/consul-plugin"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var Server = struct {
	RunMode    string `yaml:"runmode"`
	GRPCPort   string `yaml:"grpcport"`
	ServerName string `yaml:"servername"`
}{}

var Service micro.Service

func init() {
	config, err := ioutil.ReadFile("config/server.yml")
	if err != nil {
		fmt.Print(err)
	}
	yaml.Unmarshal(config, &Server)

	Service = grpc.NewService(
		micro.Name(Server.ServerName),
		micro.Registry(register.Reg),
	)
	metadata := make(map[string]string)
	metadata["gRPC.port"] = Server.GRPCPort
	Service.Server().Init(server.Address(":"+Server.GRPCPort), server.Metadata(metadata))
}
