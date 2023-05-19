package main

import (
	"fmt"
	"os"

	"github.com/gabriel-ross/stgo"
)

var (
	PORT         = ""
	DEFAULT_PORT = "8080"
)

func main() {
	loadEnv()
	s := stgo.NewServer(stgo.ServerConfig{
		PORT: PORT,
	})
	fmt.Println(s.Run())
}

func loadEnv() {
	if PORT = os.Getenv("PORT"); PORT == "" {
		PORT = DEFAULT_PORT
	}
}
