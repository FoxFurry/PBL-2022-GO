package server

import (
	"net/http"

	"github.com/FoxFurry/PBL-2022-GO/internal/httperr"
	"github.com/FoxFurry/PBL-2022-GO/internal/models"
	"github.com/gin-gonic/gin"
)

func (p *PetFeeder) RegisterPet(c *gin.Context) {
	creator, err := p.getUserFromContext(c)
	if err != nil {
		httperr.Handle(c, err)
		return
	}

	var pet models.Pet

	if err := c.ShouldBindJSON(&p); err != nil {
		httperr.Handle(c, httperr.New(err.Error(), http.StatusBadRequest))
		return
	}

	if pet.Name == "" {
		httperr.Handle(c, httperr.New("missing pet name", http.StatusBadRequest))
		return
	}

	pet.OwnerID = creator.ID

	createdPet, err := p.service.RegisterPet(c, pet)
	if err != nil {
		httperr.Handle(c, err)
		return
	}

	c.JSON(http.StatusOK, createdPet)
}

func (p *PetFeeder) GetAllPetsByOwner(c *gin.Context) {
	creator, err := p.getUserFromContext(c)
	if err != nil {
		httperr.Handle(c, err)
		return
	}

	pets, err := p.service.GetAllPetsByOwner(c, creator.ID)
	if err != nil {
		httperr.Handle(c, err)
		return
	}

	c.JSON(http.StatusOK, pets)
}

func (p *PetFeeder) DeletePet(c *gin.Context) {
	uuid := c.Param("uuid")
	if uuid == "" {
		httperr.Handle(c, httperr.New("missing uuid parameter", http.StatusBadRequest))
		return
	}

	if err := p.service.DeletePet(c, uuid); err != nil {
		httperr.Handle(c, err)
		return
	}

	c.Status(http.StatusOK)
}
