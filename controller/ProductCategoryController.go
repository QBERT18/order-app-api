package controller

import (
	"net/http"
	"strconv"

	"github.com/QBERT18/dao"
	"github.com/QBERT18/entity"
	"github.com/gin-gonic/gin"
)

type ProductCategoryController struct {
	DAO *dao.ProductCategoryDAO
}

func (c *ProductCategoryController) FindAll(ctx *gin.Context) {
	categories, err := c.DAO.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, categories)
}

func (c *ProductCategoryController) FindByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid category ID"})
		return
	}
	category, err := c.DAO.FindByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if category == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "category not found"})
		return
	}
	ctx.JSON(http.StatusOK, category)
}

func (c *ProductCategoryController) Create(ctx *gin.Context) {
	var category entity.ProductCategory
	err := ctx.BindJSON(&category)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = c.DAO.Create(&category)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, category)
}

func (c *ProductCategoryController) Update(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid category ID"})
		return
	}
	var category entity.ProductCategory
	err = ctx.BindJSON(&category)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	category.ID = uint(id)
	err = c.DAO.Update(&category)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, category)
}

func (c *ProductCategoryController) Delete(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid category ID"})
		return
	}
	err = c.DAO.Delete(&entity.ProductCategory{ID: uint(id)})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "category deleted successfully"})
}
