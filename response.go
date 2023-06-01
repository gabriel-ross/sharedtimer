package sharedtimer

import (
	"encoding/json"
	"net/http"
	"time"
)

type Response struct {
	Timestamp time.Time `json:"timestamp"`
	Timer
}

func NewResponse(t Timer) Response {
	return Response{
		Timestamp: time.Now(),
		Timer:     t,
	}
}

func WriteJSON(w http.ResponseWriter, data interface{}, statusCode int) {
	out, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(out)
	w.WriteHeader(statusCode)
}
