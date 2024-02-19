package main

import (
	"net/http"

	"github.com/joho/godotenv"
	"github.com/Nihaochingiz/go-quest/controllers"
	"github.com/Nihaochingiz/go-quest/models"
)


func main() {
  godotenv.Load()

  handler := controllers.New() 

  server := &http.Server{
		Addr:    "0.0.0.0:8008",
		Handler: handler,
  }

  models.ConnectDatabase()

  server.ListenAndServe()
}