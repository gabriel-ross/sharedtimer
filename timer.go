package stgo

import (
	"time"

	"github.com/google/uuid"
)

const (
	secondsPerMinute = 60
	secondsPerHour   = 60 * secondsPerMinute
)

type CountdownTimerConfig struct {
	Name    string `json:"name"`
	Hours   int64  `json:"hours"`
	Minutes int64  `json:"minutes"`
	Seconds int64  `json:"seconds"`
}

type countdownTimer struct {
	Id               uuid.UUID `json:"id"`
	Name             string    `json:"name"`
	InitialSeconds   int64     `json:"initialSeconds"`
	RemainingSeconds int64     `json:"remainingSeconds"`
	Paused           bool      `json:"paused"`
	cnf              CountdownTimerConfig
	cancelC          chan bool
	restartC         chan bool
	pauseC           chan bool
	resumeC          chan bool
}

func NewCountdownTimer(cnf CountdownTimerConfig) countdownTimer {
	return countdownTimer{
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

func LoadCountdownTimer(t countdownTimer) countdownTimer {
	return countdownTimer{
		Id:               t.Id,
		Name:             t.Name,
		InitialSeconds:   t.InitialSeconds,
		RemainingSeconds: t.RemainingSeconds,
		Paused:           t.Paused,
		cnf:              t.cnf,
		cancelC:          make(chan bool),
		restartC:         make(chan bool),
		pauseC:           make(chan bool),
		resumeC:          make(chan bool),
	}
}

func (t *countdownTimer) Run() {
	tick := time.NewTicker(time.Second)
	for {
		select {
		case <-t.cancelC:
			return
		case <-t.restartC:
			t.RemainingSeconds = t.InitialSeconds
		case <-t.pauseC:
			<-t.resumeC
		case <-t.resumeC: // dump redundant resume calls
		case <-tick.C:
			if t.RemainingSeconds > 0 {
				t.RemainingSeconds--
			} else {
				return
			}
		}
	}
}

func (t *countdownTimer) Cancel() {
	t.unpause()
	t.cancelC <- true
}

func (t *countdownTimer) Restart() {
	t.unpause()
	t.restartC <- true
}

func (t *countdownTimer) Pause() {
	t.Paused = true
	t.pauseC <- true
}

func (t *countdownTimer) Resume() {
	t.unpause()
}

func (t *countdownTimer) unpause() {
	if t.Paused {
		t.resumeC <- true
		t.Paused = false
	}
}

func (t *countdownTimer) Remaining() (hour, min, sec int) {
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
