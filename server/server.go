package server

import (
	"log"

	"github.com/ali-hassan-Codes/file_analyzer_2/db"
	"github.com/ali-hassan-Codes/file_analyzer_2/repositories"
	"github.com/ali-hassan-Codes/file_analyzer_2/routes"
	"github.com/ali-hassan-Codes/file_analyzer_2/services"
	"github.com/gin-gonic/gin"
	"github.com/ali-hassan-Codes/file_analyzer_2/ws"

)

func StartServer() {
	go ws.HubInstance.Run()
	database := db.InitDb()

	// User repository
	userRepo := repositories.NewUserRepository(database)

	// Signup service - pass dependency struct, returns interface
	userService := services.NewSignupService(services.SignupServiceDeps{
		Repo: userRepo,
	})

	// Login service (use same pattern if using interface)
	loginService := services.NewLoginService(services.LoginServiceDeps{
		Repo: userRepo,
	})

	// File Analyzer service
	fileRepo := repositories.NewFileAnalyzerRepository(database)
	fileService := services.NewFileAnalyzerService(services.FileAnalyzerServiceDeps{
		Repo: fileRepo,
	})

	engine := gin.Default()

	// Router now accepts interfaces
	routes.NewRouter(engine, userService, loginService, fileService)

	log.Println("✅ Server started on http://localhost:8001")
	engine.Run(":8001")
}
