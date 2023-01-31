package database

import (
	"fmt"

	"github.com/legaciespanda/location-tracker-api/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabaseAndRunMigration() {

	// databaseUser := godotenv.Load("DB_USER")
	// databasePassword := godotenv.Load("DB_PASSWORD")
	// databaseHost := godotenv.Load("DB_HOST")
	// databasePort := godotenv.Load("DB_PORT")
	// databaseName := godotenv.Load("DB_NAME")
	// databaseConfig := "?charset=utf8mb4&parseTime=True&loc=Local"

	dsn := "root:@tcp(127.0.0.1:3306)/iptracker_go?charset=utf8mb4&parseTime=True&loc=Local"
	//dsn := databaseUser + ":" + databasePassword + "@tcp(" + databaseHost + ":" + databasePort + ")/" + databaseName + databaseConfig
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	err = database.AutoMigrate(&model.User{}, &model.Track{}, &model.Location{}, &model.DeviceInfo{})
	if err != nil {
		return
	}

	fmt.Println("Migration sucessful...")

	DB = database
}
