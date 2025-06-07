package products

import (
	"fmt"

	"gorm.io/gorm"
)

func InitDbWithMockProducts(db *gorm.DB) error {
	products := make([]Product, 0, 100)
	for i := 1; i <= 100; i++ {
		products = append(products, Product{
			ID:          fmt.Sprintf("%d", i),
			Name:        fmt.Sprintf("Product %d", i),
			Description: fmt.Sprintf("Description for Product %d", i),
		})
	}

	// Insert mock products into the database if they don't exist
	for _, product := range products {
		var count int64
		if err := db.Model(&Product{}).Where("id = ?", product.ID).Count(&count).Error; err != nil {
			return fmt.Errorf("failed to check if product exists: %w", err)
		}
		if count == 0 {
			if err := db.Create(&product).Error; err != nil {
				return fmt.Errorf("failed to insert mock product: %w", err)
			}
		}
	}

	return nil
}
