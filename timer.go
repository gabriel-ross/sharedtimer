package sharedtimer

import (
	"time"

	"github.com/google/uuid"
)

const (
	secondsPerMinute = 60
	secondsPerHour   = 60 * secondsPerMinute
)

type TimerConfig struct {
	Name    string `json:"name"`
	Hours   int64  `json:"hours"`
	Minutes int64  `json:"minutes"`
	Seconds int64  `json:"seconds"`
}

type Timer struct {
	Id               uuid.UUID `json:"id"`
	Name             string    `json:"name"`
	InitialSeconds   int64     `json:"initialSeconds"`
	RemainingSeconds int64     `json:"remainingSeconds"`
	IsRunning        bool      `json:"isRunning"`
	cnf              TimerConfig
	stopC            chan bool
}

func NewTimer(cnf TimerConfig) Timer {
	return Timer{
		Id:               uuid.New(),
		Name:             cnf.Name,
		InitialSeconds:   secondsFromClock(cnf.Hours, cnf.Minutes, cnf.Seconds),
		RemainingSeconds: secondsFromClock(cnf.Hours, cnf.Minutes, cnf.Seconds),
		IsRunning:        false,
		cnf:              cnf,
		stopC:            make(chan bool),
	}
}

func (t *Timer) init() {
	t.stopC = make(chan bool)
}

func (t *Timer) Run() {
	t.init()
	t.IsRunning = true

	tick := time.NewTicker(time.Second)
	for t.RemainingSeconds > 0 {
		select {
		case <-t.stopC:
			return
		case <-tick.C:
			t.RemainingSeconds--
		}
	}
}

func (t *Timer) Cancel() {
	if t.IsRunning {
		t.stopC <- true
	}
	t.RemainingSeconds = t.InitialSeconds
}

func (t *Timer) Restart() {
	t.RemainingSeconds = t.InitialSeconds
	if t.IsRunning {
		t.stopC <- true
		t.Run()
	}
}

func (t *Timer) Pause() {
	if t.IsRunning {
		t.stopC <- true
		t.IsRunning = false
	}
}

func (t *Timer) Resume() {
	if !t.IsRunning {
		t.Run()
	}
}

func (t *Timer) Remaining() (hour, min, sec int) {
	return clockFromSeconds(t.RemainingSeconds)
}

func secondsFromClock(hours, minutes, seconds int64) int64 {
	return hours*secondsPerHour + minutes*secondsPerMinute + seconds
}

func clockFromSeconds(seconds int64) (hours, minutes, secs int) {
	secs = int(seconds)
	hours = secs / secondsPerHour
	secs -= hours * secondsPerHour
	minutes = secs / secondsPerMinute
	secs -= minutes * secondsPerMinute
	return
}
