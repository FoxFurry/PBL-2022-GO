package store

import (
	"context"
	"fmt"
	"log"

	"github.com/FoxFurry/PBL-2022-GO/internal/models"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

type DB interface {
	// User

	CreateUser(ctx context.Context, user models.User) error
	GetUserByMail(ctx context.Context, mail string) (*models.User, error)
	GetUserByUUID(ctx context.Context, userUUID string) (*models.User, error)

	// Device

	CreateDevice(ctx context.Context, device models.Device) error
	DeleteDevice(ctx context.Context, deviceUUID string) error
	GetDevicesByOwnerID(ctx context.Context, ownerID uint64) ([]models.Device, error)
	GetDeviceByUUID(ctx context.Context, deviceUUID string) (*models.Device, error)

	// Pet

	CreatePet(ctx context.Context, pet models.Pet) error
	DeletePet(ctx context.Context, petUUID string) error
	GetPetsByOwnerID(ctx context.Context, ownerID uint64) ([]models.Pet, error)
	GetPetByUUID(ctx context.Context, petUUID string) (*models.Pet, error)

	// Plan

	CreatePlan(ctx context.Context, plan models.Plan) error
	DeletePlan(ctx context.Context, planUUID string) error
	GetAllPlansByOwnerID(ctx context.Context, ownerID uint64) ([]models.Plan, error)
	GetPlanByUUID(ctx context.Context, planUUID string) (*models.Plan, error)
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
