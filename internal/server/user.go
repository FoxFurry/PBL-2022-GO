package server

import (
	"net/http"
	"time"

	"github.com/FoxFurry/PBL-2022-GO/internal/httperr"
	"github.com/FoxFurry/PBL-2022-GO/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

const (
	tokenIssuer = "petfeeder_dev"
)

func (p *PetFeeder) CreateNewUser(c *gin.Context) {
	var u models.User

	if err := c.ShouldBindJSON(&u); err != nil {
		httperr.Handle(c, httperr.New(err.Error(), http.StatusBadRequest))
		return
	}

	if err := u.ValidateAll(); err != nil {
		httperr.Handle(c, err)
		return
	}

	createdUser, err := p.service.CreateNewUser(c, u)
	if err != nil {
		httperr.Handle(c, err)
		return
	}

	c.JSON(http.StatusOK, models.User{
		UUID: createdUser.UUID,
		Mail: createdUser.Mail,
	})
}

func (p *PetFeeder) GetUserByMail(c *gin.Context) {
	var u models.User

	if err := c.ShouldBindJSON(&u); err != nil {
		httperr.Handle(c, httperr.New(err.Error(), http.StatusBadRequest))
		return
	}

	if err := u.ValidateMail(); err != nil {
		httperr.Handle(c, err)
		return
	}

	responseUser, err := p.service.GetUserByMail(c, u.Mail)
	if err != nil {
		httperr.Handle(c, err)
		return
	}

	c.JSON(http.StatusOK, responseUser)
}

func (p *PetFeeder) LoginUser(c *gin.Context) {
	var u models.User

	if err := c.ShouldBindJSON(&u); err != nil {
		httperr.Handle(c, httperr.New(err.Error(), http.StatusBadRequest))
		return
	}

	if err := u.ValidateMail(); err != nil {
		httperr.Handle(c, err)
		return
	}

	responseUser, err := p.service.AuthenticateUser(c, u)
	if err != nil {
		httperr.Handle(c, err)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uuid": responseUser.UUID,
		"exp":  time.Now().Add(time.Hour).Unix(),
		"iss":  tokenIssuer,
	})

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		httperr.Handle(c, httperr.WrapHttp(err, "could not sign token", http.StatusInternalServerError))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}
