package converter

import (
	"testing"

	"github.com/PavelBradnitski/yml-generator/internal/models" // ИЗМЕНЕНИЕ: Правильный путь импорта
)

func TestBuildYMLCatalog(t *testing.T) {
	// 1. Arrange (Подготовка): Создаем входные данные для теста
	gqlCategories := []models.GQLCategory{
		{ID: 1, Name: "Главная"},
	}

	gqlProducts := []models.GQLProduct{
		{
			ID:         123,
			Name:       "Тестовый Товар",
			Price:      100.50,
			VendorCode: "T-001",
			Category: &models.GQLCategory{
				ID:   1,
				Page: &models.GQLPage{URL: "category"},
				Parent: &models.GQLParentCategory{
					Page: &models.GQLPage{URL: "parent-category"},
				},
			},
			Page: &models.GQLPage{URL: "test-product"},
		},
		{
			ID:       456,
			Name:     "Товар со старой ценой",
			Price:    80.0,
			OldPrice: 120.0,
			Category: &models.GQLCategory{ID: 1},
		},
	}

	// 2. Act (Действие): Вызываем функцию, которую тестируем
	result := BuildYMLCatalog(gqlCategories, gqlProducts)

	// 3. Assert (Проверка): Проверяем, что результат соответствует ожиданиям
	if len(result.Shop.Offers) != 2 {
		t.Fatalf("Ожидали получить 2 товара, а получили %d", len(result.Shop.Offers))
	}

	expectedURL := "https://demo.beseller.com/parent-category/category/test-product"
	if result.Shop.Offers[0].URL != expectedURL {
		t.Errorf("URL собран неверно.\nОжидали: '%s'\nПолучили: '%s'", expectedURL, result.Shop.Offers[0].URL)
	}

	if result.Shop.Offers[0].OldPrice != 0 {
		t.Errorf("У первого товара не должно быть старой цены, а она есть: %f", result.Shop.Offers[0].OldPrice)
	}

	if result.Shop.Offers[1].OldPrice != 120.0 {
		t.Errorf("У второго товара ожидалась старая цена 120.0, а получили: %f", result.Shop.Offers[1].OldPrice)
	}
}
