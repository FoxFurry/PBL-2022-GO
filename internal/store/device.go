package store

import (
	"context"

	"github.com/FoxFurry/PBL-2022-GO/internal/models"
)

func (d *db) CreateDevice(ctx context.Context, device models.Device) error {
	_, err := d.sql.ExecContext(ctx, `INSERT INTO devices (owner_id, uuid, name, location, address) VALUES (?, ?, ?, ?, ?)`,
		device.OwnerID,
		device.UUID,
		device.Name,
		device.Location,
		device.Address,
	)
	return err
}

func (d *db) DeleteDevice(ctx context.Context, deviceUUID string) error {
	_, err := d.sql.ExecContext(ctx, `DELETE FROM devices WHERE uuid = ?`, deviceUUID)
	return err
}

func (d *db) GetDevicesByOwnerID(ctx context.Context, ownerID uint64) ([]models.Device, error) {
	rows, err := d.sql.QueryContext(ctx, `SELECT * FROM devices WHERE owner_id = ?`, ownerID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var devices []models.Device
	for rows.Next() {
		var deviceHolder models.Device

		if err := rows.Scan(
			&deviceHolder.ID,
			&deviceHolder.UUID,
			&deviceHolder.OwnerID,
			&deviceHolder.Name,
			&deviceHolder.Location,
			&deviceHolder.Address,
			&deviceHolder.CreatedAt,
			&deviceHolder.UpdatedAt,
		); err != nil {
			continue
		}

		devices = append(devices, deviceHolder)
	}

	return devices, nil
}

func (d *db) GetDeviceByUUID(ctx context.Context, deviceUUID string) (*models.Device, error) {
	var device models.Device

	if err := d.sql.GetContext(ctx, &d, `SELECT * FROM devices WHERE uuid = ?`, deviceUUID); err != nil {
		return nil, err
	}

	return &device, nil
}
