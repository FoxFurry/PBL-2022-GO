package store

import (
	"context"

	"github.com/FoxFurry/PBL-2022-GO/internal/models"
)

func (d *db) CreateUser(ctx context.Context, user models.User) error {
	_, err := d.sql.ExecContext(ctx, `INSERT INTO users (mail, password, uuid) VALUES (?, ?, ?)`, user.Mail, user.Password, user.UUID)
	return err
}

func (d *db) GetUserByMail(ctx context.Context, mail string) (*models.User, error) {
	var u models.User

	if err := d.sql.GetContext(ctx, &u, `SELECT * FROM users WHERE mail=?`, mail); err != nil {
		return nil, err
	}

	return &u, nil
}

func (d *db) GetUserByUUID(ctx context.Context, UUID string) (*models.User, error) {
	var u models.User

	if err := d.sql.GetContext(ctx, &u, `SELECT * FROM users WHERE uuid=?`, UUID); err != nil {
		return nil, err
	}

	return &u, nil
}
