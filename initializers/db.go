package initializers

import (
	"log"
	"os"

	"github.com/russelldias98/go-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	dsn := os.Getenv("DB_URL")
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database")
	}

	DB = db

	log.Println("Connected to database")

	Migrate()
}

func Migrate() {
	DB.AutoMigrate(&models.Supplier{}, &models.Product{}, &models.Product{})
}
