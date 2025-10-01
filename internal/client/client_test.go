package client

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchProducts_Success(t *testing.T) {
	// 1. Arrange (Подготовка)
	mockResponseJSON := `{
		"data": {
			"filterProduct": [
				{
					"id": 123,
					"name": "Тестовый товар из мока",
					"vendorCode": "MOCK-001"
				}
			]
		}
	}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(mockResponseJSON))
	}))
	defer server.Close()

	apiClient := NewClient(server.URL)

	// 2. Act (Действие)
	products, err := apiClient.FetchProducts()

	// 3. Assert (Проверка)
	if err != nil {
		t.Fatalf("FetchProducts() вернула неожиданную ошибку: %v", err)
	}

	if len(products) != 1 {
		t.Fatalf("Ожидали получить 1 товар, а получили %d", len(products))
	}

	if products[0].Name != "Тестовый товар из мока" {
		t.Errorf("Имя товара не совпадает. Ожидали 'Тестовый товар из мока', получили '%s'", products[0].Name)
	}
}

func TestFetchProducts_ServerError(t *testing.T) {
	// 1. Arrange
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}))
	defer server.Close()

	apiClient := NewClient(server.URL)

	// 2. Act
	_, err := apiClient.FetchProducts()

	// 3. Assert
	if err == nil {
		t.Fatal("Ожидали получить ошибку от сервера, но не получили ее.")
	}
}
