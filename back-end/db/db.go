package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joaogabrielvianna/entity"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDataBase() {

	err := godotenv.Load("./.env") // Exemplo: "./config/.env"
	if err != nil {
		log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	}
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	fmt.Println(dsn, "teste")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Erro ao conectar ao banco:", err)
	}

	DB = db
	fmt.Println("✅ Banco de dados conectado!")

	db.AutoMigrate(
		&entity.Course{},
		&entity.Year{},
		&entity.User{},
	)

}
