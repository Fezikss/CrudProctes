package postgres

import (
	"database/sql"
	"github.com/google/uuid"
	"main/models"
)

type CategoryRepo struct {
	DB *sql.DB
}

func NewCategoryRepo(db *sql.DB) CategoryRepo {
	return CategoryRepo{DB: db}
}

func (c CategoryRepo) Insert(category models.Category) (string, error) {
	id := uuid.New()

	if _, err := c.DB.Exec(`insert into category values ($1,$2)`, id, category.Name); err != nil {
		return "", err
	}
	return id.String(), nil
}

func (c CategoryRepo) GetById(id uuid.UUID) (models.Category, error) {
	category := models.Category{}
	c.DB.QueryRow(`SELECT *from category`).Scan(
		&category.Id,
		&category.Name,
		&category.CreatedAt,
		&category.UpdatedAt,
	)
	return category, nil

}

func (c CategoryRepo) GetList() ([]models.Category, error) {
	ctgList := []models.Category{}

	rows, err := c.DB.Query(`SELECT *FROM category`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		ctg := models.Category{}

		rows.Scan(
			&ctg.Id,
			&ctg.Name,
			&ctg.CreatedAt,
			&ctg.UpdatedAt,
		)

		ctgList = append(ctgList, ctg)
	}

	return ctgList, nil
}

func (c CategoryRepo) Update(category models.Category) error {
	if _, err := c.DB.Exec(`UPDATE category set name = $1 where id = $2`, category.Name, category.Id); err != nil {
		return err
	}
	return nil
}

func (c CategoryRepo) Delete(id uuid.UUID) error {

	if _, err := c.DB.Exec(`delete from category where id = $1`, id); err != nil {
		return err
	}

	return nil
}
