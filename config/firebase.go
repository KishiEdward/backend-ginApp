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
	
}