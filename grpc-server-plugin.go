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
	ConsulHost string `yaml:"ConsulHost"`
	ConsulPort string `yaml:"ConsulPort"`
	ServerName string `yaml:"servername"`
}{}

var Service micro.Service

func init() {
	config, err := ioutil.ReadFile("config/server.yml")
	if err != nil {
		fmt.Print(err)
	}
	yaml.Unmarshal(config, &Server)

	consulAddress := Server.ConsulHost + ":" + Server.ConsulPort
	reg := consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			consulAddress,
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
