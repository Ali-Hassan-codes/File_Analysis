package ws

type Hub struct {
	Clients    map[string]*Client // key = sessionID
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan []byte
}

func NewHub() *Hub {
	return &Hub{
		Clients:    make(map[string]*Client),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan []byte),
	}
}

// Global Hub instance
var HubInstance = NewHub()

// Hub main loop
func (h *Hub) Run() {
	for {
		select {

		case client := <-h.Register:
			h.Clients[client.ID] = client

		case client := <-h.Unregister:
			if _, ok := h.Clients[client.ID]; ok {
				delete(h.Clients, client.ID)
				close(client.Send)
			}

		case message := <-h.Broadcast:
			for _, client := range h.Clients {
				select {
				case client.Send <- message:
				default:
					delete(h.Clients, client.ID)
					close(client.Send)
				}
			}
		}
	}
}

// Send message to a specific client by sessionID
func (h *Hub) SendToClient(sessionID string, message []byte) {
	client, ok := h.Clients[sessionID]
	if ok {
		select {
		case client.Send <- message:
		default:
			delete(h.Clients, sessionID)
			close(client.Send)
		}
	}
}
