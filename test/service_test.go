package test

import (
	"bewell_test/internal/model"
	"bewell_test/internal/service"

	"encoding/json"
	"reflect"
	"testing"
)

func TestProcessOrders_Case1(t *testing.T) {
	// Case 1: Only one product
	inputJSON := `[
		{
			"no": 1,
			"platformProductId": "FG0A-CLEAR-IPHONE16PROMAX",
			"qty": 2,
			"unitPrice": 50,
			"totalPrice": 100
		}
	]`

	expectedJSON := `[
		{
			"no": 1,
			"productId": "FG0A-CLEAR-IPHONE16PROMAX",
			"materialId": "FG0A-CLEAR",
			"modelId": "IPHONE16PROMAX",
			"qty": 2,
			"unitPrice": 50.00,
			"totalPrice": 100.00
		},
		{
			"no": 2,
			"productId": "WIPING-CLOTH",
			"qty": 2,
			"unitPrice": 0.00,
			"totalPrice": 0.00
		},
		{
			"no": 3,
			"productId": "CLEAR-CLEANNER",
			"qty": 2,
			"unitPrice": 0.00,
			"totalPrice": 0.00
		}
	]`

	runTestCase(t, inputJSON, expectedJSON)
}

func TestProcessOrders_Case2(t *testing.T) {
	// Case 2: One product with wrong prefix
	inputJSON := `[
		{
			"no": 1,
			"platformProductId": "x2-3&FG0A-CLEAR-IPHONE16PROMAX",
			"qty": 2,
			"unitPrice": 50,
			"totalPrice": 100
		}
	]`

	expectedJSON := `[
		{
			"no": 1,
			"productId": "FG0A-CLEAR-IPHONE16PROMAX",
			"materialId": "FG0A-CLEAR",
			"modelId": "IPHONE16PROMAX",
			"qty": 2,
			"unitPrice": 50.00,
			"totalPrice": 100.00
		},
		{
			"no": 2,
			"productId": "WIPING-CLOTH",
			"qty": 2,
			"unitPrice": 0.00,
			"totalPrice": 0.00
		},
		{
			"no": 3,
			"productId": "CLEAR-CLEANNER",
			"qty": 2,
			"unitPrice": 0.00,
			"totalPrice": 0.00
		}
	]`

	runTestCase(t, inputJSON, expectedJSON)
}

func TestProcessOrders_Case3(t *testing.T) {
	// Case 3: One product with wrong prefix and has * symbol that indicates the quantity
	inputJSON := `[
		{
			"no": 1,
			"platformProductId": "x2-3&FG0A-MATTE-IPHONE16PROMAX*3",
			"qty": 1,
			"unitPrice": 90,
			"totalPrice": 90
		}
	]`

	expectedJSON := `[
		{
			"no": 1,
			"productId": "FG0A-MATTE-IPHONE16PROMAX",
			"materialId": "FG0A-MATTE",
			"modelId": "IPHONE16PROMAX",
			"qty": 3,
			"unitPrice": 30.00,
			"totalPrice": 90.00
		},
		{
			"no": 2,
			"productId": "WIPING-CLOTH",
			"qty": 3,
			"unitPrice": 0.00,
			"totalPrice": 0.00
		},
		{
			"no": 3,
			"productId": "MATTE-CLEANNER",
			"qty": 3,
			"unitPrice": 0.00,
			"totalPrice": 0.00
		}
	]`

	runTestCase(t, inputJSON, expectedJSON)
}

func TestProcessOrders_Case4(t *testing.T) {
	// Case 4: One bundle product with wrong prefix and split by / symbol into two product
	inputJSON := `[
		{
			"no": 1,
			"platformProductId": "FG0A-CLEAR-OPPOA3/%20xFG0A-CLEAR-OPPOA3-B",
			"qty": 1,
			"unitPrice": 80,
			"totalPrice": 80
		}
	]`

	expectedJSON := `[
		{
			"no": 1,
			"productId": "FG0A-CLEAR-OPPOA3",
			"materialId": "FG0A-CLEAR",
			"modelId": "OPPOA3",
			"qty": 1,
			"unitPrice": 40.00,
			"totalPrice": 40.00
		},
		{
			"no": 2,
			"productId": "FG0A-CLEAR-OPPOA3-B",
			"materialId": "FG0A-CLEAR",
			"modelId": "OPPOA3-B",
			"qty": 1,
			"unitPrice": 40.00,
			"totalPrice": 40.00
		},
		{
			"no": 3,
			"productId": "WIPING-CLOTH",
			"qty": 2,
			"unitPrice": 0.00,
			"totalPrice": 0.00
		},
		{
			"no": 4,
			"productId": "CLEAR-CLEANNER",
			"qty": 2,
			"unitPrice": 0.00,
			"totalPrice": 0.00
		}
	]`

	runTestCase(t, inputJSON, expectedJSON)
}

