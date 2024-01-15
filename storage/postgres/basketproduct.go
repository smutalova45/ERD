package postgres

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"main.go/models"
	"main.go/storage"
)

type basketproductRepo struct {
	db *sql.DB
}

func NewBasketProduct(db *sql.DB) storage.IBasketProductStorage {
	return basketproductRepo{
		db: db,
	}
}

func (bp basketproductRepo) Create(b models.CreateBP) (models.BasketProducts, error) {
	id := uuid.New()
	if _, err := bp.db.Exec(`insert into basketproducts values($1,$2,$3,$4)`, id, b.BasketID, b.ProductID, b.Quantity); err != nil {
		fmt.Println("error while inserting basket products", err.Error())
		return models.BasketProducts{}, err
	}
	return models.BasketProducts{}, nil

}
func (bp basketproductRepo) GetById(id models.PKBasketProducts) (models.BasketProducts, error) {
	bps := models.BasketProducts{}
	if err := bp.db.QueryRow(`select id,basketid,productid,quantity where id=$1`, id.ID).Scan(
		&bps.ID,
		&bps.BasketID,
		&bps.ProductID,
		bps.Quantity,
	); err != nil {
		return models.BasketProducts{}, err
	}
	return models.BasketProducts{}, nil

}
func (bp basketproductRepo) GetListRequest(models.GetAllRequestBP) (models.BasketProductsResponse, error) {
	rows, err := bp.db.Query(`select * from basketproducts`)
	if err != nil {
		return models.BasketProductsResponse{}, err
	}
	bps := []models.BasketProducts{}
	for rows.Next() {
		bpss := models.BasketProducts{}
		if err = rows.Scan(&bpss.ID, &bpss.BasketID, &bpss.ProductID, &bpss.Quantity); err != nil {
			return models.BasketProductsResponse{}, err
		}
		bps = append(bps, bpss)
	}
	response := models.BasketProductsResponse{
		BasketProducts: bps,
		Count:          len(bps),
	}
	return response, nil

}
func (bp basketproductRepo) Update(bps models.UpdateBP) (models.BasketProducts, error) {
	_, err := bp.db.Exec(`update basketproducts set quantity=$1 where id=$2 `, bps.Quantity, bps.ID)
	if err != nil {
		return models.BasketProducts{}, err
	}
	updatedbasketproduct, err := bp.GetById(models.PKBasketProducts{ID: bps.ID})
	if err != nil {
		return models.BasketProducts{}, err
	}
	return updatedbasketproduct, nil

}
func (bp basketproductRepo) Delete(id models.PKBasketProducts) error {

	if _, err := bp.db.Exec(`delete from basketproducts where id=$1`, id.ID); err != nil {
		return err
	}
	return nil
}
