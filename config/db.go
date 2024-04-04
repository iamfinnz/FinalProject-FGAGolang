package config

// Import library
import (
	"fmt"
	"log"
	"mygram/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Informasi db
var (
	host     = "localhost"
	user     = "postgres"
	password = "ilham"
	dbPort   = "5432"
	dbName   = "mygram"
	db       *gorm.DB
	err      error
)

// Function untuk membuat koneksi ke database dan auto migration
func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbName, dbPort)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("Error saat koneksi ke database :", err)
	}
	defer fmt.Println("Berhasil koneksi ke database")

	db.Debug().AutoMigrate(models.User{}, models.Comment{}, models.Photo{}, models.SocialMedia{})
}

// Function GetDB untuk mendapatkan instance dari objek database yang sudah terkoneksi
func GetDB() *gorm.DB {
	return db
}
