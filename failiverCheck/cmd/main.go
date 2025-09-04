package main

import (
	"failiverCheck/internal/api"
	"log"
)

func main() {
	log.Println("App start")
	api.StartServer()
	log.Println("App terminated")
}
