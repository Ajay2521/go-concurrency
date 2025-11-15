# Goroutines and WaitGroup

## Overview

This Go program demonstrates the fundamentals of goroutines and WaitGroup for concurrent programming. It shows how to transform sequential order processing into efficient concurrent execution with proper synchronization.

## What You'll Learn

- Creating and launching goroutines with the `go` keyword
- Synchronizing goroutines using `sync.WaitGroup`
- Avoiding common concurrency pitfalls
- Measuring performance improvements from concurrency

## Code Structure

### Order Processing System

```go
type Order struct {
    ID       int
    PrepTime time.Duration
}

func processOrder(order Order) {
    fmt.Printf("Processing order %d (prep time: %v)\n", order.ID, order.PrepTime)
    time.Sleep(order.PrepTime)
    fmt.Printf("Order %d completed\n", order.ID)
}
```

## Sequential vs Concurrent Execution

### Sequential Processing (Baseline)

```
Order 1 → Order 2 → Order 3 → Order 4 → Order 5
Total: 12 seconds
```

### Concurrent Processing with Goroutines

```
Order 1 ┐
Order 2 ├─ All run simultaneously
Order 3 ├─ Total: 4 seconds (fastest completion)
Order 4 ├─ 67% time reduction!
Order 5 ┘
```

## Key Concepts

### 1. Basic Goroutine Launch

```go
// Start a goroutine
go processOrder(order)
```

### 2. WaitGroup for Synchronization

```go
var wg sync.WaitGroup

wg.Add(1)           // Increment counter
go func() {
    defer wg.Done()  // Decrement when done
    processOrder(order)
}()
wg.Wait()           // Wait for all to complete
```

### 3. Parameter Passing to Avoid Capture Issues

```go
// CORRECT: Pass parameters explicitly
for _, order := range orders {
    wg.Add(1)
    go func(o Order) {
        defer wg.Done()
        processOrder(o)
    }(order)
}
```

## Common Patterns Demonstrated

### Multiple Goroutines with WaitGroup

```go
func goroutinesWithWaitGroup() {
    var wg sync.WaitGroup
  
    for _, order := range orders {
        wg.Add(1)
        go func(o Order) {
            defer wg.Done()
            processOrder(o)
        }(order)
    }
  
    wg.Wait()
}
```

### Anonymous Goroutines

```go
wg.Add(1)
go func(customerName string, order Order) {
    defer wg.Done()
    fmt.Printf("VIP processing for %s\n", customerName)
    processOrder(order)
}("Premium Customer", vipOrder)
```

## Performance Results

| Method     | Execution Time | Improvement |
| ---------- | -------------- | ----------- |
| Sequential | ~12 seconds    | Baseline    |
| Concurrent | ~4 seconds     | 67% faster  |

## Critical Pitfalls to Avoid

### 1. Variable Capture in Loops

```go
// WRONG - all goroutines use the last loop value
for _, order := range orders {
    go func() {
        processOrder(order) // Bug: uses final order value
    }()
}

// CORRECT - pass parameter explicitly
for _, order := range orders {
    go func(o Order) {
        processOrder(o)     // Each goroutine gets correct order
    }(order)
}
```

### 2. Missing Synchronization

```go
// WRONG - program exits before goroutines complete
go processOrder(order)
return

// CORRECT - wait for completion
var wg sync.WaitGroup
wg.Add(1)
go func() {
    defer wg.Done()
    processOrder(order)
}()
wg.Wait()
```

## Runtime Monitoring

Track goroutine lifecycle:

```go
fmt.Printf("Goroutines before: %d\n", runtime.NumGoroutine())
// Launch goroutines
fmt.Printf("Goroutines during: %d\n", runtime.NumGoroutine())
wg.Wait()
fmt.Printf("Goroutines after: %d\n", runtime.NumGoroutine())
```

## Best Practices

### ✅ Do

- Use `sync.WaitGroup` for coordination
- Pass parameters to goroutines explicitly
- Use `defer wg.Done()` for cleanup
- Add goroutines to WaitGroup before launching

### ❌ Don't

- Rely on `time.Sleep()` for synchronization
- Capture loop variables directly in goroutines
- Forget to call `wg.Wait()`
- Create goroutines without proper coordination

## When to Use This Pattern

**Ideal for:**

- Independent, parallel tasks
- I/O-bound operations
- Processing collections of data
- Background task execution

**Consider alternatives for:**

- Tasks requiring complex coordination
- Heavy CPU-bound work (limited by core count)
- Simple sequential operations

## Next Steps

After mastering goroutines and WaitGroup, explore:

- **Channels** for goroutine communication
- **Worker pools** for controlled concurrency
- **Context** for cancellation and timeouts
- **Select statements** for channel multiplexing