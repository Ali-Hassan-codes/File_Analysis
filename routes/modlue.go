package routes

import (
	"github.com/ali-hassan-Codes/file_analyzer_2/services"
	"github.com/ali-hassan-Codes/file_analyzer_2/ws"
	"github.com/gin-gonic/gin"
)

type Router struct {
	Engine       *gin.Engine
	UserService  services.ISignupService
	LoginService services.ILoginService
	FileService  services.IFileAnalyzerService
}

// Constructor
func NewRouter(
	engine *gin.Engine,
	userService services.ISignupService,
	loginService services.ILoginService,
	fileService services.IFileAnalyzerService,
) *Router {

	r := &Router{
		Engine:       engine,
		UserService:  userService,
		LoginService: loginService,
		FileService:  fileService,
	}

	// FIXED — this must be inside the function
	r.DefineRoutes()

	return r
}

// Example of using HubInstance in a handler
// Send message to a specific client
func (r *Router) NotifyClient(sessionID string, message string) {
	ws.HubInstance.SendToClient(sessionID, []byte(message))
}

// Example of broadcasting to all clients
func (r *Router) BroadcastMessage(message string) {
	ws.HubInstance.Broadcast <- []byte(message)
}
