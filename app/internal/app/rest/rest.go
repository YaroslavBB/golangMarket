package rest

import (
	"fmt"
	"net/http"
	"strconv"
	"task/internal/entity/autorisatione"
	"task/internal/entity/global"
	"task/internal/entity/producte"
	"task/internal/modules/autorisation"
	"task/internal/modules/product"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type Rest struct {
	server       *gin.Engine
	db           *sqlx.DB
	product      product.Service
	autorisation autorisation.Service
}

func NewRest(server *gin.Engine, db *sqlx.DB, product product.Service, autorisation autorisation.Service) *Rest {
	rest := &Rest{
		server:       server,
		db:           db,
		product:      product,
		autorisation: autorisation,
	}

	server.GET("/products", rest.LoadAllProduct)
	server.GET("/product/:id", rest.LoadProductByID)
	server.POST("/product/add", rest.AddNewProduct)
	server.DELETE("/delete/:id", rest.DeleteProductById)
	server.POST("/sign-up/regisration", rest.Register)

	return rest
}

func (r *Rest) Run() {
	r.server.Run(":8080")
}

func (r *Rest) Register(c *gin.Context) {
	tx, err := r.db.Beginx()
	if err != nil {
		errorMessage(c, err)
		return
	}
	defer tx.Rollback()

	var userFromLoginForm autorisatione.User

	err = c.BindJSON(&userFromLoginForm)
	fmt.Println(userFromLoginForm)
	if err != nil {
		errorMessage(c, err)
		return
	}

	_, err = r.autorisation.LoadUserByUsername(tx, userFromLoginForm.Username)

	switch err {
	case nil:
		errorMessage(c, fmt.Errorf("данный пользователь уже существует"))
		return

	case global.ErrNoDataFound:
		err = r.autorisation.SaveUser(tx, userFromLoginForm)
		if err != nil {
			errorMessage(c, err)
			return
		}
		c.JSON(http.StatusOK, gin.H{"response": "registration completed successfully"})
		return

	default:
		errorMessage(c, err)
		return
	}
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

	c.JSON(http.StatusOK, gin.H{"response": "Successfully added"})

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

	c.JSON(http.StatusOK, gin.H{"response": "Deletion was successful"})
}

func errorMessage(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
}
