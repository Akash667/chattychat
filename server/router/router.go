package router

import (
	"github.com/akrawat667/baseChat/server/internal/user"
	"github.com/akrawat667/baseChat/server/internal/ws"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func InitRouter(userHandler *user.Handler, wsHandler *ws.Handler) {
	r = gin.Default()
	r.Use(cors.Default())
	r.Use(gin.Logger())
	gin.SetMode(gin.DebugMode)
	r.POST("/signup", userHandler.CreateUser)
	r.POST("/login", userHandler.LoginUser)
	r.GET("/logout", userHandler.Logout)
	r.POST("/ws/createRoom", wsHandler.CreateRoom)
	r.GET("/ws/joinRoom/:roomId", wsHandler.JoinRoom)
	r.GET("/ws/getRooms/", wsHandler.GetRooms)
	r.GET("/ws/getClients/:roomId", wsHandler.GetClients)

}

func Start(addr string) error {
	return r.Run(addr)
}
