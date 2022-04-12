package store

import (
	"context"
	"fmt"
	"github.com/FoxFurry/petfeedergateway/internal/models"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"log"
)

type DB interface {
	CreateNewUser(ctx context.Context, u models.User) error
	GetUserByMail(ctx context.Context, mail string) (*models.User, error)
}

type db struct {
	sql *sqlx.DB
}

func initDB(d *db) {
	var (
		dbHost = viper.GetString("DATABASE_HOST")
		dbName = viper.GetString("DATABASE_NAME")
		dbPort = viper.GetString("DATABASE_PORT")
		dbUser = viper.GetString("DATABASE_USER")
		dbPass = viper.GetString("DATABASE_PASSWORD")
		err    error
	)

	if dbHost == "" ||
		dbName == "" ||
		dbPort == "" ||
		dbUser == "" ||
		dbPass == "" {
		log.Panicln("Could not open DB connection. Some of the database env are missing")
	}

	cfg := mysql.Config{
		User:   dbUser,
		Passwd: dbPass,
		Addr:   fmt.Sprintf("%s:%s", dbHost, dbPort),
		DBName: dbName,
	}

	d.sql, err = sqlx.Connect("mysql", cfg.FormatDSN())
	if err != nil {
		log.Panicf("Could not open SQL connection: %v", err)
	}
}

func CreateDB() DB {
	var d db
	initDB(&d)
	return &d
}
