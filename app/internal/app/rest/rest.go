package rest

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"task/internal/entity/autorizatione"
	"task/internal/entity/global"
	"task/internal/entity/producte"
	"task/internal/modules/autorization"
	"task/internal/modules/product"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type Rest struct {
	server       *gin.Engine
	db           *sqlx.DB
	product      product.Service
	autorization autorization.Service
}

func NewRest(server *gin.Engine, db *sqlx.DB, product product.Service, autorization autorization.Service) *Rest {
	rest := &Rest{
		server:       server,
		db:           db,
		product:      product,
		autorization: autorization,
	}

	group := server.Group("/api")
	group.GET("/products", rest.Auth, rest.LoadAllProduct)
	group.GET("/product/:id", rest.Auth, rest.LoadProductByID)
	group.POST("/product/add", rest.Auth, rest.AddNewProduct)
	group.DELETE("/delete/:id", rest.Auth, rest.DeleteProductById)
	group.POST("/sign-up/regisration", rest.Register)
	group.POST("/sign-in/login", rest.Login)

	return rest
}

func (r *Rest) Run() {
	r.server.Run(":8080")
}

func (r *Rest) Auth(c *gin.Context) {
	token, exist := getTokenByCookie(c)
	if !exist {
		errorMessage(c, fmt.Errorf("авторизация провалилась"))
		c.Abort()
		return
	}

	err := compareToken(token)
	if err != nil {
		errorMessage(c, fmt.Errorf("авторизация провалилась"))
		c.Abort()
		return
	}
}

func (r *Rest) Register(c *gin.Context) {
	tx, err := r.db.Beginx()
	if err != nil {
		errorMessage(c, err)
		return
	}
	defer tx.Rollback()

	var userFromLoginForm autorizatione.User

	err = c.BindJSON(&userFromLoginForm)
	if err != nil {
		errorMessage(c, err)
		return
	}
	_, err = r.autorization.LoadUserByUsername(tx, userFromLoginForm.Username)

	switch err {
	case nil:
		errorMessage(c, fmt.Errorf("данный пользователь уже существует"))
		return

	case global.ErrNoDataFound:
		err = r.autorization.SaveUser(tx, userFromLoginForm)
		if err != nil {
			errorMessage(c, err)
			return
		}

		err := tx.Commit()
		if err != nil {
			fmt.Println("ошибка закрытия транзакции")
		}

		token, err := createToken(userFromLoginForm.Username)
		if err != nil {
			errorMessage(c, err)
			return
		}

		SetTokenInCookie(c, token)

		c.JSON(http.StatusOK, gin.H{"response": "registration completed successfully"})

	default:
		errorMessage(c, err)
		return
	}
}

func (r *Rest) Login(c *gin.Context) {
	tx, err := r.db.Beginx()
	if err != nil {
		errorMessage(c, err)
		return
	}
	defer tx.Rollback()

	var UserFromDB *autorizatione.User
	var userFromLoginForm autorizatione.User

	err = c.BindJSON(&userFromLoginForm)
	if err != nil {
		errorMessage(c, err)
		return
	}

	UserFromDB, err = r.autorization.LoadUserByUsername(tx, userFromLoginForm.Username)

	if err != nil {
		if err == global.ErrNoDataFound {
			errorMessage(c, fmt.Errorf("неверный логин или пароль"))
			return
		}
		errorMessage(c, err)
		return
	}

	if UserFromDB.Password == userFromLoginForm.Password {
		token, err := createToken(userFromLoginForm.Username)
		if err != nil {
			errorMessage(c, err)
			return
		}

		c.SetCookie(global.TOKEN, token, 60*60*24, "/", "", false, false)
		c.JSON(http.StatusOK, gin.H{"response": "login completed successfully"})
		return
	}
	errorMessage(c, fmt.Errorf("неверный логин или пароль"))
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

	if strings.Trim(product.Name, " ") == "" || strings.Trim(product.Form, " ") == "" {
		errorMessage(c, fmt.Errorf("поля не могут быть пустыми"))
		return
	}
	if product.Amount <= 0 {
		errorMessage(c, fmt.Errorf("колличество продукта не может равняться 0 или быть отрицательным числом"))
		return
	}
	if product.Price <= 0 {
		errorMessage(c, fmt.Errorf("цена продукта не может равняться 0 или быть отрицательным числом"))
		return
	}

	err = r.product.AddNewProduct(tx, product)

	if err != nil {
		errorMessage(c, err)
		return
	}

	err = tx.Commit()
	if err != nil {
		fmt.Println("ошибка закрытия транзакции")
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

	err = tx.Commit()
	if err != nil {
		fmt.Println("ошибка закрытия транзакции")
	}

	c.JSON(http.StatusOK, gin.H{"response": "Deletion was successful"})
}

func errorMessage(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
}

func createToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"username": username})

	return token.SignedString(global.JwtKey)
}

func SetTokenInCookie(c *gin.Context, token string) {
	c.SetCookie(global.TOKEN, token, 60*60*24*365, "/", "", false, true)
}

func getTokenByCookie(c *gin.Context) (token string, exists bool) {
	cookie, err := c.Request.Cookie(global.TOKEN)
	if err != nil {
		return
	}
	token = cookie.Value
	exists = token != ""
	return
}

func compareToken(token string) error {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return global.JwtKey, nil
	})
	if err != nil {
		return err
	}

	comparedTokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	comparedToken, err := comparedTokenClaims.SignedString(global.JwtKey)
	if err != nil {
		return err
	}

	if comparedToken != token {
		return errors.New("токены не равны")
	}
	return nil
}
