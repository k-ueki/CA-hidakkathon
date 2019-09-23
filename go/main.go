package main

import (
	"hidakkathon/server"
	"log"
	"os"
)

func main() {
	log.SetFlags(log.Ldate + log.Ltime + log.Lshortfile)
	log.SetOutput(os.Stdout)

	s := server.NewServer()
	s.Init()
}
