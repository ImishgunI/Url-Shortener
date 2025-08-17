package storage

import (
	"context"
	"fmt"
	"shortener/src/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DataBase struct {
	db *pgxpool.Pool
}

type Repository interface {
	Save(original_url string, db *DataBase) (int, error)
	Get(id int, db *DataBase) (string, error)
}

func DbConnect(cfg config.DBConfig) (*DataBase, error) {
	connString := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.SSLMode)
	conn, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		return nil, err
	}
	return &DataBase{db: conn}, nil
}
