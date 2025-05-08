package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"teste/internals/core/usecase"
	"teste/internals/infra/repository"
	"teste/internals/infra/server"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		getEnv("DB_HOST", "db"),
		getEnv("DB_PORT", "5432"),
		getEnv("DB_USER", "postgres"),
		getEnv("DB_PASSWORD", "postgres"),
		getEnv("DB_NAME", "banco"),
		getEnv("DB_SSL_MODE", "disable"),
	)

	fmt.Println("Conectando no banco de dados:", connStr)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to database", err)
	}

	defer db.Close()

	bancoRepository := repository.NewBancoRepository(db)
	bancoUseCase := usecase.NewBancoUseCase(bancoRepository)

	port := getEnv("HTTP_PORT", "8080")
	srv := server.NewServer(bancoUseCase, port)
	srv.ConfigureRoutes()

	if err := srv.Start(); err != nil {
		log.Fatal("Error starting server ", err)
	}
	fmt.Println("Rodando na porta: ", port)

}
