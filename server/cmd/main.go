package main

import (
	"github.com/akrawat667/baseChat/server/db"
	"github.com/akrawat667/baseChat/server/internal/user"
	"github.com/akrawat667/baseChat/server/router"
)

func main() {
	dbObj := db.NewDatabase().GetDB()
	repoObj := user.NewRepository(dbObj)
	serviceObj := user.NewService(repoObj)
	handlerObj := user.NewHandler(serviceObj)

	router.InitRouter(handlerObj)
	router.Start("localhost:8080")
}
