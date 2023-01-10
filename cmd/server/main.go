package main

import (
	"log"

	"github.com/singlemancombat/proglog/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
  log.Println("Listening to localhost:8080 ...")
	log.Fatal(srv.ListenAndServe())
}
