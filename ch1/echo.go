package ch1

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func Echo1() {
	fmt.Println("Echo 1")
	start := time.Now()
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	secs := time.Since(start).Seconds()
	fmt.Printf("Runtime: %f s\n", secs)
}

func Echo2() {
	fmt.Println("Echo 2")
	start := time.Now()
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	secs := time.Since(start).Seconds()
	fmt.Printf("Runtime: %f s\n", secs)
}

func Echo3() {
	fmt.Println("Echo 3")
	start := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	secs := time.Since(start).Seconds()
	fmt.Printf("Runtime: %f s\n", secs)
}
