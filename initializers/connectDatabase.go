package initializers

import (
	"os"

	"github.com/kwul0208/go-restapi-gin/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() *gorm.DB {
	dns := os.Getenv("DB")
	database, err := gorm.Open(mysql.Open(dns))
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&models.Product{})
	database.AutoMigrate(&models.Users{})

	DB = database

	return database
}
