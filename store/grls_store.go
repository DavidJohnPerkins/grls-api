package store

import (
	"context"
)

type Model struct {
	Id               int64  `db:"id"`
	Is_excluded      bool   `db:"is_excluded"`
	Sobriquet        string `db:"sobriquet"`
	Principal_name   string `db:"principal_name"`
	Hotness_quotient int    `db:"hotness_quotient"`
	Nationality      string `db:"nationality"`
	Flags            string `db:"flags"`
	TH_url           string `db:"TH_url"`
}

type ModelExtended struct {
	Id               int64  `db:"id"`
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

type Interface interface {
	GetModelList(ctx context.Context) ([]Model, error)
	GetModel(ctx context.Context, id int64) (ModelExtended, error)
}

// type Interface interface {
// 	GetModelList(ctx context.Context) ([]Model, error)
// 	GetByID(ctx context.Context, id uuid.UUID) (Movie, error)
// 	Create(ctx context.Context, jsonString string) error
// 	Update(ctx context.Context, jsonString string) error
// 	Delete(ctx context.Context, jsonString string) error
// }
