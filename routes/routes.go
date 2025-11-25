package routes

import (
	"github.com/ali-hassan-Codes/file_analyzer_2/middleware"

)

func (r *Router) DefineRoutes() {
	// Public routes
	r.Engine.POST("/signup", r.SignupHandler)
	r.Engine.POST("/login", r.LoginHandler)

	// Protected routes
	auth := r.Engine.Group("/")
	auth.Use(middleware.AuthMiddleware()) // ✅ middleware applied
	{
		auth.POST("/upload", r.FileAnalyzerHandler)
	}
}
