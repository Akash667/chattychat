package main

import (
	"github.com/akrawat667/baseChat/server/db"
	"github.com/akrawat667/baseChat/server/internal/user"
	"github.com/akrawat667/baseChat/server/internal/ws"
	"github.com/akrawat667/baseChat/server/router"
)

func main() {
	dbObj := db.NewDatabase().GetDB()
	repoObj := user.NewRepository(dbObj)
	serviceObj := user.NewService(repoObj)
	handlerObj := user.NewHandler(serviceObj)
	hub := ws.NewHub()
	wsHandlerObj := ws.NewHandler(hub)
	go hub.Run()
	router.InitRouter(handlerObj, wsHandlerObj)
	router.Start("localhost:8080")
}
