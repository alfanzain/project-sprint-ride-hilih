package v1

import (
	"github.com/alfanzain/project-sprint-halo-suster/app/http/handlers"
	"github.com/alfanzain/project-sprint-halo-suster/app/http/middlewares"
	"github.com/alfanzain/project-sprint-halo-suster/app/http/repositories"
	"github.com/alfanzain/project-sprint-halo-suster/app/http/services"
)

func (i *V1Routes) MountUser() {
	gUser := i.Echo.Group("/user")

	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)
	userITService := services.NewUserITService(userRepo)
	userITHandler := handlers.NewUserITHandler(userITService)
	userNurseService := services.NewUserNurseService(userRepo)
	userNurseHandler := handlers.NewUserNurseHandler(userNurseService)

	gUser.GET("", userHandler.Get, middlewares.Authorized(), middlewares.IsRoleIT())

	gUserIT := gUser.Group("/it")
	gUserIT.POST("/register", userITHandler.Register)
	gUserIT.POST("/login", userITHandler.Login)

	gUserNurse := gUser.Group("/nurse")
	gUserNurse.POST("/register", userNurseHandler.Register, middlewares.Authorized(), middlewares.IsRoleIT())
	gUserNurse.POST("/login", userNurseHandler.Login)
	gUserNurse.PUT("/:userID", userNurseHandler.Update)
	gUserNurse.DELETE("/:userID", userNurseHandler.Destroy)
	gUserNurse.POST("/:userID/access", userNurseHandler.GrantAccess)
}
