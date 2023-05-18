package stgo

import "github.com/go-chi/chi"

func (svr *server) Routes() chi.Router {
	r := chi.NewRouter()

	r.Route("/timers", func(r chi.Router) {
		r.Post("/", svr.handleCreateTimer())
		r.Get("/", svr.handleListTimers())

		r.Route("/{id}", func(r chi.Router) {
			r.Patch("/", svr.handleUpdateTimer())
			r.Delete("/", svr.handleDeleteTimer())

			r.Get("/", svr.handleGetTimer())
			r.Post("/start", svr.handleStartTimer())
			r.Post("/cancel", svr.handleCancelTimer())
			r.Post("/restart", svr.handleRestartTimer())
			r.Post("/pause", svr.handlePauseTimer())
			r.Post("/resume", svr.handleResumeTimer())
		})
	})

	return r
}
