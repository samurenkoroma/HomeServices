package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()

	ctxWithTimeout, cancel := context.WithTimeout(ctx, 4*time.Second)
	defer cancel()

	done := make(chan struct{})
	go func() {

		time.Sleep(3 * time.Second)
		close(done)
	}()

	select {
	case <-done:
		fmt.Println("Задание выполнено")
	case <-ctxWithTimeout.Done():
		fmt.Println("Время контекста")

	}

}
