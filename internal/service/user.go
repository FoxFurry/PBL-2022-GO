package service

import (
	"context"
	"net/http"
	"strconv"

	"github.com/FoxFurry/PBL-2022-GO/internal/httperr"
	"github.com/FoxFurry/PBL-2022-GO/internal/models"
	"github.com/cespare/xxhash"
	"github.com/google/uuid"
)

func (p *service) CreateNewUser(ctx context.Context, u models.User) (*models.User, error) {
	u.Password = strconv.FormatUint(xxhash.Sum64String(u.Password), 10)
	u.UUID = uuid.New().String()

	err := p.db.CreateNewUser(ctx, u)

	if err != nil {
		return nil, handleDBError(err, "could not create user")
	}

	return &u, nil
}

func (p *service) GetUserByMail(ctx context.Context, mail string) (*models.User, error) {
	user, err := p.db.GetUserByMail(ctx, mail)
	if err != nil {
		return nil, handleDBError(err, "could not get user by mail")
	}

	return user, nil
}

func (p *service) AuthenticateUser(ctx context.Context, u models.User) (*models.User, error) {
	user, err := p.db.GetUserByMail(ctx, u.Mail)
	if err != nil {
		return nil, handleDBError(err, "could not get user by mail")
	}

	if user.Password == strconv.FormatUint(xxhash.Sum64String(u.Password), 10) { // Compare password hashes
		return user, nil
	} else {
		return nil, httperr.New("Invalid credentials", http.StatusUnauthorized)
	}
}
