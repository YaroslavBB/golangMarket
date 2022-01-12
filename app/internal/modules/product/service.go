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
	var err error
	var allID producte.AllId

	allID.ProductId, err = s.repo.GetProductIdByName(tx, product.Name)
	if err != nil {
		return err
	}
	if allID.ProductId == 0 {
		allID.ProductId, err = s.repo.AddProduct(tx, product)
		if err != nil {
			return err
		}
		allID.PriceHistoryId, err = s.repo.AddPriceHistory(tx, product)
		if err != nil {
			return err
		}
		allID.TypeId, err = s.repo.AddProductType(tx, product, allID.ProductId)
		if err != nil {
			return err
		}
		err = s.repo.AddPriceHistoryProduct(tx, allID.TypeId, allID.PriceHistoryId)
		if err != nil {
			return err
		}
	} else {
		allID.TypeId, err = s.repo.GetTypeIdByProduct(tx, product, allID.ProductId)
		if err != nil {
			return err
		}
		if allID.TypeId == 0 {
			allID.PriceHistoryId, err = s.repo.AddPriceHistory(tx, product)
			if err != nil {
				return err
			}
			allID.TypeId, err = s.repo.AddProductType(tx, product, allID.ProductId)
			if err != nil {
				return err
			}
			err = s.repo.AddPriceHistoryProduct(tx, allID.TypeId, allID.PriceHistoryId)
			if err != nil {
				return err
			}
		} else {
			err = s.repo.UpdateProductAmount(tx, product, allID.TypeId)
			if err != nil {
				return err
			}
		}
	}
	return tx.Commit()
}

func (s *service) DeleteProductById(tx *sqlx.Tx, productID int) error {
	var allID []producte.AllId
	var err error

	allID, err = s.repo.GetAllId(tx, productID)
	if err != nil {
		return err
	}

	for _, id := range allID {
		err = s.repo.DeletePriceHistoryProduct(tx, id.TypeId)
		if err != nil {
			return nil
		}
		err = s.repo.DeleteTypeProduct(tx, id.TypeId)
		if err != nil {
			return err
		}
		err = s.repo.DeletePriceHistory(tx, id.PriceHistoryId)
		if err != nil {
			return err
		}
	}
	err = s.repo.DeleteProduct(tx, productID)
	if err != nil {
		return err
	}

	return tx.Commit()
}
