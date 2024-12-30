package cmd

import (
	"go-crud/config"
	"go-crud/network"
	"go-crud/repository"
	"go-crud/service"
)

type Cmd struct {
	config *config.Config

	network    *network.Network
	repository *repository.Repository
	service    *service.Service
}

func NewCmd(filePath string) *Cmd {
	c := &Cmd{
		config: config.NewConfig(filePath),
	}

	c.repository = repository.NewRepository()
	c.service = service.NewServices(c.repository)
	c.network = network.NewNetwork(c.service)

	c.network.ServerStart(c.config.Server.Port)

	return c
}
