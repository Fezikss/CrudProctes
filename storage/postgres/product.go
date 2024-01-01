package postgres

import (
	"database/sql"
	"github.com/google/uuid"
	"main/models"
)

type ProductRepo struct {
	Db *sql.DB
}

func NewProductRepo(db *sql.DB) ProductRepo {
	return ProductRepo{Db: db}
}

func (p ProductRepo) Insert(product models.Product) (string, error) {
	id := uuid.New()

	if _, err := p.Db.Exec(`insert into product values ($1,$2,$3,$4)`, id, product.Name, product.Price, product.CategoryID); err != nil {
		return "", err
	}

	return id.String(), nil
}

func (p ProductRepo) GetById(id uuid.UUID) (models.Product, error) {
	pro := models.Product{}

	if err := p.Db.QueryRow(`select id, name, price , category_id, created_at,updated_at from product where id = $1`, id).Scan(
		&pro.Id,
		&pro.Name,
		&pro.Price,
		&pro.CategoryID,
		&pro.CreatedAt,
		&pro.UpdatedAt,
	); err != nil {
		return models.Product{}, err
	}

	return pro, nil
}

func (p ProductRepo) GetList() ([]models.Product, error) {

	products := []models.Product{}

	rows, err := p.Db.Query(`Select *from product`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		product := models.Product{}

		err := rows.Scan(
			&product.Id,
			&product.Name,
			&product.Price,
			&product.CategoryID,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil

}

func (p ProductRepo) Update(product models.Product) error {

	if _, err := p.Db.Exec(`UPDATE product set name = $1, price = $2, category_id = $3 where Id = $4`,
		product.Name, product.Price, product.CategoryID, product.Id); err != nil {
		return err
	}

	return nil

}

func (p ProductRepo) Delete(id uuid.UUID) error {

	if _, err := p.Db.Exec(`DELETE from product where id = $1`, id); err != nil {
		return err
	}
	return nil
}
