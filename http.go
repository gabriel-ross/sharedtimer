package stgo

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (svr *server) handleCreateTimer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		cnf := CountdownTimerConfig{}
		err = cnf.Bind(r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		t := NewCountdownTimer(cnf)
		svr.timers[t.Id] = &t

		WriteJSON(w, NewResponse(t), http.StatusOK)
	}
}

func (svr *server) handleListTimers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//
	}
}

func (svr *server) handleGetTimer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		id, err := uuid.Parse(chi.URLParam(r, "id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		timer, exists := svr.timers[id]
		if !exists {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		WriteJSON(w, NewResponse(*timer), http.StatusOK)
	}
}

func (svr *server) handleUpdateTimer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		id, err := uuid.Parse(chi.URLParam(r, "id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		timer, exists := svr.timers[id]
		if !exists {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		cnf := timer.cnf
		err = cnf.Bind(r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		t := NewCountdownTimer(cnf)
		t.Id = timer.Id
		svr.timers[t.Id] = &t

		WriteJSON(w, NewResponse(t), http.StatusOK)
	}
}

func (svr *server) handleDeleteTimer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		id, err := uuid.Parse(chi.URLParam(r, "id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if _, exists := svr.timers[id]; exists {
			delete(svr.timers, id)
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func (svr *server) handleStartTimer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (svr *server) handleCancelTimer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (svr *server) handleRestartTimer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (svr *server) handlePauseTimer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (svr *server) handleResumeTimer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
