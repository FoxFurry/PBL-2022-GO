package store

import (
	"context"
	"github.com/FoxFurry/petfeedergateway/internal/models"
)

func (d *db) CreateNewUser(ctx context.Context, user models.User) error {
	_, err := d.sql.ExecContext(ctx, `INSERT INTO user (mail, password, uuid) VALUES (?, ?, ?)`, user.Mail, user.Password, user.UUID)
	return err
}

func (d *db) GetUserByMail(ctx context.Context, mail string) (*models.User, error) {
	var u models.User

	err := d.sql.GetContext(ctx, &u, `SELECT * FROM user WHERE mail=?`, mail)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (d *db) GetUserByUUID(ctx context.Context, UUID string) (*models.User, error) {
	var u models.User

	err := d.sql.GetContext(ctx, &u, `SELECT * FROM user WHERE uuid=?`, UUID)
	if err != nil {
		return nil, err
	}

	return &u, nil
}
