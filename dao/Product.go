package dao

import (
	"fmt"

	"github.com/QBERT18/entity"
	"github.com/jinzhu/gorm"
)

type ProductDAO struct {
	DB *gorm.DB
}

var productDAOInstance *ProductDAO

func GetProductDAO() *ProductDAO {
	if productDAOInstance == nil {
		productDAOInstance = &ProductDAO{}
	}
	return productDAOInstance
}

func (dao *ProductDAO) FindAll() ([]entity.Product, error) {
	var products []entity.Product
	err := dao.DB.Preload("Category").Find(&products).Error
	if err != nil {
		return nil, fmt.Errorf("error finding products: %v", err)
	}
	return products, nil
}

func (dao *ProductDAO) FindByID(id uint) (*entity.Product, error) {
	var product entity.Product
	err := dao.DB.Preload("Category").First(&product, id).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, fmt.Errorf("product not found")
		}
		return nil, fmt.Errorf("error finding product: %v", err)
	}
	return &product, nil
}

func (dao *ProductDAO) Create(product *entity.Product) error {
	return dao.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(product).Error
		if err != nil {
			return fmt.Errorf("error creating product: %v", err)
		}
		return nil
	})
}

func (dao *ProductDAO) Update(product *entity.Product) error {
	return dao.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Save(product).Error
		if err != nil {
			return fmt.Errorf("error updating product: %v", err)
		}
		return nil
	})
}

func (dao *ProductDAO) Delete(product *entity.Product) error {
	return dao.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Delete(product).Error
		if err != nil {
			return fmt.Errorf("error deleting product: %v", err)
		}
		return nil
	})
}
