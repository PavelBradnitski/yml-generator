package generator

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/PavelBradnitski/yml-generator/internal/models"
)

// Тестируем успешное создание YML файла.
func TestWriteYMLFile_Success(t *testing.T) {
	// 1. Arrange (Подготовка)
	catalog := models.YMLCatalog{
		Date: "2025-10-01",
		Shop: models.YMLShop{
			Categories: []models.YMLCategory{
				{ID: 1, Name: "Тестовая категория"},
			},
			Offers: []models.YMLOffer{
				{ID: 123, Name: "Тестовый товар", Price: 100, CategoryID: 1, URL: "http://test.url"},
			},
		},
	}

	tempDir := t.TempDir()

	filePath := filepath.Join(tempDir, "shop.yml")

	// 2. Act (Действие)
	err := WriteYMLFile(catalog, filePath)

	// 3. Assert (Проверка)
	if err != nil {
		t.Fatalf("WriteYMLFile() вернула неожиданную ошибку: %v", err)
	}

	contentBytes, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("Не удалось прочитать созданный тестовый файл: %v", err)
	}

	// Определяем, каким должно быть содержимое файла.
	expectedContent := `<?xml version="1.0" encoding="UTF-8"?>
<yml_catalog date="2025-10-01">
  <shop>
    <categories>
      <category id="1">Тестовая категория</category>
    </categories>
    <offers>
      <offer id="123">
        <name>Тестовый товар</name>
        <price>100</price>
        <categoryId>1</categoryId>
        <url>http://test.url</url>
      </offer>
    </offers>
  </shop>
</yml_catalog>`

	// Сравниваем реальное содержимое с ожидаемым.
	if strings.TrimSpace(string(contentBytes)) != strings.TrimSpace(expectedContent) {
		t.Errorf("Содержимое файла не совпадает с ожидаемым.\nОжидали:\n%s\n\nПолучили:\n%s", expectedContent, string(contentBytes))
	}
}

// Тестируем случай, когда запись файла невозможна (неверный путь).
func TestWriteYMLFile_Error(t *testing.T) {
	// 1. Arrange
	catalog := models.YMLCatalog{}
	invalidPath := filepath.Join(t.TempDir(), "nonexistent", "shop.yml")

	// 2. Act
	err := WriteYMLFile(catalog, invalidPath)

	// 3. Assert
	// Мы ожидаем получить ошибку. Если ошибки нет - это провал теста.
	if err == nil {
		t.Fatal("Ожидали получить ошибку при записи по неверному пути, но не получили ее.")
	}
}
