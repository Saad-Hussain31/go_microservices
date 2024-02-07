// Entry point of the application. It initializes the database client, creates an instance of the server, and starts the server.

package main

import (
	"log"

	"github.com/Saad-Hussain31/go-microservices/internal/database"
	"github.com/Saad-Hussain31/go-microservices/internal/server"
)

func main() {
	db, err := database.NewDatabaseClient()
	if err != nil {
		log.Fatalf("failed to initialize Database Client: %s", err)
	}
	srv := server.NewEchoServer(db)
	if err := srv.Start(); err != nil {
		log.Fatal(err.Error())
	}
}
