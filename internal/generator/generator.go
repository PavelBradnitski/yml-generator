package generator

import (
	"encoding/xml"
	"fmt"
	"os"

	"github.com/PavelBradnitski/yml-generator/internal/models"
)

// WriteYMLFile маршализирует структуру YMLCatalog в XML и записывает ее в файл.
func WriteYMLFile(catalog models.YMLCatalog, filename string) error {
	file, err := xml.MarshalIndent(catalog, "", "  ")
	if err != nil {
		return fmt.Errorf("ошибка при маршалинге XML: %w", err)
	}

	output := []byte(xml.Header)
	output = append(output, file...)

	err = os.WriteFile(filename, output, 0644)
	if err != nil {
		return fmt.Errorf("ошибка при записи файла: %w", err)
	}

	return nil
}
