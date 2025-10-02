package converter

import (
	"strings"
	"time"

	"github.com/PavelBradnitski/yml-generator/internal/config"
	"github.com/PavelBradnitski/yml-generator/internal/models"
)

// BuildYMLCatalog преобразует срезы категорий и товаров из API в готовую структуру YMLCatalog.
func BuildYMLCatalog(gqlCategories []models.GQLCategory, gqlProducts []models.GQLProduct, cfg *config.Config) models.YMLCatalog {
	// Преобразуем категории
	var ymlCategories []models.YMLCategory
	for _, cat := range gqlCategories {
		category := models.YMLCategory{ID: cat.ID, Name: cat.Name}
		if cat.Parent != nil && cat.ID != cat.Parent.ID {
			category.ParentID = cat.Parent.ID
		}
		ymlCategories = append(ymlCategories, category)
	}

	// Преобразуем товары
	var ymlOffers []models.YMLOffer
	for _, p := range gqlProducts {
		var urlBuilder strings.Builder
		urlBuilder.WriteString(cfg.BaseURL)
		if p.Category != nil && p.Category.Parent != nil && p.Category.Parent.Page != nil {
			urlBuilder.WriteString("/" + p.Category.Parent.Page.URL)
		}
		if p.Category != nil && p.Category.Page != nil {
			urlBuilder.WriteString("/" + p.Category.Page.URL)
		}
		if p.Page != nil {
			urlBuilder.WriteString("/" + p.Page.URL)
		}

		offer := models.YMLOffer{
			ID:      p.ID,
			Name:    p.Name,
			BarCode: p.VendorCode,
			URL:     urlBuilder.String(),
			Price:   p.Price,
		}
		if p.OldPrice > 0 {
			offer.OldPrice = p.OldPrice
		}
		if p.Category != nil {
			offer.CategoryID = p.Category.ID
		}
		for _, img := range p.Images {
			offer.Pictures = append(offer.Pictures, cfg.BaseURL+"/pics/items/"+img.URL)
		}
		ymlOffers = append(ymlOffers, offer)
	}

	// Собираем итоговую структуру
	catalog := models.YMLCatalog{
		Date: time.Now().Format("2006-01-02 15:04"),
		Shop: models.YMLShop{
			Categories: ymlCategories,
			Offers:     ymlOffers,
		},
	}

	return catalog
}
