package main

import (
	"fmt"
	"os"

	"github.com/gabriel-ross/sharedtimer"
)

var (
	PORT         = ""
	DEFAULT_PORT = "8080"
)

func main() {
	loadEnv()
	s := sharedtimer.NewServer(sharedtimer.ServerConfig{
		PORT: PORT,
	})
	fmt.Println(s.Run())
	s.OnShutdown()
}

func loadEnv() {
	if PORT = os.Getenv("PORT"); PORT == "" {
		PORT = DEFAULT_PORT
	}
}
