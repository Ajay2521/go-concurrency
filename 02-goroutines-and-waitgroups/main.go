package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

type Order struct {
	ID       int
	PrepTime time.Duration
}

// Original sequential processing function
func processOrder(order Order) {
	fmt.Printf("📝 Order %d: Started processing\n", order.ID)
	time.Sleep(order.PrepTime)
	fmt.Printf("✅ Order %d: Ready for pickup! Time taken: %v\n", order.ID, order.PrepTime)
}

// Simple goroutine
func simpleGoroutine() {
	fmt.Printf("\n=== 1. SIMPLE GOROUTINE ===\n\n")

	order := Order{
		ID: 1, PrepTime: 2 * time.Second,
	}

	fmt.Printf("Before starting goroutine\n")

	// Start processing order in a goroutine
	go processOrder(order)

	fmt.Printf("After starting goroutine - main continues immediately!\n")

	// Wait for goroutine to complete
	time.Sleep(3 * time.Second) // or fmt.Scanln() or sync.WaitGroup in real cases
}

// Multiple goroutines processing orders concurrently
func multipleGoroutines() {
	fmt.Printf("\n=== 2. MULTIPLE GOROUTINES (Concurrent Processing) ===\n\n")
	startTime := time.Now()

	orders := []Order{
		{ID: 1, PrepTime: 2 * time.Second},
		{ID: 2, PrepTime: 3 * time.Second},
		{ID: 3, PrepTime: 1 * time.Second},
		{ID: 4, PrepTime: 4 * time.Second},
		{ID: 5, PrepTime: 2 * time.Second},
	}

	// Process all orders concurrently
	for _, order := range orders {
		go processOrder(order)
	}

	// Wait for all to complete (longest is 4 seconds)
	time.Sleep(5 * time.Second) // use sync.WaitGroup in real cases

	fmt.Printf("\n🚀 Concurrent processing time: %v\n", time.Since(startTime))
}

// Goroutines with parameters and proper synchronization
// WaitGroup is like a counter that tracks how many goroutines are still running.
// We need it to wait for all goroutines to finish before the main program exits.
//
// WaitGroup Methods:
// - Add(1): Tell WaitGroup "one more goroutine is starting" (increment counter)
// - Done(): Tell WaitGroup "this goroutine is finished" (decrement counter)
// - Wait(): Make main goroutine wait until counter reaches zero (all done)
func goroutinesWithWaitGroup() {
	fmt.Printf("\n=== 3. GOROUTINES WITH WAITGROUP (Proper Sync) ===\n\n")

	var wg sync.WaitGroup // WaitGroup to synchronize goroutines
	startTime := time.Now()

	orders := []Order{
		{ID: 1, PrepTime: 2 * time.Second},
		{ID: 2, PrepTime: 3 * time.Second},
		{ID: 3, PrepTime: 1 * time.Second},
		{ID: 4, PrepTime: 4 * time.Second},
		{ID: 5, PrepTime: 2 * time.Second},
	}

	// Process orders with proper synchronization
	for _, order := range orders {
		wg.Add(1) // Increment WaitGroup counter
		go func(o Order) {
			defer wg.Done() // Decrement counter when done
			processOrder(o)
		}(order) // Pass order as parameter to avoid closure capture issues
	}

	wg.Wait() // Wait for all goroutines to complete

	fmt.Printf("\n⏱️  Sequential Processing time: 12s\n")
	fmt.Printf("🎯 Concurrent processing time: %v\n", time.Since(startTime)) // time of the longest task
}

// Anonymous goroutines for order processing
func anonymousGoroutines() {
	fmt.Printf("\n=== 4. ANONYMOUS GOROUTINES ===\n\n")

	var wg sync.WaitGroup

	// Anonymous goroutine for rush order
	wg.Add(1)
	go func() {
		defer wg.Done()
		rushOrder := Order{ID: 1, PrepTime: 2 * time.Second}
		fmt.Printf("🔥 Rush Order: Processing immediately!\n")
		processOrder(rushOrder)
	}()

	// Anonymous goroutine with parameters
	customerName := "Alice"
	vipOrder := Order{ID: 2, PrepTime: 1 * time.Second}
	wg.Add(1)
	go func(name string, order Order) {
		defer wg.Done()
		fmt.Printf("👤 VIP Order %d for %s: Started processing\n", order.ID, name)
		time.Sleep(order.PrepTime)
		fmt.Printf("✅ VIP Order %d for %s: Ready for pickup! Time taken: %v (Priority Service)\n",
			order.ID, name, order.PrepTime)
	}(customerName, vipOrder)

	wg.Wait()
}

// Goroutine runtime information during order processing
func goroutineRuntimeInfo() {
	fmt.Println("\n=== 5. GOROUTINE RUNTIME INFO ===")

	fmt.Printf("📊 Initial goroutines count: %d\n", runtime.NumGoroutine())

	var wg sync.WaitGroup // WaitGroup to synchronize goroutines

	orders := []Order{
		{ID: 1, PrepTime: 2 * time.Second},
		{ID: 2, PrepTime: 3 * time.Second},
		{ID: 3, PrepTime: 1 * time.Second},
		{ID: 4, PrepTime: 4 * time.Second},
		{ID: 5, PrepTime: 2 * time.Second},
	}

	// Process orders with proper synchronization
	for _, order := range orders {
		wg.Add(1) // Increment WaitGroup counter
		go func() {
			defer wg.Done() // Decrement counter when done
			processOrder(order)
		}()
	}

	fmt.Printf("📈 After starting order processing, goroutines count: %d\n", runtime.NumGoroutine())

	wg.Wait() // Wait for all goroutines to complete

	fmt.Printf("📉 Final goroutines count: %d\n", runtime.NumGoroutine())
}

// Original sequential processing for comparison
func sequentialProcessing() {
	fmt.Printf("\n=== 0. SEQUENTIAL PROCESSING (Original) ===\n\n")

	startTime := time.Now()

	orders := []Order{
		{ID: 1, PrepTime: 2 * time.Second},
		{ID: 2, PrepTime: 3 * time.Second},
		{ID: 3, PrepTime: 1 * time.Second},
		{ID: 4, PrepTime: 4 * time.Second},
		{ID: 5, PrepTime: 2 * time.Second},
	}

	for _, order := range orders {
		processOrder(order)
	}

	fmt.Printf("\n⏱️  Sequential processing time: %v\n", time.Since(startTime))
}

func main() {
	fmt.Println("==========================================")
	fmt.Println("🏪 Go Concurrency: Order Processing System")
	fmt.Println("==========================================")

	// Show original sequential approach first
	sequentialProcessing()

	// Demonstrate all goroutine concepts
	simpleGoroutine()
	multipleGoroutines()
	goroutinesWithWaitGroup()
	anonymousGoroutines()
	goroutineRuntimeInfo()

	fmt.Println("\n📝 Key Learnings:")
	fmt.Println("✅ Goroutines enable concurrent order processing")
	fmt.Println("✅ Use 'go' keyword to start concurrent processing")
	fmt.Println("✅ WaitGroups provide proper synchronization")
	fmt.Println("✅ Anonymous functions can be used as goroutines")
	fmt.Println("✅ Pass parameters to avoid variable capture issues")
	fmt.Println("✅ Concurrent processing dramatically reduces total time!")
}
