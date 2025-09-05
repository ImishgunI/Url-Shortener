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
	Get(id int) (string, error)
	Update(original_url string, id int) error
	Delete(id int) error
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

func (d *DataBase) Get(id int) (string, error) {
	var url string
	err := d.db.QueryRow(context.Background(), "SELECT original_url FROM public.urls WHERE id=$1", id).Scan(&url)
	if err != nil {
		log.Printf("%v", err)
		return "", err
	}
	return url, nil
}

func (d *DataBase) Update(original_url string, id int) error {
	_, err := d.db.Query(context.Background(), "UPDATE urls SET original_url = $1 WHERE id = $2", original_url, id)
	if err != nil {
		log.Printf("%v", err)
		return err
	}
	return nil
}

func (d *DataBase) Delete(id int) error {
	_, err := d.db.Query(context.Background(), "DELETE FROM urls WHERE id=$1", id)
	if err != nil {
		log.Printf("%v", err)
		return err
	}
	return nil
}

func (d *DataBase) Close() {
	d.db.Close()
}
