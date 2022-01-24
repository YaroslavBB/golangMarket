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
	productDependencies, err := s.repo.GetProductIdAndTypeIdByName(tx, product.Name, product.Form)
	if err != nil {
		return err
	}

	switch {
	case productDependencies.ProductId != 0 && productDependencies.TypeId.Valid:
		err = s.repo.UpdateProductAmount(tx, product, int(productDependencies.TypeId.Int64))
		if err != nil {
			return err
		}

	case productDependencies.ProductId != 0 && !productDependencies.TypeId.Valid:
		productDependencies.PriceHistoryId, err = s.repo.AddPriceHistory(tx, product)
		if err != nil {
			return err
		}
		productDependencies.TypeId.Int64, err = s.repo.AddProductType(tx, product, productDependencies.ProductId)
		if err != nil {
			return err
		}
		err = s.repo.AddPriceHistoryProduct(tx, int(productDependencies.TypeId.Int64), productDependencies.PriceHistoryId)
		if err != nil {
			return err
		}

	case productDependencies.ProductId == 0:
		productDependencies.ProductId, err = s.repo.AddProduct(tx, product)
		if err != nil {
			return err
		}
		productDependencies.PriceHistoryId, err = s.repo.AddPriceHistory(tx, product)
		if err != nil {
			return err
		}
		productDependencies.TypeId.Int64, err = s.repo.AddProductType(tx, product, productDependencies.ProductId)
		if err != nil {
			return err
		}
		err = s.repo.AddPriceHistoryProduct(tx, int(productDependencies.TypeId.Int64), productDependencies.PriceHistoryId)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *service) DeleteProductById(tx *sqlx.Tx, productID int) error {
	var productDependencies []producte.ProductDependencies
	var err error

	productDependencies, err = s.repo.GetAllId(tx, productID)
	if err != nil {
		return err
	}

	for _, id := range productDependencies {
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
