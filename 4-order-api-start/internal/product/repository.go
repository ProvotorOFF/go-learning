package product

import "order-api-start/pkg/db"

type ProductRepository struct {
	Database *db.Db
}

func NewProductRepository(database *db.Db) *ProductRepository {
	return &ProductRepository{database}
}

func (repo *ProductRepository) Create(product *Product) (*Product, error) {
	res := repo.Database.Create(&product)
	if res != nil {
		return nil, res.Error
	}
	return product, nil
}

func (repo *ProductRepository) GetById(id uint64) (*Product, error) {
	var product Product
	res := repo.Database.DB.First(&product, "id = ?", id)
	if res.Error != nil {
		return nil, res.Error
	}
	return &product, nil
}

func (repo *ProductRepository) All() ([]Product, error) {
	var products []Product
	res := repo.Database.DB.Find(&products)
	if res.Error != nil {
		return nil, res.Error
	}
	return products, nil
}

func (repo *ProductRepository) Update(product *Product) (*Product, error) {
	res := repo.Database.DB.Model(&Product{}).Where("id = ?", product.ID).Updates(product)
	if res.Error != nil {
		return nil, res.Error
	}
	return product, nil
}

func (repo *ProductRepository) Delete(id uint64) error {
	res := repo.Database.DB.Delete(&Product{}, "id = ?", id)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
