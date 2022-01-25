package product

import (
	"database/sql"
	"task/internal/entity/global"
	"task/internal/entity/producte"

	"github.com/jmoiron/sqlx"
)

type repository struct {
}

// NewRepository репозиторий product
func NewRepository() Repository {
	return &repository{}
}

func (r *repository) LoadAllProducts(tx *sqlx.Tx) ([]producte.Product, error) {
	sqlQuery := `select product_id, name, date_added, included from products order by product_id`
	data := make([]producte.Product, 0)

	err := tx.Select(&data, sqlQuery)
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, global.ErrNoDataFound
	}

	return data, nil
}

func (r *repository) LoadProductFormByID(tx *sqlx.Tx, id int) ([]producte.ProductForm, error) {
	sqlQuery := `select pt.type_id, pt.form, p.name, pt.amount, ph.date_start, ph.price, ph.date_end
				from product_types pt
				join products p on p.product_id = pt.product_id
				join price_history_product hpp on pt.type_id = hpp.product_type_id
				join price_history ph on ph.history_id = hpp.price_history_id
				where p.product_id = $1 and pt.available = true`
	data := make([]producte.ProductForm, 0)

	err := tx.Select(&data, sqlQuery, id)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, global.ErrNoDataFound
	}
	return data, nil
}

func (r *repository) UpdateProductAmount(tx *sqlx.Tx, product producte.ProductForm, typeID int) error {
	sqlQuery := `update product_types set amount = amount + $1 where type_id = $2`

	_, err := tx.Exec(sqlQuery, product.Amount, typeID)

	if err != nil {
		return err
	}
	return nil
}

func (r *repository) AddProduct(tx *sqlx.Tx, product producte.ProductForm) (int, error) {
	sqlQuery := `insert into products (name, date_added, included)
				 values ($1, CURRENT_DATE, true) returning product_id`
	var productID int

	err := tx.QueryRow(sqlQuery, product.Name).Scan(&productID)
	if err != nil {
		return 0, err
	}
	return productID, nil
}

func (r *repository) AddPriceHistory(tx *sqlx.Tx, product producte.ProductForm) (int, error) {
	sqlQuery := `insert into price_history(date_start, date_end, price) values ($1, $2, $3) returning history_id`
	var priceHistoryID int

	err := tx.QueryRow(sqlQuery, product.DateStart, product.DateEnd, product.Price).Scan(&priceHistoryID)
	if err != nil {
		return 0, err
	}

	return priceHistoryID, nil
}

func (r *repository) AddProductType(tx *sqlx.Tx, product producte.ProductForm, productID int) (int64, error) {
	sqlQuery := `insert into product_types(product_id, form, amount, available) values ($1, $2, $3, true) returning type_id`
	var typeID int

	err := tx.QueryRow(sqlQuery, productID, product.Form, product.Amount).Scan(&typeID)
	if err != nil {
		return 0, err
	}

	return int64(typeID), nil
}

func (r *repository) AddPriceHistoryProduct(tx *sqlx.Tx, typeID int, priceHistoryID int) error {
	sqlQuery := `insert into price_history_product (product_type_id, price_history_id) values ($1, $2)`

	_, err := tx.Exec(sqlQuery, typeID, priceHistoryID)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) GetAllId(tx *sqlx.Tx, productID int) ([]producte.ProductDependencies, error) {
	sqlQuery := `select p.product_id, pt.type_id, ph.history_id
				from products p
				join product_types pt on p.product_id = pt.product_id
				join price_history_product hpp on pt.type_id = hpp.product_type_id
				join price_history ph on ph.history_id = hpp.price_history_id where p.product_id = $1`
	var productDependencies []producte.ProductDependencies

	err := tx.Select(&productDependencies, sqlQuery, productID)
	if err != nil {
		return nil, err
	}
	return productDependencies, nil
}

func (r *repository) DeletePriceHistoryProduct(tx *sqlx.Tx, typeID int) error {
	sqlQuery := `delete from price_history_product where price_history_product.product_type_id = $1`

	_, err := tx.Exec(sqlQuery, typeID)

	if err != nil {
		return err
	}
	return nil
}

func (r *repository) DeleteTypeProduct(tx *sqlx.Tx, typeID int) error {
	sqlQuery := `delete from product_types where product_types.type_id = $1`
	_, err := tx.Exec(sqlQuery, typeID)

	if err != nil {
		return err
	}
	return nil
}

func (r *repository) DeletePriceHistory(tx *sqlx.Tx, historyID int) error {
	sqlQuery := `delete from price_history where price_history.history_id = $1`
	_, err := tx.Exec(sqlQuery, historyID)

	if err != nil {
		return err
	}
	return nil
}

func (r *repository) DeleteProduct(tx *sqlx.Tx, productID int) error {
	sqlQuery := `delete from products where products.product_id = $1`
	_, err := tx.Exec(sqlQuery, productID)

	if err != nil {
		return err
	}
	return nil
}

func (r *repository) GetProductIdAndTypeIdByName(tx *sqlx.Tx, name, form string) (producte.ProductDependencies, error) {
	var data producte.ProductDependencies

	err := tx.Get(&data, `
		select p.product_id, pt.type_id
		from products p
		left outer join product_types pt on (p.product_id = pt.product_id and pt.form = $2)
		where p.name = $1
	`, name, form)

	if err != nil {
		if err == sql.ErrNoRows {
			return producte.ProductDependencies{}, global.ErrNoDataFound
		}
		return producte.ProductDependencies{}, err
	}

	return data, nil
}
