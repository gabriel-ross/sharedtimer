package stgo

import "github.com/go-chi/chi"

func (svr server) Routes() chi.Router {
	r := chi.NewRouter()

	r.Route("/timers", func(r chi.Router) {
		r.Post("/", svr.handleCreateTimer())
		r.Get("/", svr.handleListTimers())

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", svr.handleGetTimer())
			r.Patch("/", svr.handleUpdateTimer())
			r.Delete("/", svr.handleDeleteTimer())

			r.Put("/start", svr.handleStartTimer())
			r.Put("/cancel", svr.handleCancelTimer())
			r.Put("/restart", svr.handleRestartTimer())
			r.Put("/pause", svr.handlePauseTimer())
			r.Put("/resume", svr.handleResumeTimer())
		})
	})

	return r
}
