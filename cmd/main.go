package main

import (
	"fmt"
	"log"

	"github.com/PavelBradnitski/yml-generator/internal/client"
	"github.com/PavelBradnitski/yml-generator/internal/converter"
	"github.com/PavelBradnitski/yml-generator/internal/generator"
)

const apiURL = "https://demo.beseller.com/graphql?token=c2b05afc61b8d243e0c10b46a677cbf4eb98d09c"

func main() {
	// 1. Получение данных
	apiClient := client.NewClient(apiURL)

	fmt.Println("Получение категорий...")
	gqlCategories, err := apiClient.FetchCategories()
	if err != nil {
		log.Fatalf("Не удалось получить категории: %s", err)
	}
	fmt.Printf("Получено %d категорий.\n", len(gqlCategories))

	fmt.Println("Получение товаров...")
	gqlProducts, err := apiClient.FetchProducts()
	if err != nil {
		log.Fatalf("Не удалось получить товары: %s", err)
	}
	fmt.Printf("Получено %d товаров для YML файла.\n", len(gqlProducts))

	// ... остальной код без изменений ...
	fmt.Println("Формирование YML структуры...")
	catalog := converter.BuildYMLCatalog(gqlCategories, gqlProducts)

	fmt.Println("Запись в файл shop.yml...")
	err = generator.WriteYMLFile(catalog, "shop.yml")
	if err != nil {
		log.Fatalf("Не удалось создать YML файл: %s", err)
	}

	fmt.Println("\nГотово! Файл shop.yml успешно создан.")
}
