package product_test

import (
	"task/config"
	"task/internal/entity/producte"
	"task/internal/modules/product"
	"testing"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

const (
	confPath = "../../../config/config.yaml"
)

var testProduct = producte.ProductForm{Name: "test", Form: "test", Amount: 1, Price: 1, DateStart: time.Now(), DateEnd: time.Now()}

func TestLoadAllProducts(t *testing.T) {
	config := config.GetConfiguration(confPath)
	require.NotEmpty(t, config)

	db, err := sqlx.Open("postgres", config)
	require.NoError(t, err)
	defer db.Close()

	tx, err := db.Beginx()
	require.NoError(t, err)
	defer tx.Rollback()

	repo := product.NewRepository()

	productList, err := repo.LoadAllProducts(tx)
	require.NoError(t, err)
	require.NotEmpty(t, productList)

	id, err := repo.AddProduct(tx, producte.ProductForm{Name: "test"})
	require.NoError(t, err)
	require.NotEmpty(t, id)

	productListRefreshed, err := repo.LoadAllProducts(tx)
	require.NoError(t, err)
	require.NotEmpty(t, productListRefreshed)

	var idList []int
	for _, product := range productListRefreshed {
		idList = append(idList, product.ProductId)
	}
	require.Contains(t, idList, id)
}

func TestLoadProductFormByID(t *testing.T) {
	config := config.GetConfiguration(confPath)
	require.NotEmpty(t, config)

	db, err := sqlx.Open("postgres", config)
	require.NoError(t, err)
	defer db.Close()

	tx, err := db.Beginx()
	require.NoError(t, err)
	defer tx.Rollback()

	repo := product.NewRepository()

	t.Run("добавление продукта", func(t *testing.T) {
		id, err := repo.AddProduct(tx, testProduct)
		require.NoError(t, err)
		require.NotEmpty(t, id)

		t.Run("добавление формы", func(t *testing.T) {
			tId, err := repo.AddProductType(tx, testProduct, id)
			require.NoError(t, err)
			require.NotEmpty(t, tId)

			t.Run("добавление цен", func(t *testing.T) {
				histId, err := repo.AddPriceHistory(tx, testProduct)
				require.NoError(t, err)
				require.NotEmpty(t, histId)

				t.Run("добавление связи цены и новара", func(t *testing.T) {
					err = repo.AddPriceHistoryProduct(tx, tId, histId)
					require.NoError(t, err)

					t.Run("проверка формы", func(t *testing.T) {
						productFormList, err := repo.LoadProductFormByID(tx, id)
						require.NoError(t, err)
						require.NotEmpty(t, productFormList)

						var formList []string

						for _, productType := range productFormList {
							formList = append(formList, productType.Form)
						}

						require.Contains(t, formList, "test")
					})
				})
			})
		})
	})
}

func TestGetProductIdByName(t *testing.T) {
	config := config.GetConfiguration(confPath)
	require.NotEmpty(t, config)

	db, err := sqlx.Open("postgres", config)
	require.NoError(t, err)
	defer db.Close()

	tx, err := db.Beginx()
	require.NoError(t, err)
	defer tx.Rollback()

	repo := product.NewRepository()

	t.Run("добавление продукта", func(t *testing.T) {
		id, err := repo.AddProduct(tx, testProduct)
		require.NoError(t, err)
		require.NotEmpty(t, id)

		t.Run("получение id по имени", func(t *testing.T) {
			searchId, err := repo.GetProductIdByName(tx, testProduct.Name)
			require.NoError(t, err)
			require.NotEmpty(t, searchId)

			require.Equal(t, id, searchId)
		})
	})
}

func TestGetTypeIdByProduct(t *testing.T) {
	config := config.GetConfiguration(confPath)
	require.NotEmpty(t, config)

	db, err := sqlx.Open("postgres", config)
	require.NoError(t, err)
	defer db.Close()

	tx, err := db.Beginx()
	require.NoError(t, err)
	defer tx.Rollback()

	repo := product.NewRepository()

	t.Run("добавление продукта", func(t *testing.T) {
		id, err := repo.AddProduct(tx, testProduct)
		require.NoError(t, err)
		require.NotEmpty(t, id)

		t.Run("добавление формы", func(t *testing.T) {
			tId, err := repo.AddProductType(tx, testProduct, id)
			require.NoError(t, err)
			require.NotEmpty(t, tId)

			t.Run("проверка id", func(t *testing.T) {
				searchTID, err := repo.GetTypeIdByProduct(tx, testProduct, id)
				require.NoError(t, err)
				require.NotEmpty(t, searchTID)

				require.Equal(t, tId, searchTID)
			})
		})
	})
}

func TestUpdateProductAmount(t *testing.T) {
	config := config.GetConfiguration(confPath)
	require.NotEmpty(t, config)

	db, err := sqlx.Open("postgres", config)
	require.NoError(t, err)
	defer db.Close()

	tx, err := db.Beginx()
	require.NoError(t, err)
	defer tx.Rollback()

	repo := product.NewRepository()

	t.Run("добавление продукта", func(t *testing.T) {
		id, err := repo.AddProduct(tx, testProduct)
		require.NoError(t, err)
		require.NotEmpty(t, id)

		t.Run("добавление формы", func(t *testing.T) {
			tId, err := repo.AddProductType(tx, testProduct, id)
			require.NoError(t, err)
			require.NotEmpty(t, tId)

			t.Run("добавление цен", func(t *testing.T) {
				histId, err := repo.AddPriceHistory(tx, testProduct)
				require.NoError(t, err)
				require.NotEmpty(t, histId)

				t.Run("добавление связи цены и новара", func(t *testing.T) {
					err = repo.AddPriceHistoryProduct(tx, tId, histId)
					require.NoError(t, err)

					t.Run("обновление количества", func(t *testing.T) {
						err := repo.UpdateProductAmount(tx, testProduct, tId)
						require.NoError(t, err)

						t.Run("проверка данных", func(t *testing.T) {
							productFormList, err := repo.LoadProductFormByID(tx, id)
							require.NoError(t, err)
							require.NotEmpty(t, productFormList)

							var amountList []int

							for _, productType := range productFormList {
								amountList = append(amountList, productType.Amount)
							}
							require.NotContains(t, amountList, testProduct.Amount)
						})
					})
				})
			})

		})
	})
}

func TestAddProduct(t *testing.T) {
	config := config.GetConfiguration(confPath)
	require.NotEmpty(t, config)

	db, err := sqlx.Open("postgres", config)
	require.NoError(t, err)
	defer db.Close()

	tx, err := db.Beginx()
	require.NoError(t, err)
	defer tx.Rollback()

	repo := product.NewRepository()

	t.Run("добавление продукта", func(t *testing.T) {
		id, err := repo.AddProduct(tx, testProduct)
		require.NoError(t, err)
		require.NotEmpty(t, id)

		t.Run("проверка поиском по имени", func(t *testing.T) {
			searchId, err := repo.GetProductIdByName(tx, testProduct.Name)
			require.NoError(t, err)
			require.NotEmpty(t, searchId)

			require.Equal(t, id, searchId)
		})
	})
}

func TestAddPriceHistory(t *testing.T) {
	config := config.GetConfiguration(confPath)
	require.NotEmpty(t, config)

	db, err := sqlx.Open("postgres", config)
	require.NoError(t, err)
	defer db.Close()

	tx, err := db.Beginx()
	require.NoError(t, err)
	defer tx.Rollback()

	repo := product.NewRepository()

	t.Run("добавление продукта", func(t *testing.T) {
		id, err := repo.AddProduct(tx, testProduct)
		require.NoError(t, err)
		require.NotEmpty(t, id)

		t.Run("добавление формы", func(t *testing.T) {
			tId, err := repo.AddProductType(tx, testProduct, id)
			require.NoError(t, err)
			require.NotEmpty(t, tId)

			t.Run("добавление цен", func(t *testing.T) {
				histId, err := repo.AddPriceHistory(tx, testProduct)
				require.NoError(t, err)
				require.NotEmpty(t, histId)

				t.Run("добавление связи цены и товара", func(t *testing.T) {
					err = repo.AddPriceHistoryProduct(tx, tId, histId)
					require.NoError(t, err)

					t.Run("проверка", func(t *testing.T) {
						var idList []int
						err := tx.Select(&idList, `select history_id from price_history `)
						require.NoError(t, err)

						require.Contains(t, idList, histId)
					})
				})
			})
		})
	})
}

func TestAddProductType(t *testing.T) {
	config := config.GetConfiguration(confPath)
	require.NotEmpty(t, config)

	db, err := sqlx.Open("postgres", config)
	require.NoError(t, err)
	defer db.Close()

	tx, err := db.Beginx()
	require.NoError(t, err)
	defer tx.Rollback()

	repo := product.NewRepository()

	t.Run("добавление продукта", func(t *testing.T) {
		id, err := repo.AddProduct(tx, testProduct)
		require.NoError(t, err)
		require.NotEmpty(t, id)

		t.Run("добавление формы", func(t *testing.T) {
			tId, err := repo.AddProductType(tx, testProduct, id)
			require.NoError(t, err)
			require.NotEmpty(t, tId)

			t.Run("проверка id формы по id продукта", func(t *testing.T) {
				searchTID, err := repo.GetTypeIdByProduct(tx, testProduct, id)
				require.NoError(t, err)
				require.NotEmpty(t, searchTID)

				require.Equal(t, tId, searchTID)
			})
		})
	})
}

func TestAddPriceHistoryProduct(t *testing.T) {
	config := config.GetConfiguration(confPath)
	require.NotEmpty(t, config)

	db, err := sqlx.Open("postgres", config)
	require.NoError(t, err)
	defer db.Close()

	tx, err := db.Beginx()
	require.NoError(t, err)
	defer tx.Rollback()

	repo := product.NewRepository()

	t.Run("добавление продукта", func(t *testing.T) {
		id, err := repo.AddProduct(tx, testProduct)
		require.NoError(t, err)
		require.NotEmpty(t, id)

		t.Run("добавление формы", func(t *testing.T) {
			tId, err := repo.AddProductType(tx, testProduct, id)
			require.NoError(t, err)
			require.NotEmpty(t, tId)

			t.Run("добавление цен", func(t *testing.T) {
				histId, err := repo.AddPriceHistory(tx, testProduct)
				require.NoError(t, err)
				require.NotEmpty(t, histId)

				t.Run("добавление связи цены и новара", func(t *testing.T) {
					err = repo.AddPriceHistoryProduct(tx, tId, histId)
					require.NoError(t, err)

					t.Run("Проверка", func(t *testing.T) {
						var testID int
						err = tx.Get(&testID, `select price_history_id from price_history_product where product_type_id = $1`, tId)
						require.NoError(t, err)

						require.Equal(t, histId, testID)
					})
				})
			})
		})
	})
}

func TestGetAllId(t *testing.T) {
	config := config.GetConfiguration(confPath)
	require.NotEmpty(t, config)

	db, err := sqlx.Open("postgres", config)
	require.NoError(t, err)
	defer db.Close()

	tx, err := db.Beginx()
	require.NoError(t, err)
	defer tx.Rollback()

	repo := product.NewRepository()

	t.Run("добавление продукта", func(t *testing.T) {
		productID, err := repo.AddProduct(tx, testProduct)
		require.NoError(t, err)
		require.NotEmpty(t, productID)

		t.Run("добавление формы", func(t *testing.T) {
			tId, err := repo.AddProductType(tx, testProduct, productID)
			require.NoError(t, err)
			require.NotEmpty(t, tId)

			t.Run("добавление цен", func(t *testing.T) {
				histId, err := repo.AddPriceHistory(tx, testProduct)
				require.NoError(t, err)
				require.NotEmpty(t, histId)

				t.Run("добавление связи цены и новара", func(t *testing.T) {
					err = repo.AddPriceHistoryProduct(tx, tId, histId)
					require.NoError(t, err)

					t.Run("получение всех ID", func(t *testing.T) {
						allID, err := repo.GetAllId(tx, productID)
						require.NoError(t, err)

						var productIdList []int
						var typeIdList []int
						var historyIdList []int

						for _, id := range allID {
							productIdList = append(productIdList, id.ProductId)
							typeIdList = append(typeIdList, id.TypeId)
							historyIdList = append(historyIdList, id.PriceHistoryId)
						}

						require.Contains(t, productIdList, productID)
						require.Contains(t, typeIdList, tId)
						require.Contains(t, historyIdList, histId)

					})
				})
			})
		})
	})
}

func TestDeletePriceHistoryProduct(t *testing.T) {
	config := config.GetConfiguration(confPath)
	require.NotEmpty(t, config)

	db, err := sqlx.Open("postgres", config)
	require.NoError(t, err)
	defer db.Close()

	tx, err := db.Beginx()
	require.NoError(t, err)
	defer tx.Rollback()

	repo := product.NewRepository()

	t.Run("добавление продукта", func(t *testing.T) {
		productID, err := repo.AddProduct(tx, testProduct)
		require.NoError(t, err)
		require.NotEmpty(t, productID)

		t.Run("добавление формы", func(t *testing.T) {
			tId, err := repo.AddProductType(tx, testProduct, productID)
			require.NoError(t, err)
			require.NotEmpty(t, tId)

			t.Run("добавление цен", func(t *testing.T) {
				histId, err := repo.AddPriceHistory(tx, testProduct)
				require.NoError(t, err)
				require.NotEmpty(t, histId)

				t.Run("добавление связи цены и новара", func(t *testing.T) {
					err = repo.AddPriceHistoryProduct(tx, tId, histId)
					require.NoError(t, err)

					t.Run("удаление связи цены и товара", func(t *testing.T) {
						err = repo.DeletePriceHistoryProduct(tx, tId)
						require.NoError(t, err)

						t.Run("проверка", func(t *testing.T) {
							var typeIdList []int
							err = tx.Select(&typeIdList, `select product_type_id from price_history_product`)

							require.NotContains(t, typeIdList, tId)
						})
					})
				})
			})
		})
	})
}

func TestDeleteTypeProduct(t *testing.T) {
	config := config.GetConfiguration(confPath)
	require.NotEmpty(t, config)

	db, err := sqlx.Open("postgres", config)
	require.NoError(t, err)
	defer db.Close()

	tx, err := db.Beginx()
	require.NoError(t, err)
	defer tx.Rollback()

	repo := product.NewRepository()

	t.Run("добавление продукта", func(t *testing.T) {
		productID, err := repo.AddProduct(tx, testProduct)
		require.NoError(t, err)
		require.NotEmpty(t, productID)

		t.Run("добавление формы", func(t *testing.T) {
			tId, err := repo.AddProductType(tx, testProduct, productID)
			require.NoError(t, err)
			require.NotEmpty(t, tId)

			t.Run("удаление формы продукта", func(t *testing.T) {
				err = repo.DeleteTypeProduct(tx, tId)
				require.NoError(t, err)

				t.Run("проверка", func(t *testing.T) {
					var formIdList []int
					err = tx.Select(&formIdList, `select type_id from product_types`)
					require.NoError(t, err)

					require.NotContains(t, formIdList, tId)
				})
			})
		})
	})
}

func TestDeletePriceHistory(t *testing.T) {
	config := config.GetConfiguration(confPath)
	require.NotEmpty(t, config)

	db, err := sqlx.Open("postgres", config)
	require.NoError(t, err)
	defer db.Close()

	tx, err := db.Beginx()
	require.NoError(t, err)
	defer tx.Rollback()

	repo := product.NewRepository()

	t.Run("добавление продукта", func(t *testing.T) {
		productID, err := repo.AddProduct(tx, testProduct)
		require.NoError(t, err)
		require.NotEmpty(t, productID)

		t.Run("добавление цен", func(t *testing.T) {
			histId, err := repo.AddPriceHistory(tx, testProduct)
			require.NoError(t, err)
			require.NotEmpty(t, histId)

			t.Run("удаление цен", func(t *testing.T) {
				err = repo.DeletePriceHistory(tx, histId)
				require.NoError(t, err)

				t.Run("проверка", func(t *testing.T) {
					var histIdList []int
					err = tx.Select(&histIdList, `select history_id from price_history`)
					require.NoError(t, err)

					require.NotContains(t, histIdList, histId)
				})
			})
		})
	})
}

func TestDeleteProduct(t *testing.T) {
	config := config.GetConfiguration(confPath)
	require.NotEmpty(t, config)

	db, err := sqlx.Open("postgres", config)
	require.NoError(t, err)
	defer db.Close()

	tx, err := db.Beginx()
	require.NoError(t, err)
	defer tx.Rollback()

	repo := product.NewRepository()

	t.Run("добавление продукта", func(t *testing.T) {
		productID, err := repo.AddProduct(tx, testProduct)
		require.NoError(t, err)
		require.NotEmpty(t, productID)

		t.Run("удаление продукта", func(t *testing.T) {
			err = repo.DeleteProduct(tx, productID)
			require.NoError(t, err)

			t.Run("проверка", func(t *testing.T) {
				var productIdList []int
				err = tx.Select(&productIdList, `select product_id from products`)
				require.NoError(t, err)

				require.NotContains(t, productIdList, productID)
			})
		})
	})
}
