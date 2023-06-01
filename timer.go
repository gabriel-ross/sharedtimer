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
	Paused           bool      `json:"paused"`
	cnf              TimerConfig
	cancelC          chan bool
	restartC         chan bool
	pauseC           chan bool
	resumeC          chan bool
}

func NewTimer(cnf TimerConfig) Timer {
	return Timer{
		Id:               uuid.New(),
		Name:             cnf.Name,
		InitialSeconds:   secondsFromClock(cnf.Hours, cnf.Minutes, cnf.Seconds),
		RemainingSeconds: secondsFromClock(cnf.Hours, cnf.Minutes, cnf.Seconds),
		Paused:           false,
		cnf:              cnf,
		cancelC:          make(chan bool),
		restartC:         make(chan bool),
		pauseC:           make(chan bool),
		resumeC:          make(chan bool),
	}
}

func (t *Timer) init() {
	t.cancelC = make(chan bool)
	t.restartC = make(chan bool)
	t.pauseC = make(chan bool)
	t.resumeC = make(chan bool)
}

func (t *Timer) Run() {
	t.init()
	t.Paused = false

	tick := time.NewTicker(time.Second)
	for t.RemainingSeconds > 0 {
		select {
		case <-t.cancelC:
			return
		case <-t.restartC:
			t.RemainingSeconds = t.InitialSeconds
		case <-t.pauseC:
			<-t.resumeC
		case <-t.resumeC: // dump redundant resume calls
		case <-tick.C:
			t.RemainingSeconds--
		}
	}
}

func (t *Timer) Cancel() {
	t.unpause()
	t.cancelC <- true
}

func (t *Timer) Restart() {
	t.unpause()
	t.restartC <- true
}

func (t *Timer) Pause() {
	t.Paused = true
	t.pauseC <- true
}

func (t *Timer) Resume() {
	t.unpause()
}

func (t *Timer) unpause() {
	if t.Paused {
		t.resumeC <- true
		t.Paused = false
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
