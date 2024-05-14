package v1

import (
	"github.com/alfanzain/project-sprint-halo-suster/app/http/handlers"
	"github.com/alfanzain/project-sprint-halo-suster/app/repositories"
	"github.com/alfanzain/project-sprint-halo-suster/app/services"
)

func (i *V1Routes) MountUser() {
	gUser := i.Echo.Group("/user")

	userITHandler := handlers.NewUserITHandler(services.NewUserITService(
		repositories.NewUserITRepository(),
	))

	gUserIT := gUser.Group("/it")

	gUserIT.POST("/register", userITHandler.Register)
	// gUserIT.POST("/login", userITHandler.Login)
}
