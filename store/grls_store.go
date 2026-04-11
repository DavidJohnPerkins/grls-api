package store

import (
	"context"
)

type Model struct {
	Id               int    `db:"id"`
	Is_excluded      bool   `db:"is_excluded"`
	Sobriquet        string `db:"sobriquet"`
	Principal_name   string `db:"principal_name"`
	Hotness_quotient int    `db:"hotness_quotient"`
	Nationality      string `db:"nationality"`
	Ranking          string `db:"ranking"`
	Flags            string `db:"flags"`
	TH_url           string `db:"TH_url"`
	Movie_count      int    `db:"movie_count"`
}

type ModelExtended struct {
	Id               int    `db:"id"`
	Is_excluded      bool   `db:"is_excluded"`
	Sobriquet        string `db:"sobriquet"`
	Principal_name   string `db:"principal_name"`
	Aliases          string `db:"aliases"`
	Hotness_quotient int    `db:"hotness_quotient"`
	Ranking          string `db:"ranking"`
	Year_of_birth    string `db:"year_of_birth"`
	Nationality      string `db:"nationality"`
	Flags            string `db:"flags"`
	Comment          string `db:"comment"`
	Movie_count      int    `db:"movie_count"`
	TH_url           string `db:"TH_url"`
	RF_url           string `db:"RF_url"`
	FA_url           string `db:"FA_url"`
	BR_url           string `db:"BR_url"`
	PF_url           string `db:"PF_url"`
	PR_url           string `db:"PR_url"`
	AR_url           string `db:"AR_url"`
}

type Movie struct {
	Id           int    `db:"id"`
	Title        string `db:"title"`
	Comment      string `db:"comment"`
	Rating       string `db:"rating"`
	Participants int    `db:"participants"`
	Flags        string `db:"flags"`
	Names        string `db:"names"`
	Image_folder string `db:"image_folder"`
}

type Interface interface {
	GetModelList(ctx context.Context, term string) ([]Model, error)
	GetModel(ctx context.Context, id int) (ModelExtended, error)
	GetMovieList(ctx context.Context) ([]Movie, error)
}

// type Interface interface {
// 	GetModelList(ctx context.Context) ([]Model, error)
// 	GetByID(ctx context.Context, id uuid.UUID) (Movie, error)
// 	Create(ctx context.Context, jsonString string) error
// 	Update(ctx context.Context, jsonString string) error
// 	Delete(ctx context.Context, jsonString string) error
// }
