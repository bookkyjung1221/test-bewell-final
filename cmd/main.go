package main

import (
	"bewell_test/internal/model"
	"bewell_test/internal/service"

	"encoding/json"
	"fmt"
	"log"
)

func main() {
	// Create a new order service
	orderService := service.NewOrderService()

	// Read input data (this would come from an API or file in production)
	// For demonstration, we'll use a hardcoded example:
	inputOrdersJSON := `[
{
"no": 1,
"platformProductId": "FG0A-CLEAR-OPPOA3/%20xFG0A-CLEAR-OPPOA3-B/FG0A-MAT",
"qty": 1,
"unitPrice": 120,
"totalPrice": 120
}
	]`

	// Parse the input JSON
	var inputOrders []model.InputOrder
	if err := json.Unmarshal([]byte(inputOrdersJSON), &inputOrders); err != nil {
		log.Fatalf("Failed to parse input orders: %v", err)
	}

	// Process orders
	cleanedOrders := orderService.ProcessOrders(inputOrders)

	// Output the results as JSON
	outputJSON, err := json.MarshalIndent(cleanedOrders, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal output: %v", err)
	}

	fmt.Println(string(outputJSON))

	// In a real application, you might want to persist or send this data somewhere
}
