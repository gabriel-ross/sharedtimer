package main

import (
	"fmt"
	"time"

	"github.com/gabriel-ross/stgo"
)

func main() {
	t := stgo.NewCountdownTimer(stgo.CountdownTimerConfig{
		Hours:   1,
		Minutes: 10,
		Seconds: 10,
	})

	done := make(chan bool)
	go func() {
		t.Run()
		done <- true
	}()

	time.Sleep(2 * time.Second)
	fmt.Println(t.Remaining())
	time.Sleep(2 * time.Second)
	fmt.Println(t.Remaining())
	t.Pause()
	time.Sleep(2 * time.Second)
	fmt.Println(t.Remaining())
	t.Resume()
	time.Sleep(2 * time.Second)
	fmt.Println(t.Remaining())
	time.Sleep(time.Second)
	t.Restart()
	time.Sleep(time.Second)
	fmt.Println(t.Remaining())
	t.Cancel()

	// TODO: ensure restart is working
	// TODO: cancel is holding

	<-done
	fmt.Println("timer done")
}
