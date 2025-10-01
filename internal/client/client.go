package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/PavelBradnitski/yml-generator/internal/models"
)

// Client - это клиент для работы с GraphQL API.
type Client struct {
	apiURL     string
	httpClient *http.Client
}

// NewClient создает новый экземпляр API клиента.
func NewClient(apiURL string) *Client {
	return &Client{
		apiURL:     apiURL,
		httpClient: &http.Client{},
	}
}

// FetchCategories получает данные о категориях.
func (c *Client) FetchCategories() ([]models.GQLCategory, error) {
	query := `query GetCategoriesForYML { filterCategory { id name parentCategory { id } } }`
	requestBody, _ := json.Marshal(models.GraphQLRequest{Query: query})
	resp, err := c.executeRequest(requestBody)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseBody, _ := io.ReadAll(resp.Body)
	var categoryResp models.GQLCategoryResponse
	if err := json.Unmarshal(responseBody, &categoryResp); err != nil {
		return nil, fmt.Errorf("ошибка распаковки JSON категорий: %w", err)
	}
	return categoryResp.Data.FilterCategory, nil
}

// FetchProducts получает данные о товарах.
func (c *Client) FetchProducts() ([]models.GQLProduct, error) {
	query := `query GetProductsForYML {
			  filterProduct(filter: {statusId: 1}) {
				id name price oldPrice vendorCode
				category { id parentCategory { page { url } } page { url } }
				images { image } page { url }
			  }
			}`
	requestBody, _ := json.Marshal(models.GraphQLRequest{Query: query})
	resp, err := c.executeRequest(requestBody)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseBody, _ := io.ReadAll(resp.Body)
	var productResp models.GQLProductResponse
	if err := json.Unmarshal(responseBody, &productResp); err != nil {
		return nil, fmt.Errorf("ошибка распаковки JSON товаров: %w", err)
	}
	return productResp.Data.FilterProduct, nil
}

// executeRequest выполняет основной HTTP POST запрос.
func (c *Client) executeRequest(body []byte) (*http.Response, error) {
	resp, err := http.Post(c.apiURL, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("ошибка при отправке HTTP-запроса: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		errorBody, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		return nil, fmt.Errorf("получен неверный статус-код: %d, тело ответа: %s", resp.StatusCode, string(errorBody))
	}
	return resp, nil
}
