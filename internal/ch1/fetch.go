package ch1

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func Fecth() {
	for _, url := range os.Args[1:] {
		url = treatUrl(url)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		var buf bytes.Buffer
		bytes, err := io.Copy(&buf, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}

		fmt.Printf("Content copied:\n%s\n", buf.String())
		fmt.Printf("Status code: %s\n", resp.Status)
		fmt.Printf("Bytes copied: %d\n", bytes)
	}
}

func treatUrl(url string) string {
	if strings.HasPrefix(url, "http://") {
		return url
	}
	return "http://" + url
}
