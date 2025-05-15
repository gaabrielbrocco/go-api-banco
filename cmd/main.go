package main

import (
	"database/sql"
	"fmt"
	"log"
	"teste/internal/infra/server"
	"teste/pkg/di"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func initConfig() {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath("/app")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading configuration file: %v", err)
	}
}

func runMigrations(db *sql.DB) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("Error creating migrate driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://config/database/migrations",
		"postgres", driver)
	if err != nil {
		log.Fatalf("Error starting migrate: %v", err)
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Error applying migrations: %v", err)
	}

	fmt.Println("Migrations applied successfully")
}

func main() {
	initConfig()

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		viper.GetString("DB_HOST"),
		viper.GetString("DB_PORT"),
		viper.GetString("DB_USER"),
		viper.GetString("DB_PASSWORD"),
		viper.GetString("DB_NAME"),
		viper.GetString("DB_SSL_MODE"),
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting database:", err)
	}
	defer db.Close()

	runMigrations(db)

	bancoController := di.NewBancoController(db)
	usuarioController := di.NewUsuarioController(db)
	contaController := di.NewContaController(db)

	port := viper.GetString("HTTP_PORT")
	srv := server.NewServer(bancoController, usuarioController, contaController, port)
	srv.ConfigureRoutes()

	if err := srv.Start(); err != nil {
		log.Fatal("Error starting app:", err)
	}

	fmt.Println("Run on port:", port)
}
