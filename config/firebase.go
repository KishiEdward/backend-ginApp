package config

import (
	"fmt"
	"log"
	"os"

	"github.com/KishiEdward/backend_ginApp/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDatabase() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbname)

	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Log semua query SQL [cite: 602]
	}

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), gormConfig) // [cite: 605]
	if err != nil {
		log.Fatalf("Gagal koneksi ke database: %v", err) // [cite: 608]
	}

	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Gagal mendapatkan sql.DB: %v", err) // [cite: 616]
	}
	sqlDB.SetMaxOpenConns(25) // [cite: 618]
	sqlDB.SetMaxIdleConns(10)
}