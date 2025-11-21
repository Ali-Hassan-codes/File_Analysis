package routes

func (r *Router) DefineRoutes() {
	r.Engine.POST("/signup", r.SignupHandler)
	r.Engine.POST("/login", r.LoginHandler)
	r.Engine.POST("/upload", r.FileAnalyzerHandler) 
}
