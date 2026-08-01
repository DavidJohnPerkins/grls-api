package api

import (
	chi "github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func (s *Server) routes() {
	s.router.Use(render.SetContentType(render.ContentTypeJSON))
	s.router.Route("/api/grls", func(r chi.Router) {

		// /api/grls/model?id=123
		r.Get("/model", s.handleModelList)

		// /api/grls/model/get?id=123
		r.Get("/model/get", s.handleGetModel)

		// /api/grls/model/associates?id=123
		r.Get("/model/associates", s.handleGetModelAssociates)

		// /api/grls/modelsearch?term=abc
		r.Get("/modelsearch", s.handleModelList)

		// /api/grls/model/create   (POST)
		r.Post("/model/create", s.handleCreateMovie)

		// /api/grls/movies?model_id=123
		r.Get("/movies", s.handleMovieList)

		// /api/grls/contactsheet?images=["aaliyah.jpg","aislin.jpg"]
		r.Get("/contactsheet", s.handleContactSheet)

		// /api/grls/add/flags
		r.Get("/add/flags", s.handleFlagList)

		// /api/grls/add/attr?attr_abbrev=xyz
		r.Get("/add/attr", s.handleAttrDescList)
	})

	// s.router.Route("/api/grls", func(r chi.Router) {
	// 	r.Get("/model", s.handleModelList)
	// 	r.Route("/model/{id}", func(r chi.Router) {
	// 		r.Get("/", s.handleGetModel)
	// 	})
	// 	r.Route("/model/associates/{id}", func(r chi.Router) {
	// 		r.Get("/", s.handleGetModelAssociates)
	// 	})
	// 	r.Route("/modelsearch/{term}", func(r chi.Router) {
	// 		r.Get("/", s.handleModelList)
	// 	})
	// 	r.Route("/model/create", func(r chi.Router) {
	// 		r.Post("/", s.handleCreateMovie)
	// 	})
	// 	//r.Get("/movies/", s.handleMovieList)
	// 	r.Route("/movies/{model_id}", func(r chi.Router) {
	// 		r.Get("/", s.handleMovieList)
	// 	})

	// 	r.Get("/contactsheet", s.handleContactSheet)

	// 	r.Get("/add/flags", s.handleFlagList)

	// 	r.Route("/add/attr/{attr_abbrev}", func(r chi.Router) {
	// 		r.Get("/", s.handleAttrDescList)
	// 	})
	// })
}
