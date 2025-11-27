package ws

import (
	"github.com/gorilla/websocket"
)

// Each WebSocket client
type Client struct {
	ID   string   
	Hub  *Hub
	Conn *websocket.Conn
	Send chan []byte
}

// Listen for outgoing messages
func (c *Client) WritePump() {
	for msg := range c.Send {
		c.Conn.WriteMessage(websocket.TextMessage, msg)
	}
}
