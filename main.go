package main

import "fmt"

func main() {
	// Create a new blockchain
	blockchain := NewBlockchain()

	// Create some sample transactions
	transaction1 := Transaction{
		TransactionID:   "1",
		ProductID:       "ABC123",
		ProductName:     "Widget",
		ProductType:     "Electronics",
		Quantity:        10,
		UnitOfMeasure:   "pieces",
		TransactionType: "Transfer",
		FromParty:       Party{ID: "MANUF001", Name: "Manufacturer X", Role: "Manufacturer"},
		ToParty:         Party{ID: "DIST001", Name: "Distributor Y", Role: "Distributor"},
		Timestamp:       1623456780,
		LocationData:    Location{Latitude: 37.7749, Longitude: -122.4194},
		Conditions:      Conditions{Temperature: 25, Humidity: 60},
		Notes:           "Regular shipment",
	}

	transaction2 := Transaction{
		TransactionID:   "2",
		ProductID:       "XYZ789",
		ProductName:     "Gadget",
		ProductType:     "Electronics",
		Quantity:        5,
		UnitOfMeasure:   "pieces",
		TransactionType: "Transfer",
		FromParty:       Party{ID: "DIST001", Name: "Distributor Y", Role: "Distributor"},
		ToParty:         Party{ID: "RET001", Name: "Retailer Z", Role: "Retailer"},
		Timestamp:       1623543210,
		LocationData:    Location{Latitude: 40.7128, Longitude: -74.0060},
		Conditions:      Conditions{Temperature: 20, Humidity: 55},
		Notes:           "Urgent delivery",
	}

	// Add the transactions to the blockchain
	blockchain.AddBlock([]Transaction{transaction1})
	blockchain.AddBlock([]Transaction{transaction2})

	// Print the blockchain
	for _, block := range blockchain.Blocks {
		fmt.Printf("Block %d:\n", block.BlockNumber)
		fmt.Printf("Version: %s\n", block.Version)
		fmt.Printf("Previous Hash: %s\n", block.PreviousHash)
		fmt.Printf("Timestamp: %d\n", block.Timestamp)
		fmt.Printf("Transactions:\n")
		for _, tx := range block.Transactions {
			fmt.Printf("  - Transaction ID: %s\n", tx.TransactionID)
			fmt.Printf("    Product ID: %s\n", tx.ProductID)
			fmt.Printf("    Product Name: %s\n", tx.ProductName)
			fmt.Printf("    Quantity: %d\n", tx.Quantity)
			// Print other transaction details as needed
		}
		fmt.Println()
	}

	// Validate the blockchain
	isValid := blockchain.IsValid()
	fmt.Printf("Is blockchain valid? %t\n", isValid)
}
