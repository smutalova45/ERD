package postgres

import (
	"database/sql"

	"github.com/google/uuid"
	"main.go/models"
	"main.go/storage"
)

type basketRepo struct {
	db *sql.DB
}

func NewBasket(db *sql.DB) storage.IBasketsStorage {
	return basketRepo{
		db: db,
	}
}
func (b basketRepo) Create(createbasket models.CreateBaskets) (models.Baskets, error) {
	id := uuid.New()
	if _, err := b.db.Exec(`insert into basket values($1,$2,$3)`, id, createbasket.CustomerID, createbasket.TotalSum); err != nil {
		return models.Baskets{}, err
	}
	return models.Baskets{}, nil
}
func (b basketRepo) GetById(id models.PKBaskets) (models.Baskets, error) {
	basket := models.Baskets{}
	if err := b.db.QueryRow(`select id,customerid,totalsum where id=$1`, id.ID).Scan(
		&basket.ID,
		&basket.CustomerID,
		&basket.TotalSum,
	); err != nil {
		return models.Baskets{}, err
	}
	return models.Baskets{}, nil

}
func (b basketRepo) GetListRequest(models.GetALLRequestBasket) (models.BasketResponse, error) {
	rows, err := b.db.Query(`select * from basket`)
	if err != nil {
		return models.BasketResponse{}, err
	}
	baskets := []models.Baskets{}
	for rows.Next() {
		basket := models.Baskets{}
		if err = rows.Scan(&basket.ID, &basket.CustomerID, &basket.TotalSum); err != nil {
			return models.BasketResponse{}, err
		}
		baskets = append(baskets, basket)
	}
	response := models.BasketResponse{
		Baskets: baskets,
		Count:   len(baskets),
	}
	return response, nil
}
func (b basketRepo) Update(basket models.UpdateBaskets) (models.Baskets, error) {
	_, err := b.db.Exec(`update basket set totalsum=$1 where id=$2`, basket.TotalSum, basket.ID)
	if err != nil {
		return models.Baskets{}, err
	}
	updatedbasket, err := b.GetById(models.PKBaskets{ID: basket.ID})
	if err != nil {
		return models.Baskets{}, err
	}
	return updatedbasket, nil
}

func (b basketRepo) Delete(id models.PKBaskets) error {
	if _, err := b.db.Exec(`delete from basketproducts where id=$1`, id.ID); err != nil {
		return err
	}
	if _, err := b.db.Exec(`delete from basket where id=$1`, id.ID); err != nil {
		return err
	}
	return nil

}
