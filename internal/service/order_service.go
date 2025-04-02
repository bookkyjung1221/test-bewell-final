package service

import (
	"bewell_test/internal/model"
	"bewell_test/internal/parser"

	"strings"
)

// OrderService handles the business logic for processing orders
type OrderService struct {
	productParser *parser.ProductParser
	// Map of texture to cleaner product ID
	cleanerMap map[string]string
}

// NewOrderService creates a new OrderService
func NewOrderService() *OrderService {
	// Initialize with predefined cleaners for each texture type
	cleanerMap := map[string]string{
		"CLEAR":   "CLEAR-CLEANNER",
		"MATTE":   "MATTE-CLEANNER",
		"PRIVACY": "PRIVACY-CLEANNER",
		// Add more texture-cleaner mappings as needed
	}

	return &OrderService{
		productParser: parser.NewProductParser(),
		cleanerMap:    cleanerMap,
	}
}

// ProcessOrders transforms input orders to cleaned orders with complementary items
func (s *OrderService) ProcessOrders(inputOrders []model.InputOrder) []model.CleanedOrder {
	var cleanedOrders []model.CleanedOrder

	// Tracks the next order number to use
	nextOrderNo := 1

	// Maps to track total quantities for complementary items
	wipingClothQty := 0
	cleanerQuantities := make(map[string]int)

	// Process each input order
	for _, inputOrder := range inputOrders {
		// Parse the products from the platform product ID
		parsedProducts := s.productParser.ParseProducts(inputOrder.PlatformProductId)

		if len(parsedProducts) == 0 {
			// Skip this order if no valid product could be parsed
			continue
		}

		// Calculate unit price if this is a bundle
		unitPrice := inputOrder.UnitPrice

		if len(parsedProducts) >= 1 {
			sum := 0

			for _, v := range parsedProducts {
				sum += v.Qty
			}

			unitPrice = inputOrder.UnitPrice / float64(sum)

		}

		// Process each parsed product
		for _, product := range parsedProducts {
			// Calculate the actual quantity based on input order and parsed product
			actualQty := product.Qty * inputOrder.Qty

			// Add the main product
			cleanedOrders = append(cleanedOrders, model.CleanedOrder{
				No:         nextOrderNo,
				ProductId:  product.ProductId,
				MaterialId: product.MaterialId,
				ModelId:    product.ModelId,
				Qty:        actualQty,
				UnitPrice:  unitPrice,
				TotalPrice: unitPrice * float64(actualQty),
			})
			nextOrderNo++

			// Track quantities for complementary items
			wipingClothQty += actualQty

			// Extract texture ID for cleaner products
			textureParts := strings.Split(product.MaterialId, "-")
			if len(textureParts) >= 2 {
				textureId := textureParts[1]
				if cleanerProductId, exists := s.cleanerMap[textureId]; exists {
					cleanerQuantities[cleanerProductId] += actualQty
				}
			}
		}
	}

	// Add wiping cloth as complementary item
	if wipingClothQty > 0 {
		cleanedOrders = append(cleanedOrders, model.CleanedOrder{
			No:         nextOrderNo,
			ProductId:  "WIPING-CLOTH",
			Qty:        wipingClothQty,
			UnitPrice:  0,
			TotalPrice: 0,
		})
		nextOrderNo++
	}

	// Add cleaner products as complementary items
	for cleanerProductId, qty := range cleanerQuantities {
		cleanedOrders = append(cleanedOrders, model.CleanedOrder{
			No:         nextOrderNo,
			ProductId:  cleanerProductId,
			Qty:        qty,
			UnitPrice:  0,
			TotalPrice: 0,
		})
		nextOrderNo++
	}

	return cleanedOrders
}
