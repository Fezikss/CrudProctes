package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"main/config"
)

type Store struct {
	Db              *sql.DB
	ProductStorege  ProductRepo
	CategoryStorege CategoryRepo
}

func New(cfg config.Config) (Store, error) {
	url := fmt.Sprintf(`host=%s port=%s user=%s password=%s database=%s sslmode=disable`, cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDB)
	db, err := sql.Open("postgres", url)
	if err != nil {
		return Store{}, err
	}

	productRepo := NewProductRepo(db)
	categoryRepo := NewCategoryRepo(db)

	return Store{
		Db:              db,
		ProductStorege:  productRepo,
		CategoryStorege: categoryRepo,
	}, nil
}
