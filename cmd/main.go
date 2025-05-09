package main

import (
	"database/sql"
	"fmt"
	"log"
	"teste/internals/core/usecase"
	"teste/internals/infra/repository"
	"teste/internals/infra/server"

	"github.com/spf13/viper"
	_ "github.com/lib/pq"
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
		log.Fatal("Erro ao conectar no banco de dados:", err)
	}
	defer db.Close()

	bancoRepository := repository.NewBancoRepository(db)
	bancoUseCase := usecase.NewBancoUseCase(bancoRepository)

	port := viper.GetString("HTTP_PORT")
	srv := server.NewServer(bancoUseCase, port)
	srv.ConfigureRoutes()

	if err := srv.Start(); err != nil {
		log.Fatal("Erro ao iniciar o servidor:", err)
	}

	fmt.Println("Rodando na porta:", port)
}
