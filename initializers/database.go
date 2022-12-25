package initializers

import (
	"log"
	"time"

	"github.com/broemp/brettspielTrackerAPI/entity"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	dsn := buildConnectionString()
	counter := 1
connection:
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println("Failed to connect to Database: ", err)
		time.Sleep(time.Duration(counter) * time.Second)
		counter += 2
		goto connection
	}
	DB.AutoMigrate(&entity.Boardgame{}, &entity.Collection{})
}

func buildConnectionString() string {
	dbIP := viper.GetString("dbIP")
	dbUsername := viper.GetString("dbUsername")
	dbPassword := viper.GetString("dbPassword")
	dbName := viper.GetString("dbName")
	dbPort := viper.GetString("dbPort")

	return "host=" + dbIP + " user=" + dbUsername + " password=" + dbPassword + " dbname=" + dbName + " port=" + dbPort + " sslmode=disable TimeZone=Europe/Berlin"
}
