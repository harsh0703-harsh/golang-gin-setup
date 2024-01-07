package initializers

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {

	var err error

	dsn := "host=localhost user=postgres password=root dbname=mydb port=5432 sslmode=disable"

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {

		panic("Failed to connect with db")
	}
}
