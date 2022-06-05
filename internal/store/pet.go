package store

import (
	"context"

	"github.com/FoxFurry/PBL-2022-GO/internal/models"
)

func (d *db) CreatePet(ctx context.Context, pet models.Pet) error {
	_, err := d.sql.ExecContext(ctx, `INSERT INTO pets (plan_id, uuid, name) VALUES (?, ?, ?)`,
		pet.PlanID,
		pet.UUID,
		pet.Name,
	)
	return err
}

func (d *db) DeletePet(ctx context.Context, petUUID string) error {
	_, err := d.sql.ExecContext(ctx, `DELETE FROM devices WHERE uuid = ?`, petUUID)
	return err
}

func (d *db) GetPetsByOwnerID(ctx context.Context, ownerID uint64) ([]models.Pet, error) {
	rows, err := d.sql.QueryContext(ctx, `SELECT * FROM pets WHERE owner_id = ?`, ownerID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var pets []models.Pet
	for rows.Next() {
		var petHolder models.Pet

		if err := rows.Scan(
			&petHolder.ID,
			&petHolder.UUID,
			&petHolder.OwnerID,
			&petHolder.PlanID,
			&petHolder.Name,
			&petHolder.CreatedAt,
			&petHolder.UpdatedAt,
		); err != nil {
			continue
		}

		pets = append(pets, petHolder)
	}

	return pets, nil
}

func (d *db) GetPetByUUID(ctx context.Context, petUUID string) (*models.Pet, error) {
	var pet models.Pet

	if err := d.sql.GetContext(ctx, &pet, `SELECT * FROM pets WHERE uuid=?`, petUUID); err != nil {
		return nil, err
	}

	return &pet, nil
}
