package server

import (
	"github.com/FoxFurry/petfeedergateway/internal/httperr"
	"github.com/FoxFurry/petfeedergateway/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
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

	c.JSON(200, responseUser)
}
