package api

import (
	"dperkins/grls-api/store"
	"net/http"

	"github.com/go-chi/render"
)

type modelResponse struct {
	Id               int64  `json:"id"`
	Is_excluded      bool   `json:"is_excluded"`
	Sobriquet        string `json:"sobriquet"`
	Principal_name   string `json:"principal_name"`
	Hotness_quotient int64  `json:"hotness_quotient"`
	Nationality      string `json:"nationality"`
	Flags            string `json:"flags"`
	TH_url           string `json:"th_url"`
}

func NewModelResponse(m store.Model) modelResponse {
	return modelResponse{
		Id:               m.Id,
		Is_excluded:      m.Is_excluded,
		Sobriquet:        m.Sobriquet,
		Principal_name:   m.Principal_name,
		Hotness_quotient: m.Hotness_quotient,
		Nationality:      m.Nationality,
		Flags:            m.Flags,
		TH_url:           m.TH_url,
	}
}

func NewModelListResponse(models []store.Model) []render.Renderer {
	list := []render.Renderer{}
	for _, model := range models {
		mr := NewModelResponse(model)
		list = append(list, mr)
	}
	return list
}

func (mr modelResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *Server) handleModelList(w http.ResponseWriter, r *http.Request) {
	models, err := s.store.GetModelList(r.Context())
	if err != nil {
		render.Render(w, r, ErrInternalServerError)
		return
	}

	render.RenderList(w, r, NewModelListResponse(models))
}
