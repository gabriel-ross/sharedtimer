package stgo

import "time"

const (
	secondsPerMinute = 60
	secondsPerHour   = 60 * secondsPerMinute
)

type CountdownTimerConfig struct {
	Hours   int64
	Minutes int64
	Seconds int64
}

type countdownTimer struct {
	cnf              CountdownTimerConfig
	InitialSeconds   int64
	RemainingSeconds int64
	cancelC          chan bool
	restartC         chan bool
	pauseC           chan bool
	resumeC          chan bool
	paused           bool
}

func NewCountdownTimer(cnf CountdownTimerConfig) countdownTimer {
	return countdownTimer{
		cnf:              cnf,
		InitialSeconds:   secondsFromClock(cnf.Hours, cnf.Minutes, cnf.Seconds),
		RemainingSeconds: secondsFromClock(cnf.Hours, cnf.Minutes, cnf.Seconds),
		cancelC:          make(chan bool),
		restartC:         make(chan bool),
		pauseC:           make(chan bool),
		resumeC:          make(chan bool),
		paused:           false,
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
	t.paused = true
	t.pauseC <- true
}

func (t *countdownTimer) Resume() {
	t.unpause()
}

func (t *countdownTimer) unpause() {
	if t.paused {
		t.resumeC <- true
		t.paused = false
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
