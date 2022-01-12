package product_test

import (
	"task/config"
	"task/internal/entity/producte"
	"task/internal/modules/product"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
)

func TestDeleteProductById(t *testing.T) {
	config := config.GetConfiguration(confPath)
	require.NotEmpty(t, config)

	db, err := sqlx.Open("postgres", config)
	require.NoError(t, err)
	defer db.Close()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := product.NewMockRepository(ctrl)
	service := product.NewService(mockRepo)

	tx, err := db.Beginx()
	require.NoError(t, err)
	defer tx.Rollback()

	testID := producte.AllId{
		ProductId:      9,
		TypeId:         6,
		PriceHistoryId: 0,
	}

	testIDList := make([]producte.AllId, 0)
	testIDList = append(testIDList, testID)

	t.Run("удаление продукта", func(t *testing.T) {
		mockRepo.EXPECT().GetAllId(tx, 98).Return(testIDList, nil).Times(1)
		mockRepo.EXPECT().DeletePriceHistoryProduct(tx, testID.TypeId).Return(nil).Times(1)
		mockRepo.EXPECT().DeleteTypeProduct(tx, testID.TypeId).Return(nil).Times(1)
		mockRepo.EXPECT().DeletePriceHistory(tx, testID.PriceHistoryId).Return(nil).Times(1)
		mockRepo.EXPECT().DeleteProduct(tx, 98).Return(nil).Times(1)

		service.DeleteProductById(tx, 98)
	})
}

func TestAddNewProduct(t *testing.T) {
	config := config.GetConfiguration(confPath)
	require.NotEmpty(t, config)

	db, err := sqlx.Open("postgres", config)
	require.NoError(t, err)
	defer db.Close()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := product.NewMockRepository(ctrl)
	service := product.NewService(mockRepo)

	tx, err := db.Beginx()
	require.NoError(t, err)
	defer tx.Rollback()

	testProduct = producte.ProductForm{Name: "test", Form: "test", Amount: 1, Price: 1, DateStart: time.Now(), DateEnd: time.Now()}

	t.Run("Добавление продукта ", func(t *testing.T) {
		mockRepo.EXPECT().GetProductIdByName(tx, testProduct.Name).Return(0, nil).Times(1)
		mockRepo.EXPECT().AddProduct(tx, testProduct).Return(1, nil).Times(1)
		mockRepo.EXPECT().AddPriceHistory(tx, testProduct).Return(1, nil).Times(1)
		mockRepo.EXPECT().AddProductType(tx, testProduct, 1).Return(1, nil).Times(1)
		mockRepo.EXPECT().AddPriceHistoryProduct(tx, 1, 1).Return(nil).Times(1)

		service.AddNewProduct(tx, testProduct)
	})

	t.Run("Добавление формы в существующий продукт", func(t *testing.T) {
		mockRepo.EXPECT().GetProductIdByName(tx, testProduct.Name).Return(14, nil).Times(1)
		mockRepo.EXPECT().GetTypeIdByProduct(tx, testProduct, 14).Return(0, nil).Times(1)
		mockRepo.EXPECT().AddPriceHistory(tx, testProduct).Return(17, nil).Times(1)
		mockRepo.EXPECT().AddProductType(tx, testProduct, 14).Return(19, nil).Times(1)
		mockRepo.EXPECT().AddPriceHistoryProduct(tx, 19, 17).Return(nil).Times(1)

		service.AddNewProduct(tx, testProduct)
	})

	t.Run("Увеличение колличества товара", func(t *testing.T) {
		mockRepo.EXPECT().GetProductIdByName(tx, testProduct.Name).Return(12, nil).Times(1)
		mockRepo.EXPECT().GetTypeIdByProduct(tx, testProduct, 12).Return(32, nil).Times(1)
		mockRepo.EXPECT().UpdateProductAmount(tx, testProduct, 32).Return(nil).Times(1)

		service.AddNewProduct(tx, testProduct)
	})
}
