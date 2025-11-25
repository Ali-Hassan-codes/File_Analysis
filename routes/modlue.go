package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/ali-hassan-Codes/file_analyzer_2/services"
)

type Router struct {
    Engine       *gin.Engine
    UserService  services.ISignupService
    LoginService services.ILoginService
    FileService  services.IFileAnalyzerService
}

// Constructor
func NewRouter(engine *gin.Engine, userService services.ISignupService, loginService services.ILoginService, fileService services.IFileAnalyzerService) *Router {
    r := &Router{
        Engine:       engine,
        UserService:  userService,
        LoginService: loginService,
        FileService:  fileService,
    }

    r.DefineRoutes()
    return r
}
