package routes

import (
	"net/http"
	"strings"

	"github.com/ali-hassan-Codes/file_analyzer_2/ws"
	"github.com/gin-gonic/gin"
)

// FileAnalyzerHandler handles file upload, analysis, and WebSocket notification
func (r *Router) FileAnalyzerHandler(c *gin.Context) {
	// Extract session ID from Authorization header
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing or invalid"})
		return
	}
	sessionID := strings.TrimPrefix(authHeader, "Bearer ")

	// Get the uploaded file
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file is uploaded"})
		return
	}

	// Save the file temporarily
	filename := "./tmppp/" + file.Filename
	if err := c.SaveUploadedFile(file, filename); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	// Call the service to analyze the file
	result, err := r.FileService.AnalyzeFile(filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Send WebSocket message to the specific client
	go ws.HubInstance.SendToClient(sessionID, []byte("File "+file.Filename+" analyzed successfully!"))

	// Return JSON response
	c.JSON(http.StatusOK, result)
}
