package postgres

import (
	"database/sql"

	"github.com/google/uuid"
	"main.go/models"
	"main.go/storage"
)

type categoryRepo struct {
	db *sql.DB
}

func NewCategory(db *sql.DB) storage.ICategory {
	return categoryRepo{
		db: db,
	}
}

func (c categoryRepo) Create(createcategory models.CreateCategory) (models.Category, error) {
	id := uuid.New()
	if _, err := c.db.Exec(`insert into category values($1,$2)`, id, createcategory.Name); err != nil {
		return models.Category{}, err
	}
	return models.Category{}, nil

}
func (c categoryRepo) GetByID(id models.PKCategory) (models.Category, error) {
	category := models.Category{}
	if err := c.db.QueryRow(`select id, name_ where id=$1`, id.ID).Scan(
		&category.ID,
		&category.Name,
	); err != nil {
		return models.Category{}, err
	}
	return models.Category{}, nil
}

func (c categoryRepo) GetListRequest(models.GetAllRequestCategory) (models.CategoryResponse, error) {
	rows, err := c.db.Query(`select * from category`)
	if err != nil {
		return models.CategoryResponse{}, err
	}
	categories := []models.Category{}
	for rows.Next() {
		category := models.Category{}
		if err := rows.Scan(&category.ID, &category.Name); err != nil {
			return models.CategoryResponse{}, err
		}
		categories = append(categories, category)
	}
	response := models.CategoryResponse{
		Categories: categories,
		Count:      len(categories),
	}
	return response, nil

}
func (c categoryRepo) Update(category models.Category) (models.Category, error) {
	_, err := c.db.Exec(`update category set name_=$1 where id=$2`, category.Name, category.ID)
	if err != nil {
		return models.Category{}, err
	}
	updatedcategory, err := c.GetByID(models.PKCategory{ID: category.ID})
	if err != nil {
		return models.Category{}, err
	}
	return updatedcategory, nil

}
func (c categoryRepo) Delete(id models.PKCategory) error {
	if _, err := c.db.Exec(`delete from products where id=$1`, id.ID); err != nil {
		return err
	}
	if _, err := c.db.Exec(`delete from category where id=$1`, id.ID); err != nil {
		return err
	}
	return nil

}
