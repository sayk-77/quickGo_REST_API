package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDataBaseConnection() (*gorm.DB, error) {
	dbc := "host=localhost user=postgres password=postgres dbname=CargoTransportation sslmode=disable"

	db, err := gorm.Open(postgres.Open(dbc), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}

func CloseConnection(db *gorm.DB) error {
	sqlDB, err := db.DB()

	if err != nil {
		return err
	}

	return sqlDB.Close()
}
