package service

import (
	"context"
	"fmt"

	"github.com/FoxFurry/PBL-2022-GO/internal/models"
	"github.com/google/uuid"
)

func (p *service) RegisterPet(ctx context.Context, pet models.Pet) (*models.Pet, error) {
	pet.UUID = uuid.New().String()

	if err := p.db.CreatePet(ctx, pet); err != nil {
		return nil, handleDBError(err, "could not create pet")
	}

	return &pet, nil
}

func (p *service) DeletePet(ctx context.Context, petUUID string) error {
	if err := p.db.DeletePet(ctx, petUUID); err != nil {
		return handleDBError(err, "could not delete pet")
	}

	return nil
}

func (p *service) GetAllPetsByOwner(ctx context.Context, ownerID uint64) ([]models.Pet, error) {
	pets, err := p.db.GetPetsByOwnerID(ctx, ownerID)
	if err != nil {
		return nil, handleDBError(err, fmt.Sprintf("could not get all pets by owner id: %d", ownerID))
	}

	return pets, nil
}
