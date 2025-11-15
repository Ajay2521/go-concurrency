package main

import (
	"fmt"
	"time"
)

type Order struct {
	ID       int
	PrepTime time.Duration
}

func processOrder(order Order) {
	// Print order start message
	fmt.Printf("📝 Order %d: Started processing\n", order.ID)

	// Simulate order processing time (blocking operation)
	time.Sleep(order.PrepTime)

	// Print order completion message with time taken
	fmt.Printf("✅ Order %d : Ready for pickup! Time taken: %v\n\n", order.ID, order.PrepTime)
}

func main() {
	fmt.Println("🏪 Sequential Synchronous Order Processing System")
	fmt.Println("⏰ Processing started\n")

	// Record start time for total processing calculation
	startTime := time.Now()

	// Create orders
	orders := []Order{
		{ID: 1, PrepTime: 2 * time.Second},
		{ID: 2, PrepTime: 3 * time.Second},
		{ID: 3, PrepTime: 1 * time.Second},
		{ID: 4, PrepTime: 4 * time.Second},
		{ID: 5, PrepTime: 2 * time.Second},
	}

	// Process orders sequentially (one after another)
	for _, order := range orders {
		processOrder(order)
	}

	// Calculate and display total processing time
	fmt.Printf("⏱️  Total processing time: %v\n", time.Since(startTime)) // 2 + 3 + 1 + 4 + 2 = 12 seconds
	fmt.Println("🔄 Note: Orders processed sequentially - one after another")
}
