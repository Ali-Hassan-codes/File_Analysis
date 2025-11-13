package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// FileAnalyzerHandler handles file upload and analysis
func (r *Router) FileAnalyzerHandler(c *gin.Context) {
	// Get the uploaded file
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file is uploaded"})
		return
	}

	// Save the file temporarily
	filename := "./tmp/" + file.Filename
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

	// Return JSON response
	c.JSON(http.StatusOK, result)
}
