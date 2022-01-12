package rest

import (
	"fmt"
	"net/http"
	"strconv"
	"task/internal/entity/global"
	"task/internal/entity/producte"
	"task/internal/modules/product"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type Rest struct {
	server  *gin.Engine
	db      *sqlx.DB
	product product.Service
}

func NewRest(server *gin.Engine, db *sqlx.DB, product product.Service) *Rest {
	rest := &Rest{
		server:  server,
		db:      db,
		product: product,
	}

	server.GET("/products", rest.LoadAllProduct)
	server.GET("/product/:id", rest.LoadProductByID)
	server.POST("/product/add", rest.AddNewProduct)
	server.DELETE("/delete/:id", rest.DeleteProductById)

	return rest
}

func (r *Rest) Run() {
	r.server.Run(":8080")
}

func (r *Rest) LoadAllProduct(c *gin.Context) {
	tx, err := r.db.Beginx()
	if err != nil {
		errorMessage(c, err)
		return
	}
	defer tx.Rollback()

	productList, err := r.product.LoadAllProducts(tx)
	if err != nil {
		if err == global.ErrNoDataFound {
			errorMessage(c, fmt.Errorf("нет продуктов"))
			return
		}

		errorMessage(c, err)
		return
	}

	c.JSON(http.StatusOK, productList)

}

func (r *Rest) LoadProductByID(c *gin.Context) {
	tx, err := r.db.Beginx()
	if err != nil {
		errorMessage(c, err)
		return
	}
	defer tx.Rollback()

	idParam := c.Param("id")
	var id int
	id, err = strconv.Atoi(idParam)
	if err != nil {
		errorMessage(c, err)
		return
	}

	productFormList, err := r.product.LoadProductFormByID(tx, id)
	if err != nil {
		if err == global.ErrNoDataFound {
			errorMessage(c, fmt.Errorf("нет форм продуктов"))
			return
		}
		errorMessage(c, err)
		return
	}

	c.JSON(http.StatusOK, productFormList)
}

func (r *Rest) AddNewProduct(c *gin.Context) {
	tx, err := r.db.Beginx()
	if err != nil {
		errorMessage(c, err)
		return
	}
	defer tx.Rollback()

	var product producte.ProductForm
	err = c.BindJSON(&product)
	if err != nil {
		errorMessage(c, err)
		return
	}

	err = r.product.AddNewProduct(tx, product)

	if err != nil {
		errorMessage(c, err)
		return
	}

	c.JSON(http.StatusOK, "Successfully added")

}

func (r *Rest) DeleteProductById(c *gin.Context) {
	tx, err := r.db.Beginx()
	if err != nil {
		errorMessage(c, err)
		return
	}
	defer tx.Rollback()

	var id int

	idParam := c.Param("id")
	id, err = strconv.Atoi(idParam)
	if err != nil {
		errorMessage(c, err)
		return
	}

	err = r.product.DeleteProductById(tx, id)
	if err != nil {
		errorMessage(c, err)
		return
	}

	c.JSON(http.StatusOK, "Deletion was successful")
}

func errorMessage(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
}
