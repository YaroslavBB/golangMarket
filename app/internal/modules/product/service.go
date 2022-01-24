package product

import (
	"task/internal/entity/producte"

	"github.com/jmoiron/sqlx"
)

type service struct {
	repo Repository
}

// NewService по product
func NewService(r Repository) Service {
	return &service{
		r,
	}
}

func (s *service) LoadAllProducts(tx *sqlx.Tx) ([]producte.Product, error) {
	return s.repo.LoadAllProducts(tx)
}

func (s *service) LoadProductFormByID(tx *sqlx.Tx, id int) ([]producte.ProductForm, error) {
	return s.repo.LoadProductFormByID(tx, id)
}

func (s *service) AddNewProduct(tx *sqlx.Tx, product producte.ProductForm) error {
	allID, err := s.repo.GetProductIdAndTypeIdByName(tx, product.Name, product.Form)
	if err != nil {
		return err
	}

	switch {
	case allID.ProductId != 0 && allID.TypeId.Valid:
		err = s.repo.UpdateProductAmount(tx, product, int(allID.TypeId.Int64))
		if err != nil {
			return err
		}

	case allID.ProductId != 0 && !allID.TypeId.Valid:
		allID.PriceHistoryId, err = s.repo.AddPriceHistory(tx, product)
		if err != nil {
			return err
		}
		allID.TypeId.Int64, err = s.repo.AddProductType(tx, product, allID.ProductId)
		if err != nil {
			return err
		}
		err = s.repo.AddPriceHistoryProduct(tx, int(allID.TypeId.Int64), allID.PriceHistoryId)
		if err != nil {
			return err
		}

	case allID.ProductId == 0:
		allID.ProductId, err = s.repo.AddProduct(tx, product)
		if err != nil {
			return err
		}
		allID.PriceHistoryId, err = s.repo.AddPriceHistory(tx, product)
		if err != nil {
			return err
		}
		allID.TypeId.Int64, err = s.repo.AddProductType(tx, product, allID.ProductId)
		if err != nil {
			return err
		}
		err = s.repo.AddPriceHistoryProduct(tx, int(allID.TypeId.Int64), allID.PriceHistoryId)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *service) DeleteProductById(tx *sqlx.Tx, productID int) error {
	var allID []producte.AllId
	var err error

	allID, err = s.repo.GetAllId(tx, productID)
	if err != nil {
		return err
	}

	for _, id := range allID {
		err = s.repo.DeletePriceHistoryProduct(tx, int(id.TypeId.Int64))
		if err != nil {
			return err
		}
		err = s.repo.DeleteTypeProduct(tx, int(id.TypeId.Int64))
		if err != nil {
			return err
		}
		err = s.repo.DeletePriceHistory(tx, id.PriceHistoryId)
		if err != nil {
			return err
		}
	}
	return s.repo.DeleteProduct(tx, productID)
}
