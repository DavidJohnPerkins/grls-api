package store

import (
	"context"
)

type Model struct {
	Id               int64  `db:"id"`
	Is_excluded      bool   `db:"is_excluded"`
	Sobriquet        string `db:"sobriquet"`
	Principal_name   string `db:"principal_name"`
	Hotness_quotient int64  `db:"hotness_quotient"`
	Nationality      string `db:"nationality"`
	Flags            string `db:"flags"`
	TH_url           string `db:"TH_url"`
}

type Interface interface {
	GetModelList(ctx context.Context) ([]Model, error)
}

// type Interface interface {
// 	GetModelList(ctx context.Context) ([]Model, error)
// 	GetByID(ctx context.Context, id uuid.UUID) (Movie, error)
// 	Create(ctx context.Context, jsonString string) error
// 	Update(ctx context.Context, jsonString string) error
// 	Delete(ctx context.Context, jsonString string) error
// }
