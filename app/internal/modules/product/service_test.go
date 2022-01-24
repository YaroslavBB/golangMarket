package product_test

import (
	"database/sql"
	"task/internal/entity/global"
	"task/internal/entity/producte"
	"task/internal/modules/product"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestDeleteProductById(t *testing.T) {
	r := require.New(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := product.NewMockRepository(ctrl)
	service := product.NewService(mockRepo)

	testID := producte.ProductDependencies{
		ProductId:      98,
		TypeId:         sql.NullInt64{Int64: 18, Valid: true},
		PriceHistoryId: 32,
	}

	testIDList := make([]producte.ProductDependencies, 0)
	testIDList = append(testIDList, testID)

	t.Run("удаление продукта", func(t *testing.T) {
		mockRepo.EXPECT().GetAllId(gomock.Any(), testID.ProductId).Return(testIDList, nil).Times(1)
		mockRepo.EXPECT().DeletePriceHistoryProduct(gomock.Any(), int(testID.TypeId.Int64)).Return(nil).Times(1)
		mockRepo.EXPECT().DeleteTypeProduct(gomock.Any(), int(testID.TypeId.Int64)).Return(nil).Times(1)
		mockRepo.EXPECT().DeletePriceHistory(gomock.Any(), testID.PriceHistoryId).Return(nil).Times(1)
		mockRepo.EXPECT().DeleteProduct(gomock.Any(), testID.ProductId).Return(nil).Times(1)

		r.NoError(service.DeleteProductById(nil, testID.ProductId))
	})

	t.Run("удаление продукта - ошибка в получении id", func(t *testing.T) {
		mockRepo.EXPECT().GetAllId(gomock.Any(), testID.ProductId).Return(nil, global.ErrNoDataFound).Times(1)

		r.Equal(global.ErrNoDataFound, service.DeleteProductById(nil, testID.ProductId))
	})

	t.Run("удаление продукта - ошибка в удалении связи истории цен и формы", func(t *testing.T) {
		mockRepo.EXPECT().GetAllId(gomock.Any(), testID.ProductId).Return(testIDList, nil).Times(1)
		mockRepo.EXPECT().DeletePriceHistoryProduct(gomock.Any(), int(testID.TypeId.Int64)).Return(global.ErrNoDataFound).Times(1)

		r.Error(service.DeleteProductById(nil, testID.ProductId))
	})

	t.Run("удаление продукта - ошибка в удалении формы", func(t *testing.T) {
		mockRepo.EXPECT().GetAllId(gomock.Any(), testID.ProductId).Return(testIDList, nil).Times(1)
		mockRepo.EXPECT().DeletePriceHistoryProduct(gomock.Any(), int(testID.TypeId.Int64)).Return(nil).Times(1)
		mockRepo.EXPECT().DeleteTypeProduct(gomock.Any(), int(testID.TypeId.Int64)).Return(global.ErrNoDataFound).Times(1)

		r.Error(service.DeleteProductById(nil, testID.ProductId))
	})

	t.Run("удаление продукта - ошибка в удалении истории цен", func(t *testing.T) {
		mockRepo.EXPECT().GetAllId(gomock.Any(), testID.ProductId).Return(testIDList, nil).Times(1)
		mockRepo.EXPECT().DeletePriceHistoryProduct(gomock.Any(), int(testID.TypeId.Int64)).Return(nil).Times(1)
		mockRepo.EXPECT().DeleteTypeProduct(gomock.Any(), int(testID.TypeId.Int64)).Return(nil).Times(1)
		mockRepo.EXPECT().DeletePriceHistory(gomock.Any(), testID.PriceHistoryId).Return(global.ErrNoDataFound).Times(1)

		r.Error(service.DeleteProductById(nil, testID.ProductId))
	})

	t.Run("удаление продукта - ошибка в удалении продукта", func(t *testing.T) {
		mockRepo.EXPECT().GetAllId(gomock.Any(), testID.ProductId).Return(testIDList, nil).Times(1)
		mockRepo.EXPECT().DeletePriceHistoryProduct(gomock.Any(), int(testID.TypeId.Int64)).Return(nil).Times(1)
		mockRepo.EXPECT().DeleteTypeProduct(gomock.Any(), int(testID.TypeId.Int64)).Return(nil).Times(1)
		mockRepo.EXPECT().DeletePriceHistory(gomock.Any(), testID.PriceHistoryId).Return(nil).Times(1)
		mockRepo.EXPECT().DeleteProduct(gomock.Any(), testID.ProductId).Return(global.ErrNoDataFound).Times(1)

		r.Error(service.DeleteProductById(nil, testID.ProductId))
	})
}

func TestAddNewProduct(t *testing.T) {
	r := require.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := product.NewMockRepository(ctrl)
	service := product.NewService(mockRepo)

	testProduct = producte.ProductForm{Name: "test", Form: "test", Amount: 1, Price: 1, DateStart: time.Now(), DateEnd: time.Now()}

	t.Run("Добавление нового продукта", func(t *testing.T) {
		mockRepo.EXPECT().GetProductIdAndTypeIdByName(gomock.Any(), testProduct.Name, testProduct.Form).Return(producte.ProductDependencies{}, nil).Times(1)
		mockRepo.EXPECT().AddProduct(gomock.Any(), testProduct).Return(12, nil).Times(1)
		mockRepo.EXPECT().AddPriceHistory(gomock.Any(), testProduct).Return(34, nil).Times(1)
		mockRepo.EXPECT().AddProductType(gomock.Any(), testProduct, 12).Return(int64(43), nil).Times(1)
		mockRepo.EXPECT().AddPriceHistoryProduct(gomock.Any(), int(43), 34).Return(nil).Times(1)

		r.NoError(service.AddNewProduct(nil, testProduct))
	})
	t.Run("Добавление формы в существующий продукт", func(t *testing.T) {
		mockRepo.EXPECT().GetProductIdAndTypeIdByName(gomock.Any(), testProduct.Name, testProduct.Form).
			Return(producte.ProductDependencies{ProductId: 23, TypeId: sql.NullInt64{Int64: 0, Valid: false}}, nil).Times(1)
		mockRepo.EXPECT().AddPriceHistory(gomock.Any(), testProduct).Return(13, nil).Times(1)
		mockRepo.EXPECT().AddProductType(gomock.Any(), testProduct, 23).Return(int64(12), nil).Times(1)
		mockRepo.EXPECT().AddPriceHistoryProduct(gomock.Any(), int(12), 13).Return(nil).Times(1)

		r.NoError(service.AddNewProduct(nil, testProduct))
	})
	t.Run("Увеличение колличества продукта", func(t *testing.T) {
		mockRepo.EXPECT().GetProductIdAndTypeIdByName(gomock.Any(), testProduct.Name, testProduct.Form).
			Return(producte.ProductDependencies{ProductId: 14, TypeId: sql.NullInt64{Int64: 32, Valid: true}}, nil).Times(1)
		mockRepo.EXPECT().UpdateProductAmount(gomock.Any(), testProduct, 32).Return(nil).Times(1)

		r.NoError(service.AddNewProduct(nil, testProduct))
	})
	t.Run("Ошибка при добавлении нового продукта на этапе получения id формы и продукта", func(t *testing.T) {
		mockRepo.EXPECT().GetProductIdAndTypeIdByName(gomock.Any(), testProduct.Name, testProduct.Form).
			Return(producte.ProductDependencies{}, global.ErrNoDataFound).Times(1)

		r.Equal(global.ErrNoDataFound, service.AddNewProduct(nil, testProduct))
	})
	t.Run("ошибка при добавлени товара на этапе добавления продукта", func(t *testing.T) {
		mockRepo.EXPECT().GetProductIdAndTypeIdByName(gomock.Any(), testProduct.Name, testProduct.Form).Return(producte.ProductDependencies{}, nil).Times(1)
		mockRepo.EXPECT().AddProduct(gomock.Any(), testProduct).Return(0, global.ErrNoDataFound).Times(1)

		r.Equal(global.ErrNoDataFound, service.AddNewProduct(nil, testProduct))
	})
	t.Run("ошибка при добавлении продукта на этапе добавления истории цен", func(t *testing.T) {
		mockRepo.EXPECT().GetProductIdAndTypeIdByName(gomock.Any(), testProduct.Name, testProduct.Form).Return(producte.ProductDependencies{}, nil).Times(1)
		mockRepo.EXPECT().AddProduct(gomock.Any(), testProduct).Return(12, nil).Times(1)
		mockRepo.EXPECT().AddPriceHistory(gomock.Any(), testProduct).Return(0, global.ErrNoDataFound).Times(1)

		r.Equal(global.ErrNoDataFound, service.AddNewProduct(nil, testProduct))
	})
	t.Run("ошибка при добавлении продукта на этапе добавления формы", func(t *testing.T) {
		mockRepo.EXPECT().GetProductIdAndTypeIdByName(gomock.Any(), testProduct.Name, testProduct.Form).Return(producte.ProductDependencies{}, nil).Times(1)
		mockRepo.EXPECT().AddProduct(gomock.Any(), testProduct).Return(12, nil).Times(1)
		mockRepo.EXPECT().AddPriceHistory(gomock.Any(), testProduct).Return(34, nil).Times(1)
		mockRepo.EXPECT().AddProductType(gomock.Any(), testProduct, 12).Return(int64(0), global.ErrNoDataFound).Times(1)

		r.Equal(global.ErrNoDataFound, service.AddNewProduct(nil, testProduct))
	})
	t.Run("ошибка при добавлении продукта на этапе добавления связи истории цен и формы", func(t *testing.T) {
		mockRepo.EXPECT().GetProductIdAndTypeIdByName(gomock.Any(), testProduct.Name, testProduct.Form).Return(producte.ProductDependencies{}, nil).Times(1)
		mockRepo.EXPECT().AddProduct(gomock.Any(), testProduct).Return(12, nil).Times(1)
		mockRepo.EXPECT().AddPriceHistory(gomock.Any(), testProduct).Return(34, nil).Times(1)
		mockRepo.EXPECT().AddProductType(gomock.Any(), testProduct, 12).Return(int64(43), nil).Times(1)
		mockRepo.EXPECT().AddPriceHistoryProduct(gomock.Any(), int(43), 34).Return(global.ErrNoDataFound).Times(1)

		r.Equal(global.ErrNoDataFound, service.AddNewProduct(nil, testProduct))
	})
	t.Run("ошибка при увеличении колличства продукта", func(t *testing.T) {
		mockRepo.EXPECT().GetProductIdAndTypeIdByName(gomock.Any(), testProduct.Name, testProduct.Form).
			Return(producte.ProductDependencies{ProductId: 14, TypeId: sql.NullInt64{Int64: 32, Valid: true}}, nil).Times(1)
		mockRepo.EXPECT().UpdateProductAmount(gomock.Any(), testProduct, 32).Return(global.ErrNoDataFound).Times(1)

		r.Equal(global.ErrNoDataFound, service.AddNewProduct(nil, testProduct))
	})
}
