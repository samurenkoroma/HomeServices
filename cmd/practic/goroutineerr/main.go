package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func ping(url string, respCh chan int, errCh chan error) {
	resp, err := http.Get(url)
	if err != nil {
		errCh <- err
		return
	}
	respCh <- resp.StatusCode

}

func main() {
	path := flag.String("file", "url.txt", "Файл со списком url для проверки")
	flag.Parse()

	file, err := os.ReadFile(*path)
	if err != nil {
		panic(err.Error())
	}

	urlSlice := strings.Split(string(file), "\n")

	pingCh := make(chan int)
	errCh := make(chan error)

	for _, url := range urlSlice {
		go ping(url, pingCh, errCh)
	}

	for range urlSlice {

		select {
		case err := <-errCh:
			fmt.Println(err)
		case res := <-pingCh:
			fmt.Println(res)
		}

	}

}
