package storage

import (
	"context"
	"fmt"
	"log"
	"log/slog"
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
	connString := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?search_path=public&sslmode=%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.SSLMode)
	conn, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		log.Printf("%v", err)
		return nil, err
	}
	slog.Info("Succesful connection to DataBase")
	return &DataBase{db: conn}, nil
}

func (d *DataBase) Save(original_url string) (int, error) {
	var id int
	err := d.db.QueryRow(context.Background(), "INSERT INTO public.urls (original_url) VALUES ($1) RETURNING id", original_url).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (d *DataBase) Get(id int) string {
	var url string
	row := d.db.QueryRow(context.Background(), "SELECT original_url FROM public.urls WHERE id=$1", id)
	row.Scan(url)
	return url
}

func (d *DataBase) Close() {
	d.db.Close()
}
