package server

import (
	"log"

	"github.com/ali-hassan-Codes/file_analyzer_2/db"
	"github.com/ali-hassan-Codes/file_analyzer_2/repositories"
	"github.com/ali-hassan-Codes/file_analyzer_2/routes"
	"github.com/ali-hassan-Codes/file_analyzer_2/services"
	"github.com/gin-gonic/gin"
)

func StartServer() {
	database := db.InitDb()

	userRepo := repositories.NewUserRepository(database)
	userService := services.NewSignupService(userRepo)

	// Login service
	loginService := services.NewLoginService(userRepo)

	// File Analyzer service
	fileRepo := repositories.NewFileAnalyzerRepository(database)
	fileService := services.NewFileAnalyzerService(fileRepo)

	engine := gin.Default()

	routes.NewRouter(engine, userService, loginService, fileService)

	log.Println("✅ Server started on http://localhost:8001")
	engine.Run(":8001")
}
