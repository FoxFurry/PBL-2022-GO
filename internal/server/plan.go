package server

import (
	"net/http"

	"github.com/FoxFurry/PBL-2022-GO/internal/httperr"
	"github.com/FoxFurry/PBL-2022-GO/internal/models"
	"github.com/gin-gonic/gin"
)

func (p *PetFeeder) CreatePlan(c *gin.Context) {
	creator, err := p.getUserFromContext(c)
	if err != nil {
		httperr.Handle(c, err)
		return
	}

	var plan models.Plan

	if err := c.ShouldBindJSON(&plan); err != nil {
		httperr.Handle(c, httperr.New(err.Error(), http.StatusBadRequest))
		return
	}

	if plan.Name == "" {
		httperr.Handle(c, httperr.New("missing pet name", http.StatusBadRequest))
		return
	}

	plan.OwnerID = creator.ID

	createdPlan, err := p.service.CreatePlan(c, plan)
	if err != nil {
		httperr.Handle(c, err)
		return
	}

	c.JSON(http.StatusOK, createdPlan)
}

func (p *PetFeeder) GetPlansByUser(c *gin.Context) {
	creator, err := p.getUserFromContext(c)
	if err != nil {
		httperr.Handle(c, err)
		return
	}

	plans, err := p.service.GetAllPlansByOwner(c, creator.ID)
	if err != nil {
		httperr.Handle(c, err)
		return
	}

	c.JSON(http.StatusOK, plans)
}

func (p *PetFeeder) DeletePlan(c *gin.Context) {
	uuid := c.Param("planUUID")
	if uuid == "" {
		httperr.Handle(c, httperr.New("missing uuid parameter", http.StatusBadRequest))
		return
	}

	if err := p.service.DeletePlan(c, uuid); err != nil {
		httperr.Handle(c, err)
		return
	}

	c.Status(http.StatusOK)
}
