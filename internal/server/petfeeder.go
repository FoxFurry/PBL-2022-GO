package server

import (
	"github.com/FoxFurry/petfeedergateway/internal/service"
	"github.com/FoxFurry/petfeedergateway/internal/store"
	"github.com/gin-gonic/gin"
)

type PetFeeder struct {
	service service.Service
	gEng    *gin.Engine
}

func New(datastore store.DB) (*PetFeeder, error) {
	ginEngine := gin.Default()

	p := PetFeeder{
		service: service.New(datastore),
		gEng:    ginEngine,
	}

	v1 := ginEngine.Group("/v1")
	{
		user := v1.Group("/user")
		{
			user.POST("/create", p.CreateNewUser)   // /v1/user/create
			user.GET("/getbymail", p.GetUserByMail) // /v1/user/getbymail
			user.POST("/login", p.LoginUser)        // /v1/user/login
		}
	}

	return &p, nil
}

func (p *PetFeeder) Run() {
	p.gEng.Run(":8080")
}

// Add generic helper functions here
