package service

import (
	"context"
	"fmt"

	"github.com/FoxFurry/PBL-2022-GO/internal/models"
)

func (p *service) RegisterDevice(ctx context.Context, device models.Device) (*models.Device, error) {
	err := p.db.CreateDevice(ctx, device)
	if err != nil {
		return nil, handleDBError(err, "could not create device")
	}

	return &device, nil
}

func (p *service) DeleteDevice(ctx context.Context, deviceUUID string) error {
	err := p.db.DeleteDevice(ctx, deviceUUID)
	if err != nil {
		return handleDBError(err, "could not delete device")
	}

	return nil
}

func (p *service) GetAllDevicesByOwner(ctx context.Context, ownerID uint64) ([]models.Device, error) {
	devices, err := p.db.GetDevicesByOwnerID(ctx, ownerID)
	if err != nil {
		return nil, handleDBError(err, fmt.Sprintf("could not get all devices by owner id: %d", ownerID))
	}

	return devices, nil
}
