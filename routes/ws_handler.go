package routes

import (
	"log"
	"net/http"

	"github.com/ali-hassan-Codes/file_analyzer_2/ws"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

// WebSocketHandler upgrades connection and registers client into Hub
func (r *Router) WebSocketHandler(c *gin.Context) {
    sessionID := c.Query("session_id")
    if sessionID == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "session_id is required"})
        return
    }

    conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
    if err != nil {
        log.Println("WebSocket upgrade failed:", err)
        return
    }

    client := &ws.Client{
        ID:   sessionID,
        Hub:  ws.HubInstance,
        Conn: conn,
        Send: make(chan []byte, 256),
    }

    // Register client
    ws.HubInstance.Register <- client
    log.Println("Client connected with session ID:", sessionID)

    // MUST RUN BOTH
    go client.WritePump()
    go client.ReadPump()
}
