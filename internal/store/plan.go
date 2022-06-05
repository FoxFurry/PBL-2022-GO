package store

import (
	"context"

	"github.com/FoxFurry/PBL-2022-GO/internal/models"
)

func (d *db) CreatePlan(ctx context.Context, plan models.Plan) error {
	_, err := d.sql.ExecContext(ctx, `INSERT INTO plans (uuid, name, owner_id) VALUES (?, ?, ?)`,
		plan.UUID,
		plan.Name,
		plan.OwnerID,
	)
	return err
}

func (d *db) DeletePlan(ctx context.Context, planUUID string) error {
	_, err := d.sql.ExecContext(ctx, `DELETE FROM plans WHERE uuid = ?`, planUUID)
	return err
}

func (d *db) GetAllPlansByOwnerID(ctx context.Context, ownerID uint64) ([]models.Plan, error) {
	rows, err := d.sql.QueryContext(ctx, `SELECT * FROM plans WHERE owner_id = ?`, ownerID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var plans []models.Plan
	for rows.Next() {
		var planHolder models.Plan

		err = rows.Scan(&planHolder.ID, &planHolder.UUID, &planHolder.OwnerID, &planHolder.Name, &planHolder.CreatedAt, &planHolder.UpdatedAt)
		if err != nil {
			continue
		}

		plans = append(plans, planHolder)
	}

	return plans, nil
}

func (d *db) GetPlanByUUID(ctx context.Context, planUUID string) (*models.Plan, error) {
	var plan models.Plan

	if err := d.sql.GetContext(ctx, &d, `SELECT * FROM plans WHERE uuid = ?`, planUUID); err != nil {
		return nil, err
	}

	return &plan, nil
}
