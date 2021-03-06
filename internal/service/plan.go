package service

import (
	"context"
	"fmt"

	"github.com/FoxFurry/PBL-2022-GO/internal/models"
	"github.com/google/uuid"
)

func (p *service) CreatePlan(ctx context.Context, plan models.Plan) (*models.Plan, error) {
	plan.UUID = uuid.New().String()

	if err := p.db.CreatePlan(ctx, plan); err != nil {
		return nil, handleDBError(err, "could not get plan")
	}

	return &plan, nil
}

func (p *service) DeletePlan(ctx context.Context, planUUID string) error {
	if err := p.db.DeletePlan(ctx, planUUID); err != nil {
		return handleDBError(err, "could not delete plan")
	}

	return nil
}

func (p *service) GetAllPlansByOwner(ctx context.Context, ownerID uint64) ([]models.Plan, error) {
	plans, err := p.db.GetAllPlansByOwnerID(ctx, ownerID)
	if err != nil {
		return nil, handleDBError(err, fmt.Sprintf("could not get plans by owner id: %d", ownerID))
	}

	return plans, nil
}

func (p *service) GetPlanByUUID(ctx context.Context, planUUID string) (*models.Plan, error) {
	plan, err := p.db.GetPlanByUUID(ctx, planUUID)
	if err != nil {
		return nil, handleDBError(err, "could not get plan")
	}

	return plan, nil
}
