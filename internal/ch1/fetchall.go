package ch1

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func Fetchall() {
	start := time.Now()
	ch := make(chan string)

	for _, url := range os.Args[1:] {
		go fecth(url, ch) // Start a routine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel
	}
	secs := time.Since(start).Seconds()
	fmt.Printf("%.2fs elapsed\n", secs)
}

func fecth(url string, ch chan<- string) {
	start := time.Now()

	url = treatUrl(url)
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nBytes, err := io.Copy(io.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v\n", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nBytes, url)
}
