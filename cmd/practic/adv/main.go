package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func sumPart(arr []int, ch chan int) {
	res := 0
	for _, v := range arr {
		res += v
	}
	ch <- res
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	numGoroutines := 5

	ch := make(chan int, numGoroutines)

	partSize := len(arr) / numGoroutines

	for i := 0; i < numGoroutines; i++ {
		start := i * partSize
		end := start + partSize
		go sumPart(arr[start:end], ch)
	}

	totalSum := 0

	for i := 0; i < numGoroutines; i++ {
		totalSum += <-ch
	}

	fmt.Printf("Сумма %d\n", totalSum)

	t := time.Now()
	code := make(chan int)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			requestCh(code)
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(code)
	}()
	for res := range code {
		fmt.Printf("Code: %d\n", res)
	}
	fmt.Println(time.Since(t))
}

func requestCh(codeCh chan int) {
	resp, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println(err)
	}
	codeCh <- resp.StatusCode
}

// func main() {
// 	t := time.Now()

// 	var wg sync.WaitGroup

// 	for i := 0; i < 10; i++ {
// 		wg.Add(1)
// 		// Вариант 1
// 		go requestWg(&wg)

// 		// Вариант 2
// 		go func() {
// 			request()
// 			wg.Done()
// 		}()

// 	}

// 	wg.Wait()
// 	fmt.Println(time.Since(t))
// }

func requestWg(wg *sync.WaitGroup) {
	defer wg.Done()
	request()
}
func request() {
	resp, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Code: %d\n", resp.StatusCode)
}
