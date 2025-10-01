package models

import "encoding/xml"

// --- Структуры для GraphQL запроса ---

// GraphQLRequest Структура для GraphQL запроса
type GraphQLRequest struct {
	Query string `json:"query"`
}

// GQLPage Вложенная структура для страницы
type GQLPage struct {
	URL string `json:"url"`
}

// --- Структуры для данных из API (Категории) ---

// GQLParentCategory Вложенная структура для родительской категории
type GQLParentCategory struct {
	ID   int      `json:"id"`
	Page *GQLPage `json:"page"`
}

// GQLCategory Основная структура для категорий
type GQLCategory struct {
	ID     int                `json:"id"`
	Name   string             `json:"name"`
	Page   *GQLPage           `json:"page"`
	Parent *GQLParentCategory `json:"parentCategory"`
}

// GQLCategoryResponse Ответ от GraphQL для категорий
type GQLCategoryResponse struct {
	Data struct {
		FilterCategory []GQLCategory `json:"filterCategory"`
	} `json:"data"`
}

// --- Структуры для данных из API (Товары) ---

// GQLImage Вложенная структура для изображений
type GQLImage struct {
	URL string `json:"image"`
}

// GQLProduct Основная структура для товаров
type GQLProduct struct {
	ID         int          `json:"id"`
	Name       string       `json:"name"`
	Price      float64      `json:"price"`
	OldPrice   float64      `json:"oldPrice"`
	VendorCode string       `json:"vendorCode"`
	Category   *GQLCategory `json:"category"`
	Images     []GQLImage   `json:"images"`
	Page       *GQLPage     `json:"page"`
}

// GQLProductResponse Ответ от GraphQL для товаров
type GQLProductResponse struct {
	Data struct {
		FilterProduct []GQLProduct `json:"filterProduct"`
	} `json:"data"`
}

// --- Структуры для генерации YML файла ---

// YMLCategory Корневой элемент YML
type YMLCategory struct {
	XMLName  xml.Name `xml:"category"`
	ID       int      `xml:"id,attr"`
	ParentID int      `xml:"parentId,attr,omitempty"`
	Name     string   `xml:",chardata"`
}

// YMLOffer Элемент предложения (товара) в YML
type YMLOffer struct {
	XMLName    xml.Name `xml:"offer"`
	ID         int      `xml:"id,attr"`
	Name       string   `xml:"name"`
	Price      float64  `xml:"price"`
	OldPrice   float64  `xml:"oldprice,omitempty"`
	CategoryID int      `xml:"categoryId"`
	URL        string   `xml:"url"`
	Pictures   []string `xml:"picture,omitempty"`
	BarCode    string   `xml:"barcode,omitempty"`
}

// YMLCurrency Элемент валюты в YML
type YMLCurrency struct {
	ID   string `xml:"id,attr"`
	Rate string `xml:"rate,attr"`
}

// YMLShop Элемент магазина в YML
type YMLShop struct {
	XMLName    xml.Name      `xml:"shop"`
	Categories []YMLCategory `xml:"categories>category"`
	Offers     []YMLOffer    `xml:"offers>offer"`
}

// YMLCatalog Корневой элемент YML файла
type YMLCatalog struct {
	XMLName xml.Name `xml:"yml_catalog"`
	Date    string   `xml:"date,attr"`
	Shop    YMLShop  `xml:"shop"`
}
