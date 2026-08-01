package api

import (
	"dperkins/grls-api/store"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	chi "github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type modelResponse struct {
	Id               int    `json:"id"`
	Is_excluded      bool   `json:"is_excluded"`
	Sobriquet        string `json:"sobriquet"`
	Principal_name   string `json:"principal_name"`
	Hotness_quotient int    `json:"hotness_quotient"`
	Nationality      string `json:"nationality"`
	Ranking          string `json:"ranking"`
	Flags            string `json:"flags"`
	TH_url           string `json:"th_url"`
	Movie_count      int    `json:"movie_count"`
}

type modelExtendedReponse struct {
	Id               int    `json:"id"`
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

type modelAssociateReponse struct {
	Id             int    `json:"id"`
	Sobriquet      string `json:"sobriquet"`
	Principal_name string `json:"principal_name"`
}

type movieResponse struct {
	Id           int    `json:"id"`
	Title        string `json:"title"`
	Comment      string `json:"comment"`
	Rating       string `json:"rating"`
	Participants int    `json:"participants"`
	Flags        string `json:"flags"`
	Names        string `json:"names"`
	Image_folder string `json:"image_folder"`
}

type attrDescResponse struct {
	L2_desc string `json:"l2_desc"`
}

type flagResponse struct {
	Flag_abbrev string `json:"flag_abbrev"`
}

type contactSheetResponse struct {
	Image_name string `json:"image_name"`
	Model_id   int    `json:"model_id"`
}

func NewModelResponse(m store.Model) modelResponse {
	return modelResponse{
		Id:               m.Id,
		Is_excluded:      m.Is_excluded,
		Sobriquet:        m.Sobriquet,
		Principal_name:   m.Principal_name,
		Hotness_quotient: m.Hotness_quotient,
		Nationality:      m.Nationality,
		Ranking:          m.Ranking,
		Flags:            m.Flags,
		TH_url:           m.TH_url,
		Movie_count:      m.Movie_count,
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

func NewModelAssociateResponse(m store.ModelAssociate) modelAssociateReponse {
	return modelAssociateReponse{
		Id:             m.Id,
		Sobriquet:      m.Sobriquet,
		Principal_name: m.Principal_name,
	}
}

func NewMovieResponse(m store.Movie) movieResponse {
	return movieResponse{
		Id:           m.Id,
		Title:        m.Title,
		Comment:      m.Comment,
		Rating:       m.Rating,
		Participants: m.Participants,
		Flags:        m.Flags,
		Names:        m.Names,
		Image_folder: m.Image_folder,
	}
}

func NewAttrDescResponse(m store.AttrDesc) attrDescResponse {
	return attrDescResponse{
		L2_desc: m.L2_desc,
	}
}

func NewFlagResponse(m store.Flag) flagResponse {
	return flagResponse{
		Flag_abbrev: m.Flag_abbrev,
	}
}

func NewContactSheetResponse(m store.ContactSheet) contactSheetResponse {
	return contactSheetResponse{
		Image_name: m.Image_name,
		Model_id:   m.Model_id,
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

func NewModelAssociateListResponse(associates []store.ModelAssociate) []render.Renderer {

	list := []render.Renderer{}
	for _, associate := range associates {
		mr := NewModelAssociateResponse(associate)
		list = append(list, mr)
	}
	return list
}

func NewMovieListResponse(movies []store.Movie) []render.Renderer {

	list := []render.Renderer{}
	for _, movie := range movies {
		mr := NewMovieResponse(movie)
		list = append(list, mr)
	}
	return list
}

func NewAttrDescListResponse(desc []store.AttrDesc) []render.Renderer {

	list := []render.Renderer{}
	for _, d := range desc {
		mr := NewAttrDescResponse(d)
		list = append(list, mr)
	}
	return list
}

func NewFlagListResponse(desc []store.Flag) []render.Renderer {

	list := []render.Renderer{}
	for _, d := range desc {
		mr := NewFlagResponse(d)
		list = append(list, mr)
	}
	return list
}

func NewContactSheetListResponse(cs []store.ContactSheet) []render.Renderer {

	list := []render.Renderer{}
	for _, d := range cs {
		mr := NewContactSheetResponse(d)
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

func (mr modelAssociateReponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (mr movieResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (mr attrDescResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (mr flagResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (mr contactSheetResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *Server) handleModelList(w http.ResponseWriter, r *http.Request) {

	termParam := r.URL.Query().Get("term")
	if termParam == "" {
		termParam = "%"
	} else {
		termParam = strings.Replace(termParam, "~", "%", -1)
	}

	flagsParam := r.URL.Query().Get("flags")
	fmt.Printf("flagsParam: %v", flagsParam)

	flags := []string{}
	if flagsParam != "" {
		json.Unmarshal([]byte(flagsParam), &flags)
	}

	combined := map[string]interface{}{
		"search_term":  termParam,
		"search_flags": flags,
	}

	combinedJSON, err := json.Marshal(combined)
	if err != nil {
		log.Println("error marshaling combined JSON:", err)
	}

	fmt.Println("combinedJSON", string(combinedJSON))

	models, err := s.store.GetModelList(r.Context(), string(combinedJSON))
	if err != nil {
		render.Render(w, r, ErrInternalServerError)
		return
	}

	render.RenderList(w, r, NewModelListResponse(models))
}

func (s *Server) handleGetModel(w http.ResponseWriter, r *http.Request) {

	idParam := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idParam)

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

	render.Render(w, r, NewModelExtendedResponse(model))
}

func (s *Server) handleGetModelAssociates(w http.ResponseWriter, r *http.Request) {

	idParam := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idParam)

	associates, err := s.store.GetModelAssociateList(r.Context(), id)
	if err != nil {
		render.Render(w, r, ErrInternalServerError)
		return
	}

	render.RenderList(w, r, NewModelAssociateListResponse(associates))
}

func (s *Server) handleMovieList(w http.ResponseWriter, r *http.Request) {

	idParam := chi.URLParam(r, "model_id")
	model_id, _ := strconv.Atoi(idParam)

	movies, err := s.store.GetMovieList(r.Context(), model_id)
	if err != nil {
		render.Render(w, r, ErrInternalServerError)
		return
	}

	render.RenderList(w, r, NewMovieListResponse(movies))
}

func (s *Server) handleAttrDescList(w http.ResponseWriter, r *http.Request) {

	termParam := r.URL.Query().Get("attr_abbrev")

	desc, err := s.store.GetAttrDescList(r.Context(), termParam)
	if err != nil {
		render.Render(w, r, ErrInternalServerError)
		return
	}

	render.RenderList(w, r, NewAttrDescListResponse(desc))
}

func (s *Server) handleFlagList(w http.ResponseWriter, r *http.Request) {

	typeParam := r.URL.Query().Get("type")
	desc, err := s.store.GetFlagList(r.Context(), typeParam)
	if err != nil {
		render.Render(w, r, ErrInternalServerError)
		return
	}

	render.RenderList(w, r, NewFlagListResponse(desc))
}

func (s *Server) handleContactSheet(w http.ResponseWriter, r *http.Request) {

	imageParam := r.URL.Query().Get("images")
	i, err := s.store.GetContactSheet(r.Context(), imageParam)
	if err != nil {
		render.Render(w, r, ErrInternalServerError)
		return
	}

	render.RenderList(w, r, NewContactSheetListResponse(i))
}

func (s *Server) handleCreateMovie(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}
	jsonString := string(body)

	if err := s.store.CreateModel(r.Context(), jsonString); err != nil {
		// DB or SP failure → 500, not 400
		log.Printf("err: %v", err)
		render.Render(w, r, ErrInternalServerError)
		return
	}

	// Success
	render.Status(r, http.StatusCreated)
	render.JSON(w, r, map[string]string{"status": "ok"})

}
