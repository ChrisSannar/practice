package main

import (
	"context"
	"fmt"
	"time"
)

/**
* NOTE:
* - Context's are good for long running functions
* - Always create contexts first thing in function
* - Ensure you always cancel (use `defer`)
* - Make sure you `<-ctx.Done()` in concurrent function
* - Use contexts for deadline, cancel, and request-scoped data
* - Try not to nest
* - Minimize creation (resource overhead)
* - Avoid passing unnecessary data
 */
func main() {
	fmt.Println("Context")
	// basics()
	// withTimeout()
	withValue()
}

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Processing stopped")
			return
		default:
			fmt.Println("Proessing data...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func basics() {
	ctx, cancel := context.WithCancel(context.Background())

	go worker(ctx)
	time.Sleep(2 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)
}

func slowOperation(ctx context.Context) {
	select {
	case <-time.After(5 * time.Second):
		fmt.Println("Data processing completed")
	case <-ctx.Done():
		fmt.Println("Processing timed out:", ctx.Err())
	}
}

func withTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	slowOperation(ctx)
}

type contextKey string

const taskIDKey contextKey = "taskID"

func taskWorker(ctx context.Context) {
	if taskID, ok := ctx.Value(taskIDKey).(int); ok {
		fmt.Printf("Processing task: %d\n", taskID)
	} else {
		fmt.Println("No task ID found in context")
	}
}

func withValue() {
	ctx := context.WithValue(context.Background(), taskIDKey, 42)
	taskWorker(ctx)
}
