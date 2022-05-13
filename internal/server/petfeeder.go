package server

import (
	"github.com/FoxFurry/PBL-2022-GO/internal/service"
	"github.com/FoxFurry/PBL-2022-GO/internal/store"
	"github.com/FoxFurry/PBL-2022-GO/internal/util"
	"github.com/gin-gonic/gin"
)

type PetFeeder struct {
	service service.Service
	gEng    *gin.Engine
	jwt     util.JWTProvider
}

func New(datastore store.DB) (*PetFeeder, error) {
	ginEngine := gin.Default()

	p := PetFeeder{
		service: service.New(datastore),
		gEng:    ginEngine,
		jwt:     util.NewJWT(),
	}

	v1 := ginEngine.Group("/v1")
	{
		user := v1.Group("/user")
		{
			user.POST("/create", p.CreateNewUser)   // /v1/user/create
			user.GET("/getbymail", p.GetUserByMail) // /v1/user/getbymail
			user.POST("/login", p.LoginUser)        // /v1/user/login
		}

		device := v1.Group("/device")
		{
			device.POST("/register") // /v1/device/register
		}
	}

	return &p, nil
}

func (p *PetFeeder) Run() {
	p.gEng.Run(":8080")
}

// Add generic helper functions here
