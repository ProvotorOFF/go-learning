package main_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	main "order-api-start/cmd"
	"order-api-start/configs"
	"order-api-start/internal/order"
	"order-api-start/internal/product"
	"order-api-start/internal/user"
	"order-api-start/pkg/db"
	"order-api-start/pkg/jwt"
	"testing"

	"github.com/stretchr/testify/require"
)

func setUp(t *testing.T, config *configs.Config) (*user.User, *product.Product) {

	database, err := db.NewDb(config)
	require.NoError(t, err)

	testUser := user.User{
		Phone: "79876543232",
	}
	database.Create(&testUser)

	productTest := product.Product{
		Name:        "Test Product",
		Description: "Some product",
		Images:      []string{"https://example.com/image1.jpg"},
	}
	database.Create(&productTest)

	t.Cleanup(func() {
		database.Unscoped().Where("id = ?", testUser.ID).Delete(&user.User{})
		database.Unscoped().Where("id = ?", productTest.ID).Delete(&product.Product{})
		database.Exec("delete from orders where user_id = ?", testUser.ID)
	})

	return &testUser, &productTest
}

func TestOrderSuccess(t *testing.T) {
	config := configs.LoadConfig()
	ts := httptest.NewServer(main.App())
	defer ts.Close()

	testUser, testProduct := setUp(t, config)

	token, err := jwt.NewJWT(config.Secret).Create(testUser.Phone)
	require.NoError(t, err)

	payload := order.OrderCreateRequest{
		Products: []uint{testProduct.ID},
	}

	payloadJSON, err := json.Marshal(payload)
	require.NoError(t, err)

	req, err := http.NewRequest(
		http.MethodPost,
		ts.URL+"/order",
		bytes.NewBuffer(payloadJSON),
	)
	require.NoError(t, err)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	require.Equal(t, http.StatusCreated, resp.StatusCode)
}
