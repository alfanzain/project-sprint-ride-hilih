package v1

import (
	"github.com/alfanzain/project-sprint-halo-suster/app/http/handlers"
	"github.com/alfanzain/project-sprint-halo-suster/app/repositories"
	"github.com/alfanzain/project-sprint-halo-suster/app/services"
)

func (i *V1Routes) MountUser() {
	gUser := i.Echo.Group("/user")

	// userRepo := repositories.NewUserRepository()
	// userService := services.NewUserService(userRepo)
	// userHandler := handlers.NewUserHandler(userService)

	// gUser.GET("", userHandler.Get, middlewares.Authorized(), middlewares.IsRoleIT())

	userITRepo := repositories.NewUserITRepository()
	userITService := services.NewUserITService(userITRepo)
	userITHandler := handlers.NewUserITHandler(userITService)

	gUserIT := gUser.Group("/it")

	gUserIT.POST("/register", userITHandler.Register)
	gUserIT.POST("/login", userITHandler.Login)
}
