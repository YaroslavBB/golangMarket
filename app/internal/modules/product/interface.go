package product

import (
	"task/internal/entity/producte"

	"github.com/jmoiron/sqlx"
)

// Repository по product
type Repository interface {
	LoadAllProducts(tx *sqlx.Tx) ([]producte.Product, error)
	LoadProductFormByID(tx *sqlx.Tx, id int) ([]producte.ProductForm, error)
	UpdateProductAmount(tx *sqlx.Tx, product producte.ProductForm, typeID int) error
	AddProduct(tx *sqlx.Tx, product producte.ProductForm) (int, error)
	AddPriceHistory(tx *sqlx.Tx, product producte.ProductForm) (int, error)
	AddProductType(tx *sqlx.Tx, product producte.ProductForm, productID int) (int64, error)
	AddPriceHistoryProduct(tx *sqlx.Tx, typeID int, priceHistoryID int) error
	GetAllId(tx *sqlx.Tx, productID int) ([]producte.ProductDependencies, error)
	DeletePriceHistoryProduct(tx *sqlx.Tx, typeID int) error
	DeleteTypeProduct(tx *sqlx.Tx, typeID int) error
	DeletePriceHistory(tx *sqlx.Tx, historyID int) error
	DeleteProduct(tx *sqlx.Tx, productID int) error
	GetProductIdAndTypeIdByName(tx *sqlx.Tx, name, form string) (producte.ProductDependencies, error)
}

// Service по product
type Service interface {
	LoadAllProducts(tx *sqlx.Tx) ([]producte.Product, error)
	LoadProductFormByID(tx *sqlx.Tx, id int) ([]producte.ProductForm, error)
	AddNewProduct(tx *sqlx.Tx, product producte.ProductForm) error
	DeleteProductById(tx *sqlx.Tx, productID int) error
}
