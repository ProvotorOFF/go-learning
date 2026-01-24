package order

import (
	"errors"
	"order-api-start/internal/product"
	"order-api-start/pkg/db"

	"gorm.io/gorm"
)

var Err = map[string]string{
	"INCORRECT_PRODUCT_IDS": "Incorrect product ids given",
}

type OrderRepository struct {
	Database *db.Db
}

func NewOrderRepository(database *db.Db) *OrderRepository {
	return &OrderRepository{database}
}

func (repo *OrderRepository) CreateFromRequest(req *OrderCreateRequest, id uint) (*Order, error) {
	var order Order

	err := repo.Database.Transaction(func(tx *gorm.DB) error {
		var products []product.Product

		if err := tx.Where("id in ?", req.Products).Find(&products).Error; err != nil {
			return err
		}

		if len(products) != len(req.Products) {
			return errors.New(Err["INCORRECT_PRODUCT_IDS"])
		}

		order = Order{
			UserID:   id,
			Products: products,
		}

		if err := tx.Create(&order).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (repo *OrderRepository) GetById(id uint64) (*Order, error) {
	var order Order
	if err := repo.Database.Preload("Products").First(&order, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func (repo *OrderRepository) FindForUser(id uint) ([]Order, error) {
	var orders []Order
	if err := repo.Database.Preload("Products").Find(&orders, "user_id = ?", id).Error; err != nil {
		return nil, err
	}
	return orders, nil
}
