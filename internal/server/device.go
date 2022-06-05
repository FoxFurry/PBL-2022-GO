package server

import (
	"net/http"

	"github.com/FoxFurry/PBL-2022-GO/internal/httperr"
	"github.com/FoxFurry/PBL-2022-GO/internal/models"
	"github.com/gin-gonic/gin"
)

func (p *PetFeeder) RegisterDevice(c *gin.Context) {
	creator, err := p.getUserFromContext(c)
	if err != nil {
		httperr.Handle(c, err)
		return
	}

	var d models.Device

	if err := c.ShouldBindJSON(&d); err != nil {
		httperr.Handle(c, httperr.New(err.Error(), http.StatusBadRequest))
		return
	}

	if d.Name == "" || d.Address == "" || d.Location == "" {
		httperr.Handle(c, httperr.New("missing required fields", http.StatusBadRequest))
		return
	}

	d.OwnerID = creator.ID

	createdDevice, err := p.service.RegisterDevice(c, d)
	if err != nil {
		httperr.Handle(c, err)
		return
	}

	c.JSON(http.StatusOK, createdDevice)
}

func (p *PetFeeder) DeleteDevice(c *gin.Context) {
	uuid := c.Param("uuid")
	if uuid == "" {
		httperr.Handle(c, httperr.New("missing uuid parameter", http.StatusBadRequest))
		return
	}

	if err := p.service.DeleteDevice(c, uuid); err != nil {
		httperr.Handle(c, err)
		return
	}

	c.Status(http.StatusOK)
}

func (p *PetFeeder) GetDevicesByOwnerUUID(c *gin.Context) {
	uuid := c.Param("uuid")

	if uuid == "" {
		httperr.Handle(c, httperr.New("missing uuid parameter", http.StatusBadRequest))
		return
	}

	user, err := p.service.GetUserByUUID(c, uuid)
	if err != nil {
		httperr.Handle(c, httperr.Wrap(err, "could not get user"))
		return
	}

	devices, err := p.service.GetAllDevicesByOwner(c, user.ID)
	if err != nil {
		httperr.Handle(c, httperr.Wrap(err, "could not get devices"))
		return
	}

	c.JSON(http.StatusOK, devices)
}
