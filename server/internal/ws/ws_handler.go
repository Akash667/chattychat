package ws

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Handler struct {
	hub *Hub
}

func NewHandler(h *Hub) *Handler {
	return &Handler{hub: h}
}

type CreateRoomReq struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (h *Handler) CreateRoom(c *gin.Context) {
	var req CreateRoomReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}

	h.hub.Rooms[req.ID] = &Room{
		ID:      req.ID,
		Name:    req.Name,
		Clients: make(map[string]*Client)}
	log.Println(req.ID)
	log.Println(req.Name)
	c.JSON(http.StatusOK, req)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *Handler) JoinRoom(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error() + "upgrading to websocket failed"})
		return
	}
	roomID := c.Param("roomId")
	clientID := c.Query("userId")
	username := c.Query("username")

	cl := &Client{
		Conn:     conn,
		Message:  make(chan *Message, 10),
		ID:       clientID,
		RoomID:   roomID,
		Username: username,
	}

	m := &Message{
		Content:  "User has joined the room",
		RoomID:   roomID,
		Username: username,
	}
	h.hub.Register <- cl
	h.hub.Broadcast <- m

	go cl.writeMessage()
	cl.readMessage(h.hub)

}

type RoomRes struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (h *Handler) GetRooms(c *gin.Context) {

	rooms := []*RoomRes{}

	for _, r := range h.hub.Rooms {

		rooms = append(rooms, &RoomRes{
			ID:   r.ID,
			Name: r.Name,
		})
	}

	c.JSON(http.StatusOK, rooms)

}

type ClientRes struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (h *Handler) GetClients(c *gin.Context) {

	clients := []ClientRes{}
	roomID := c.Param("roomId")

	room, ok := h.hub.Rooms[roomID]
	if !ok {
		c.JSON(http.StatusNotFound, []ClientRes{})
		return
	}

	for _, r := range room.Clients {
		clients = append(clients, ClientRes{
			ID:   r.ID,
			Name: r.Username,
		})
	}

	c.JSON(http.StatusOK, clients)

}
