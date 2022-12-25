package main

import (
	"github.com/QBERT18/controller"
	"github.com/QBERT18/dao"
	"github.com/QBERT18/db"
	"github.com/QBERT18/entity"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db := db.GetMySQLDB()
	defer db.Close()

	db.AutoMigrate(&entity.Product{}, &entity.ProductCategory{})

	r := gin.New()

	productDAO := dao.ProductDAO{DB: db}
	productCategoryDAO := dao.ProductCategoryDAO{DB: db}

	productController := controller.ProductController{DAO: &productDAO}
	productCategoryController := controller.ProductCategoryController{DAO: &productCategoryDAO}

	r.GET("/products", productController.FindAll)
	r.GET("/products/:id", productController.FindByID)
	r.POST("/products", productController.Create)
	r.PUT("/products/:id", productController.Update)
	r.DELETE("/products/:id", productController.Delete)

	r.GET("/categories", productCategoryController.FindAll)
	r.GET("/categories/:id", productCategoryController.FindByID)
	r.POST("/categories", productCategoryController.Create)
	r.PUT("/categories/:id", productCategoryController.Update)
	r.DELETE("/categories/:id", productCategoryController.Delete)

	r.Run(":8080")
}
