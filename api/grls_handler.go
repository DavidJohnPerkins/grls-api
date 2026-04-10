package api

import (
	"dperkins/grls-api/store"
	"errors"
	"net/http"
	"strconv"

	chi "github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type modelResponse struct {
	Id               int64  `json:"id"`
	Is_excluded      bool   `json:"is_excluded"`
	Sobriquet        string `json:"sobriquet"`
	Principal_name   string `json:"principal_name"`
	Hotness_quotient int    `json:"hotness_quotient"`
	Nationality      string `json:"nationality"`
	Flags            string `json:"flags"`
	TH_url           string `json:"th_url"`
}

type modelExtendedReponse struct {
	Id               int64  `json:"id"`
	Is_excluded      bool   `json:"is_excluded"`
	Sobriquet        string `json:"sobriquet"`
	Principal_name   string `json:"principal_name"`
	Aliases          string `json:"aliases"`
	Hotness_quotient int    `json:"hotness_quotient"`
	Ranking          string `json:"ranking"`
	Year_of_birth    string `json:"year_of_birth"`
	Nationality      string `json:"nationality"`
	Flags            string `json:"flags"`
	Comment          string `json:"comment"`
	Movie_count      int    `json:"movie_count"`
	TH_url           string `json:"TH_url"`
	RF_url           string `json:"RF_url"`
	FA_url           string `json:"FA_url"`
	BR_url           string `json:"BR_url"`
	PF_url           string `json:"PF_url"`
	PR_url           string `json:"PR_url"`
	AR_url           string `json:"AR_url"`
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

func NewModelExtendedResponse(m store.ModelExtended) modelExtendedReponse {
	return modelExtendedReponse{
		Id:               m.Id,
		Is_excluded:      m.Is_excluded,
		Sobriquet:        m.Sobriquet,
		Principal_name:   m.Principal_name,
		Aliases:          m.Aliases,
		Hotness_quotient: m.Hotness_quotient,
		Ranking:          m.Ranking,
		Year_of_birth:    m.Year_of_birth,
		Nationality:      m.Nationality,
		Flags:            m.Flags,
		Comment:          m.Comment,
		Movie_count:      m.Movie_count,
		TH_url:           m.TH_url,
		RF_url:           m.RF_url,
		FA_url:           m.FA_url,
		BR_url:           m.BR_url,
		PF_url:           m.PF_url,
		PR_url:           m.PR_url,
		AR_url:           m.AR_url,
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

func (mr modelExtendedReponse) Render(w http.ResponseWriter, r *http.Request) error {
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

func (s *Server) handleGetModel(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, _ := strconv.ParseInt(idParam, 10, 64)
	//id, err := strconv.Atoi(idParam)
	//if err != nil {
	//	render.Render(w, r, ErrBadRequest)
	//	return
	//}

	model, err := s.store.GetModel(r.Context(), id)
	if err != nil {
		var rnfErr *store.RecordNotFoundError
		if errors.As(err, &rnfErr) {
			render.Render(w, r, ErrRecordNotFound)
		} else {
			render.Render(w, r, ErrInternalServerError)
		}
		return
	}

	// Send the response and STOP
	render.Render(w, r, NewModelExtendedResponse(model))
}
