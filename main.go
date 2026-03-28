package main

import ( 
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kishiEdward/backend_ginApp/config"
	"github.com/kishiEdward/backend_ginApp/routes"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("File .env tidak ditemukan, menggunakan environment variable sistem")
	}
	
	config.InitFirebase()
	config.InitDatabase()