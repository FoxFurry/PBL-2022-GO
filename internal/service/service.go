package service

import (
	"context"
	"fmt"
	"github.com/FoxFurry/petfeedergateway/internal/httperr"
	"github.com/FoxFurry/petfeedergateway/internal/models"
	"github.com/FoxFurry/petfeedergateway/internal/store"
	"github.com/go-sql-driver/mysql"
	"net/http"
)

type Service interface {
	CreateNewUser(ctx context.Context, u models.User) (*models.User, error)
	GetUserByMail(ctx context.Context, mail string) (*models.User, error)
	AuthenticateUser(ctx context.Context, u models.User) (string, error)
}

type service struct {
	db store.DB
}

func New(datastore store.DB) Service {
	return &service{
		db: datastore,
	}
}

func handleDBError(err error, msg string) error {
	if mysqlErr, ok := err.(*mysql.MySQLError); ok {
		switch mysqlErr.Number {
		case 1062:
			return httperr.New(fmt.Sprintf("%s: entry already exists", msg), http.StatusBadRequest)
		case 1741:
			return httperr.New(fmt.Sprintf("%s: key not found", msg), http.StatusNotFound)
		}
	}
	// TODO: Change in live environment
	return httperr.New(fmt.Sprintf("%s: unknown internal error: %v", msg, err), http.StatusInternalServerError)
}
