package main

import (
	"fmt"
	"log/slog"
	"os"
	"sync"

	"github.com/PavelBradnitski/yml-generator/internal/client"
	"github.com/PavelBradnitski/yml-generator/internal/config"
	"github.com/PavelBradnitski/yml-generator/internal/converter"
	"github.com/PavelBradnitski/yml-generator/internal/generator"
	"github.com/PavelBradnitski/yml-generator/internal/models"
)

func main() {
	opts := &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelInfo,
	}
	handler := slog.NewTextHandler(os.Stdout, opts)
	logger := slog.New(handler)
	slog.SetDefault(logger)
	cfg, err := config.Load()
	if err != nil {
		slog.Error("Ошибка при загрузке конфигурации", "error", err)
	}
	slog.Info("Конфигурация успешно загружена")

	apiURL := fmt.Sprintf("%s?token=%s", cfg.APIURL, cfg.APIToken)
	apiClient := client.NewClient(apiURL)

	var wg sync.WaitGroup
	var errCat, errProd error
	var gqlCategories []models.GQLCategory
	var gqlProducts []models.GQLProduct
	slog.Info("Конфигурация успешно загружена")

	wg.Add(2)

	// Горутина для получения категорий
	go func() {
		defer wg.Done()
		fmt.Println("Получение категорий...")
		gqlCategories, errCat = apiClient.FetchCategories()
		if errCat == nil {
			slog.Info("Категории успешно получены", "count", len(gqlCategories))
		}
	}()

	// Горутина для получения товаров
	go func() {
		defer wg.Done()
		fmt.Println("Получение товаров...")
		gqlProducts, errProd = apiClient.FetchProducts()
		if errProd == nil {
			slog.Info("Товары успешно получены", "count", len(gqlProducts))
		}
	}()

	wg.Wait()
	slog.Debug("Все горутины завершили работу")

	if errCat != nil {
		slog.Error("Не удалось получить категории", "error", errCat)
		os.Exit(1)
	}
	if errProd != nil {
		slog.Error("Не удалось получить товары", "error", errProd)
		os.Exit(1)
	}

	slog.Info("Формирование YML структуры...")
	catalog := converter.BuildYMLCatalog(gqlCategories, gqlProducts, cfg)

	slog.Info("Запись YML файла", "filename", cfg.OutputFilename)
	err = generator.WriteYMLFile(catalog, cfg.OutputFilename)
	if err != nil {
		slog.Error("Не удалось создать YML файл", "error", err, "filename", cfg.OutputFilename)
		os.Exit(1)
	}

	slog.Info("Генерация успешно завершена", "filename", cfg.OutputFilename)
}
