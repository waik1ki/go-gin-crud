package service

import (
	"go-crud/repository"
	"sync"
)

var (
	serviceInit     sync.Once
	serviceInstance *Service
)

type Service struct {
	// repository
	repository *repository.Repository

	User *User
}

func NewServices(rep *repository.Repository) *Service {
	serviceInit.Do(func() {
		serviceInstance = &Service{
			repository: rep,
		}

		serviceInstance.User = newUserService(rep.User)
	})

	return serviceInstance
}
