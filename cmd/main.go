package main

import (
	"log"
	"net/http"
	"os"
	"taskify/internal/database"
	"taskify/internal/models"
	"taskify/internal/router"

	"github.com/joho/godotenv"
)

func main() {
	// Carregar variáveis de ambiente
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  Arquivo .env não encontrado, usando variáveis do sistema")
	}

	// Conectar ao banco
	database.Connect()

	// Migrar tabelas
	database.DB.AutoMigrate(&models.Task{}, &models.User{})

	// Iniciar servidor
	r := router.NewRouter()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 Servidor rodando em http://localhost:%s\n", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal("❌ Erro ao iniciar servidor:", err)
	}
}
