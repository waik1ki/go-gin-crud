package network

import (
	"go-crud/service"

	"github.com/gin-gonic/gin"
)

type Network struct {
	engin *gin.Engine

	service *service.Service
}

func NewNetwork(service *service.Service) *Network {
	r := &Network{
		engin: gin.New(),
	}

	newUserRouter(r, service.User)

	return r
}

func (n *Network) ServerStart(port string) error {
	return n.engin.Run(":8080")
}
