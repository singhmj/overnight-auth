package db

import (
	"auth/db/models"
	"context"
	"database/sql"
)

type Store struct {
	*models.Queries
	db *sql.DB
}

func NewStore(sqlDB *sql.DB) Store {
	return Store{
		db:      sqlDB,
		Queries: models.New(sqlDB),
	}
}

func (store *Store) execTx(ctx context.Context, fn func(querie *models.Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	query := models.New(tx)
	err = fn(query)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return rbErr
		}

		return err
	}

	return tx.Commit()
}
