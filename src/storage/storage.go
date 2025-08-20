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
	Save(original_url string) (int, error)
	Get(id int) string
}

type Closer interface {
	Close()
}

func DbConnect(cfg config.DBConfig) (*DataBase, error) {
	connString := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.SSLMode)
	conn, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		return nil, err
	}
	return &DataBase{db: conn}, nil
}

func (d *DataBase) Save(original_url string) (int, error) {
	_, err := d.db.Exec(context.Background(), "INSERT INTO urls (original_url) VALUES ($1)", original_url)
	if err != nil {
		return 0, err
	}
	var id int
	row := d.db.QueryRow(context.Background(), "SELECT id FROM urls WHERE original_urls=$1", original_url)
	row.Scan(&id)
	return id, nil
}

func (d *DataBase) Get(id int) string {
	var url string
	row := d.db.QueryRow(context.Background(), "SELECT original_url FROM urls WHERE id=$1", id)
	row.Scan(url)
	return url
}

func (d *DataBase) Close() {
	d.db.Close()
}
