package producte

import (
	"database/sql"
	"time"
)

type Product struct {
	ProductId int       `db:"product_id" json:"productId"`
	Name      string    `db:"name" json:"name"`
	Included  bool      `db:"included" json:"included"`
	DateAdded time.Time `db:"date_added" json:"dateAdded"`
}

type ProductForm struct {
	ProductId     int       `db:"product_id" json:"productId"`
	IdProductForm int       `db:"type_id" json:"typeId"`
	Name          string    `db:"name" json:"name"`
	Form          string    `db:"form" json:"form"`
	Amount        int       `db:"amount" json:"amount"`
	Price         int       `db:"price" json:"price"`
	DateAdded     time.Time `db:"date_added" json:"dateAdded"`
	DateStart     time.Time `db:"date_start" json:"dateStart"`
	DateEnd       time.Time `db:"date_end" json:"dateEnd"`
}

type ProductDependencies struct {
	ProductId      int           `db:"product_id"`
	TypeId         sql.NullInt64 `db:"type_id"`
	PriceHistoryId int           `db:"history_id"`
}