func TestProcessOrders_Case6(t *testing.T) {
	// Case 6: One bundle product with wrong prefix and have / symbol and * symbol

	inputJSON := `[
		{
"no": 1,
"platformProductId": "--FG0A-CLEAR-OPPOA3*2/FG0A-MATTE-OPPOA3",
"qty": 1,
"unitPrice": 120,
"totalPrice": 120
}

		
	]`

	expectedJSON := `[
		{
"no": 1,
"productId": "FG0A-CLEAR-OPPOA3",
"materialId": "FG0A-CLEAR",
"modelId": "OPPOA3",
"qty": 2,
"unitPrice": 40.00,
"totalPrice": 80.00
},
{
"no": 2,
"productId": "FG0A-MATTE-OPPOA3",
"materialId": "FG0A-MATTE",
"modelId": "OPPOA3",
"qty": 1,
"unitPrice": 40.00,
"totalPrice": 40.00
},
{
"no": 3,
"productId": "WIPING-CLOTH",
"qty": 3,
"unitPrice": 0.00,
"totalPrice": 0.00
},
{
"no": 4,
"productId": "CLEAR-CLEANNER",
"qty": 2,
"unitPrice": 0.00,
"totalPrice": 0.00
},
{
"no": 5,
"productId": "MATTE-CLEANNER",
"qty": 1,
"unitPrice": 0.00,
"totalPrice": 0.00
}
	]`

	runTestCase(t, inputJSON, expectedJSON)
}
func TestProcessOrders_Case7(t *testing.T) {
	// Case 7: one product and one bundle product with wrong prefix and have / symbol and * symbol
	inputJSON := `[
		{
			"no": 1,
			"platformProductId": "--FG0A-CLEAR-OPPOA3*2/FG0A-MATTE-OPPOA3*2",
			"qty": 1,
			"unitPrice": 160,
			"totalPrice": 160
		},
		{
			"no": 2,
			"platformProductId": "FG0A-PRIVACY-IPHONE16PROMAX",
			"qty": 1,
			"unitPrice": 50,
			"totalPrice": 50
		}
	]`

	expectedJSON := `[
		{
			"no": 1,
			"productId": "FG0A-CLEAR-OPPOA3",
			"materialId": "FG0A-CLEAR",
			"modelId": "OPPOA3",
			"qty": 2,
			"unitPrice": 40.00,
			"totalPrice": 80.00
		},
		{
			"no": 2,
			"productId": "FG0A-MATTE-OPPOA3",
			"materialId": "FG0A-MATTE",
			"modelId": "OPPOA3",
			"qty": 2,
			"unitPrice": 40.00,
			"totalPrice": 80.00
		},
		{
			"no": 3,
			"productId": "FG0A-PRIVACY-IPHONE16PROMAX",
			"materialId": "FG0A-PRIVACY",
			"modelId": "IPHONE16PROMAX",
			"qty": 1,
			"unitPrice": 50.00,
			"totalPrice": 50.00
		},
		{
			"no": 4,
			"productId": "WIPING-CLOTH",
			"qty": 5,
			"unitPrice": 0.00,
			"totalPrice": 0.00
		},
		{
			"no": 5,
			"productId": "CLEAR-CLEANNER",
			"qty": 2,
			"unitPrice": 0.00,
			"totalPrice": 0.00
		},
		{
			"no": 6,
			"productId": "MATTE-CLEANNER",
			"qty": 2,
			"unitPrice": 0.00,
			"totalPrice": 0.00
		},
		{
			"no": 7,
			"productId": "PRIVACY-CLEANNER",
			"qty": 1,
			"unitPrice": 0.00,
			"totalPrice": 0.00
		}
	]`

	runTestCase(t, inputJSON, expectedJSON)
}

// Helper function to run a test case
func runTestCase(t *testing.T, inputJSON, expectedJSON string) {
	// Parse input and expected output
	var inputOrders []model.InputOrder
	if err := json.Unmarshal([]byte(inputJSON), &inputOrders); err != nil {
		t.Fatalf("Failed to parse input JSON: %v", err)
	}

	var expectedOrders []model.CleanedOrder
	if err := json.Unmarshal([]byte(expectedJSON), &expectedOrders); err != nil {
		t.Fatalf("Failed to parse expected JSON: %v", err)
	}

	// Create service and process orders
	orderService := service.NewOrderService()
	actualOrders := orderService.ProcessOrders(inputOrders)

	// Compare results
	if !reflect.DeepEqual(actualOrders, expectedOrders) {
		// For better debugging, print both actual and expected as JSON
		actualJSON, _ := json.MarshalIndent(actualOrders, "", "  ")
		t.Errorf("Test failed.\nExpected: %s\nActual: %s", expectedJSON, string(actualJSON))
	}
}
