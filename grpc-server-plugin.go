package grpc_server_plugin

import (
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-micro/service/grpc"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"time"
)

var Server = struct {
	RunMode    string `yaml:"runmode"`
	GRPCPort   string `yaml:"grpcport"`
	ServerName string `yaml:"servername"`
}{}

var Registry = struct {
	Type string `yaml:"type"`
	Host string `yaml:"host"`
}{}

var Service micro.Service

func init() {
	config, err := ioutil.ReadFile("config/server.yml")
	registyConfig, err := ioutil.ReadFile("config/registry.yml")
	if err != nil {
		fmt.Print(err)
	}
	yaml.Unmarshal(config, &Server)
	yaml.Unmarshal(registyConfig, &Registry)

	reg := consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			Registry.Host,
		}
	})

	Service = grpc.NewService(
		micro.Name("grpc-"+Server.ServerName),
		micro.Registry(reg),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*15),
	)
	Service.Server().Init(server.Address(":" + Server.GRPCPort))
}
