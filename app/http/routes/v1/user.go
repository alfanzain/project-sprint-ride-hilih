package v1

import (
	"github.com/alfanzain/project-sprint-halo-suster/app/http/handlers"
	"github.com/alfanzain/project-sprint-halo-suster/app/http/middlewares"
	"github.com/alfanzain/project-sprint-halo-suster/app/repositories"
	"github.com/alfanzain/project-sprint-halo-suster/app/services"
)

func (i *V1Routes) MountUser() {
	gUser := i.Echo.Group("/user")

	r := repositories.NewUserITRepository()
	s := services.NewUserITService(r)
	userITHandler := handlers.NewUserITHandler(s)

	gUserIT := gUser.Group("/it")

	gUserIT.POST("/register", userITHandler.Register)
	gUserIT.POST("/login", userITHandler.Login)
	gUserIT.GET("", userITHandler.GetUsers, middlewares.Authorized(), middlewares.IsRoleIT())
}
