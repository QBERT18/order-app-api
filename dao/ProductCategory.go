package dao

import (
	"fmt"

	"github.com/QBERT18/entity"
	"github.com/jinzhu/gorm"
)

type ProductCategoryDAO struct {
	DB *gorm.DB
}

var productCategoryDAOInstance *ProductCategoryDAO

func GetProductCategoryDAO() *ProductCategoryDAO {
	if productCategoryDAOInstance == nil {
		productCategoryDAOInstance = &ProductCategoryDAO{}
	}
	return productCategoryDAOInstance
}

func (dao *ProductCategoryDAO) FindAll() ([]entity.ProductCategory, error) {
	var categories []entity.ProductCategory
	err := dao.DB.Find(&categories).Error
	if err != nil {
		return nil, fmt.Errorf("error finding categories: %v", err)
	}
	return categories, nil
}

func (dao *ProductCategoryDAO) FindByID(id uint) (*entity.ProductCategory, error) {
	var category entity.ProductCategory
	err := dao.DB.First(&category, id).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, fmt.Errorf("category not found")
		}
		return nil, fmt.Errorf("error finding category: %v", err)
	}
	return &category, nil
}

func (dao *ProductCategoryDAO) Create(category *entity.ProductCategory) error {
	return dao.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(category).Error
		if err != nil {
			return fmt.Errorf("error creating category: %v", err)
		}
		return nil
	})
}

func (dao *ProductCategoryDAO) Update(category *entity.ProductCategory) error {
	return dao.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Save(category).Error
		if err != nil {
			return fmt.Errorf("error updating category: %v", err)
		}
		return nil
	})
}

func (dao *ProductCategoryDAO) Delete(category *entity.ProductCategory) error {
	return dao.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Delete(category).Error
		if err != nil {
			return fmt.Errorf("error deleting category: %v", err)
		}
		return nil
	})
}
