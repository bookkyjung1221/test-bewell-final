package parser

import (
	"regexp"
	"strconv"
	"strings"
)

// ProductParser handles parsing of platform-specific product IDs
type ProductParser struct {
	// Regex to extract the core product code pattern (e.g., FG0A-CLEAR-IPHONE16PROMAX)
	productCodeRegex *regexp.Regexp
	// Regex to extract quantity when marked with * symbol
	quantityRegex *regexp.Regexp
}

// NewProductParser creates a new product parser with compiled regexes
func NewProductParser() *ProductParser {
	return &ProductParser{
		productCodeRegex: regexp.MustCompile(`(FG0[0-9A-Z]\-[A-Z]+\-[A-Z0-9\-]+)(?:\*(\d+))?`),
		quantityRegex:    regexp.MustCompile(`\*(\d+)$`),
	}
}

// ParseProducts parses a platform product ID into one or more standardized products
// Returns a slice of parsed products with their respective quantities
func (p *ProductParser) ParseProducts(platformProductId string) []struct {
	ProductId  string
	MaterialId string
	ModelId    string
	Qty        int
} {
	var result []struct {
		ProductId  string
		MaterialId string
		ModelId    string
		Qty        int
	}

	// Check if this is a bundle (products separated by /)
	productStrings := strings.Split(platformProductId, "/")

	for _, productStr := range productStrings {
		matches := p.productCodeRegex.FindStringSubmatch(productStr)
		if len(matches) < 2 {
			continue // Skip if no valid product code found
		}

		productCode := matches[1]
		// Parse parts to extract material and model IDs
		parts := strings.Split(productCode, "-")
		if len(parts) < 3 {
			continue // Skip if format is not as expected
		}

		materialId := parts[0] + "-" + parts[1]
		modelId := strings.Join(parts[2:], "-")

		// Parse quantity if specified with * symbol
		qty := 1
		if len(matches) > 2 && matches[2] != "" {
			if parsedQty, err := strconv.Atoi(matches[2]); err == nil && parsedQty > 0 {
				qty = parsedQty
			}
		}

		result = append(result, struct {
			ProductId  string
			MaterialId string
			ModelId    string
			Qty        int
		}{
			ProductId:  productCode,
			MaterialId: materialId,
			ModelId:    modelId,
			Qty:        qty,
		})
	}

	return result
}
