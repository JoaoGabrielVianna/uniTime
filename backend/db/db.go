package db

import (
	"fmt"
	"os"

	"github.com/joaogabrielvianna/config"
	"github.com/joaogabrielvianna/entity"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDataBase() error {
	logger := config.GetLogger("PostgreSQL")

	err := godotenv.Load("./.env") // Exemplo: "./config/.env"
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Erro ao carregar o arquivo .env: %v", err))
	}
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Erro ao conectar ao banco de dados: %v", err))
	}

	DB = db
	logger.SuccessLog("Banco de dados conectado!")

	db.AutoMigrate(
		&entity.Course{},
		&entity.Calendar{},
		&entity.CalendarEvent{},
		&entity.User{},
	)

	// Resetar o contador do auto-incremento
	err = ResetAutoIncrement(db) // Passando o db do GORM
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Erro ao resetar o auto-incremento: %v", err))
		return err
	}
	return nil
}

func ResetAutoIncrement(db *gorm.DB) error {
	// Verificar se a sequência existe antes de tentar resetá-la
	var exists bool
	queryCheck := "SELECT EXISTS (SELECT 1 FROM pg_class WHERE relname = 'users_id_seq')"
	db.Raw(queryCheck).Scan(&exists)

	if !exists {
		return fmt.Errorf("a sequência users_id_seq não existe")
	}

	query := "ALTER SEQUENCE users_id_seq RESTART WITH 1"
	result := db.Exec(query)
	if result.Error != nil {
		return fmt.Errorf("erro ao resetar o auto-incremento: %v", result.Error)
	}
	return nil
}
