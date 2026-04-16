package api

import (
	chi "github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func (s *Server) routes() {
	s.router.Use(render.SetContentType(render.ContentTypeJSON))

	//s.router.Get("/health", s.handleGetHealth)

	// s.router.Route("/api/movies", func(r chi.Router) {
	// 	r.Get("/", s.handleListMovies)
	// 	r.Post("/", s.handleCreateMovie)
	// 	r.Delete("/", s.handleDeleteMovie)
	// 	r.Route("/{id}", func(r chi.Router) {
	// 		r.Get("/", s.handleGetMovie)
	// 		r.Put("/", s.handleUpdateMovie)
	// 	})
	// })
	s.router.Route("/api/grls", func(r chi.Router) {
		r.Get("/", s.handleModelList)
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", s.handleGetModel)
		})
		r.Route("/search/{term}", func(r chi.Router) {
			r.Get("/", s.handleModelList)
		})

		//r.Get("/movies/", s.handleMovieList)
		r.Route("/movies/{model_id}", func(r chi.Router) {
			r.Get("/", s.handleMovieList)
		})

		r.Get("/add/flags/{flag_type}", s.handleFlagList)

		r.Route("/add/attr/{attr_abbrev}", func(r chi.Router) {
			r.Get("/", s.handleAttrDescList)
		})
	})
}
