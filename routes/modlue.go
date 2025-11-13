package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ali-hassan-Codes/file_analyzer_2/services"
)

type Router struct {
	Engine       *gin.Engine
	UserService  *services.SignupService
	LoginService *services.LoginService
	FileService  *services.FileAnalyzerService
}

// Constructor
func NewRouter(engine *gin.Engine, userService *services.SignupService, loginService *services.LoginService, fileService *services.FileAnalyzerService) *Router {
	return &Router{
		Engine:       engine,
		UserService:  userService,
		LoginService: loginService,
		FileService:  fileService,
	}
}
