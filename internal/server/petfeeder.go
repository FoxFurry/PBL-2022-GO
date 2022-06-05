package server

import (
	"net/http"

	"github.com/FoxFurry/PBL-2022-GO/internal/httperr"
	"github.com/FoxFurry/PBL-2022-GO/internal/models"
	"github.com/FoxFurry/PBL-2022-GO/internal/service"
	"github.com/FoxFurry/PBL-2022-GO/internal/store"
	"github.com/FoxFurry/PBL-2022-GO/internal/util"
	"github.com/gin-gonic/gin"
)

const (
	authSchema  = "Bearer " // Space if required by auth header standard
	uuidKey     = "uuid"
	tokenIssuer = "petfeeder_dev"
	tokenSecret = "5dd0bf305c1eb5b832dbc4169c84ba0aa51704da74b8d2e953dca7b276ee8b0c821e8a764f16fd183c50ca9d9b655cf6159564a1554da81ee16fe01866a462e225ad779b472a62b15d2861c54579875709da3e025e916ab3ac89b165359d0ac529e3739a513eb0de1a2350ab9f741"
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
	v1.Use(p.corsMiddleware)
	{
		user := v1.Group("/user")
		{
			user.POST("/register", p.RegisterUser)  // /v1/user/create
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

func (p *PetFeeder) jwtMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	if len(authHeader) <= len(authSchema) {
		httperr.Handle(c, httperr.New("missing or invalid JWT token", http.StatusUnauthorized))
		return
	}

	token := authHeader[len(authSchema):]

	uuid, err := p.jwt.ValidateToken(token, tokenIssuer, []byte(tokenSecret))
	if err != nil {
		httperr.Handle(c, httperr.WrapHttp(err, "could not validate JWT token", http.StatusUnauthorized))
		return
	}

	c.Set(uuidKey, uuid)
	c.Next()
}

func (p *PetFeeder) corsMiddleware(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}

	c.Next()
}

func (p *PetFeeder) getUserFromContext(c *gin.Context) (*models.User, error) {
	userUUID := c.GetString(uuidKey)
	if userUUID == "" {
		return nil, httperr.New("user uuid missing from context", http.StatusBadRequest)
	}

	user, err := p.service.GetUserByUUID(c, userUUID)
	if err != nil {
		return nil, httperr.Wrap(err, "could not get user from the context")
	}

	return user, nil
}
