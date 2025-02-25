package main

import (
	"notes/internal/service"
	"notes/logs"

	"github.com/labstack/echo/v4"
)

func main() {
	logger := logs.NewLogger(false)
	db, err := PostgresConnection()
	if err != nil {
		logger.Fatal(err)
	}
	svc := service.NewService(db, logger)

	router := echo.New()
	api := router.Group("api")

	//Notes
	api.GET("/note/:id", svc.GetNoteById, svc.JWTCheck)
	api.GET("/note", svc.GetNotes, svc.JWTCheck)
	api.POST("/notes", svc.CreateNewNote, svc.JWTCheck)
	api.DELETE("/note/:id", svc.DeleteNote, svc.JWTCheck)
	api.PUT("/note/:id", svc.UpdateNote, svc.JWTCheck)

	//Users
	api.POST("/registration", svc.Registration)
	api.POST("/login", svc.Login)

	router.Logger.Fatal(router.Start(":8000"))
}
