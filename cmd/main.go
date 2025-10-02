package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/PavelBradnitski/yml-generator/internal/client"
	"github.com/PavelBradnitski/yml-generator/internal/config"
	"github.com/PavelBradnitski/yml-generator/internal/converter"
	"github.com/PavelBradnitski/yml-generator/internal/generator"
	"github.com/PavelBradnitski/yml-generator/internal/models"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Ошибка при загрузке конфигурации: %s", err)
	}

	apiURL := fmt.Sprintf("%s?token=%s", cfg.APIURL, cfg.APIToken)
	apiClient := client.NewClient(apiURL)

	var wg sync.WaitGroup
	var errCat, errProd error
	var gqlCategories []models.GQLCategory
	var gqlProducts []models.GQLProduct
	fmt.Println("Запуск получения данных...")

	wg.Add(2)

	// Горутина для получения категорий
	go func() {
		defer wg.Done()
		fmt.Println("Получение категорий...")
		gqlCategories, errCat = apiClient.FetchCategories()
		if errCat == nil {
			fmt.Printf("Получено %d категорий.\n", len(gqlCategories))
		}
	}()

	// Горутина для получения товаров
	go func() {
		defer wg.Done()
		fmt.Println("Получение товаров...")
		gqlProducts, errProd = apiClient.FetchProducts()
		if errProd == nil {
			fmt.Printf("Получено %d товаров для YML файла.\n", len(gqlProducts))
		}
	}()

	wg.Wait()

	if errCat != nil {
		log.Fatalf("Не удалось получить категории: %s", errCat)
	}
	if errProd != nil {
		log.Fatalf("Не удалось получить товары: %s", errProd)
	}

	fmt.Println("Формирование YML структуры...")
	catalog := converter.BuildYMLCatalog(gqlCategories, gqlProducts, cfg)

	fmt.Printf("Запись в файл %s...\n", cfg.OutputFilename)
	err = generator.WriteYMLFile(catalog, cfg.OutputFilename)
	if err != nil {
		log.Fatalf("Не удалось создать YML файл: %s", err)
	}

	fmt.Println("\nГотово! Файл shop.yml успешно создан.")
}
