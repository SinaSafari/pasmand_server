package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"pasmand/internal/models"
)

var DB *gorm.DB

func SetupDatabase() {
	var err error
	//connectionString := os.Getenv("DB_CONNECTION_STRING")
	//if len(connectionString) == 0 {
	//	log.Fatal("connection string is not provided as environment variable")
	//}
	DB, err = gorm.Open(postgres.Open("postgres://db_user:db_password@postgresdb:5432/pasmand_db?sslmode=disable"), &gorm.Config{})
	if err != nil {
		log.Fatal("connection is not established: " + err.Error())
	}

	// TODO: add auto migration feature here
	err = DB.AutoMigrate(&models.User{}, &models.Address{}, &models.Category{}, &models.Product{})
	if err != nil {
		log.Fatal(err)
	}
}
