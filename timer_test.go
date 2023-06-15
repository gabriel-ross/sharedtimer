package sharedtimer

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	allowedError = 1
	hours        = 1
	minutes      = 1
	seconds      = 10
)

// Expected run time: ~6 seconds
func TestTimer(t *testing.T) {
	fmt.Printf("running timer test. expected time: %d seconds\n", 6)
	timer := NewTimer(TimerConfig{
		Hours:   hours,
		Minutes: minutes,
		Seconds: seconds,
	})
	secondsElapsed := 0

	go timer.Run()

	var _, _, secs int

	secondsElapsed++
	time.Sleep(time.Second)
	_, _, secs = timer.Remaining()
	assertApproximateEqual(t, seconds-secondsElapsed, secs)

	timer.Pause()
	time.Sleep(2 * time.Second)
	assertApproximateEqual(t, seconds-secondsElapsed, secs)

	go timer.Run()
	secondsElapsed += 2
	time.Sleep(2 * time.Second)
	_, _, secs = timer.Remaining()
	assertApproximateEqual(t, seconds-secondsElapsed, secs)

	timer.Restart()
	secondsElapsed = 1
	time.Sleep(time.Second)
	assertApproximateEqual(t, seconds-secondsElapsed, secs)

	timer.Cancel()
	assert.Equal(t, timer.InitialSeconds, timer.RemainingSeconds)
	assert.False(t, timer.IsRunning)
}

func assertApproximateEqual(t *testing.T, expected, actual int) bool {
	err := (expected - actual) / actual
	return assert.LessOrEqual(t, err, allowedError)
}
