package store

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/microsoft/go-mssqldb"
)

const driverName = "sqlserver"

type SqlServerGrlsStore struct {
	databaseUrl string
	dbx         *sqlx.DB
}

func NewSqlServerGrlsStore(databaseUrl string) *SqlServerGrlsStore {
	return &SqlServerGrlsStore{
		databaseUrl: databaseUrl,
	}
}

func noOpMapper(s string) string {
	return s
}

func (s *SqlServerGrlsStore) connect(ctx context.Context) error {
	dbx, err := sqlx.ConnectContext(ctx, driverName, s.databaseUrl)
	if err != nil {
		log.Printf("DB connect failed: %v", err)
		return err
	}

	dbx.MapperFunc(noOpMapper)
	s.dbx = dbx
	return nil
}

func (s *SqlServerGrlsStore) close() error {
	return s.dbx.Close()
}

func (s *SqlServerGrlsStore) GetModelList(ctx context.Context) ([]Model, error) {
	err := s.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer s.close()

	var models []Model
	r, err := s.dbx.QueryxContext(
		ctx, `
		EXEC GRLS.r_model_card_list '{"model_id": -1}';
	`)

	if err != nil {
		return nil, err
	}
	defer r.Close()

	for r.Next() {
		var m Model
		if err := r.StructScan(&m); err != nil {
			log.Printf("failed: %v", err)
			return nil, err
		}
		models = append(models, m)
	}

	return models, nil
}

func (s *SqlServerGrlsStore) GetModel(ctx context.Context, id int64) (ModelExtended, error) {
	err := s.connect(ctx)
	if err != nil {
		return ModelExtended{}, err
	}
	defer s.close()

	var model ModelExtended
	jsonBody := fmt.Sprintf(`{"model_id": %d}`, id)

	r, err := s.dbx.QueryxContext(
		ctx, `
		EXEC GRLS.r_model @p_input_json = @json`,
		sql.Named("json", jsonBody))

	if err != nil {
		return ModelExtended{}, err
	}
	defer r.Close()

	if r.Next() {
		if err := r.StructScan(&model); err != nil {
			log.Printf("err1: %v", err)

			return ModelExtended{}, err
		}
	} else {
		//return ModelExtended{}, sql.ErrNoRows
		return ModelExtended{}, &RecordNotFoundError{id}
	}
	return model, nil
}
