# Go Exercises

Exercices and examples for the book "The Go Programming Language" by Alan A. A. Donovan and Brian W. Kernighan.

## Chapter 1

### [Echo1, Echo2, Echo3](internal/ch1/echo.go)

#### How to run

```bash
go run cmd/main/main.go
# No parameter
```

#### Topics covered

- Basic syntax
- Command line arguments
- Basic I/O

### [Dup1, Dup3](internal/ch1/dup.go)

#### How to run

```bash
# Dup1 - no parameter
go run cmd/main/main.go

# Dup3 - filename
go run cmd/main/main.go assets/dup3_sample.txt
```

#### Topics covered

- Maps
- File I/O
- Command line arguments

### [Lissajous](internal/ch1/lissajous.go)

#### How to run

```bash
go build cmd/lissajous/lissajous.go && ./lissajous > assets/lissajous.gif
# Run from a different model
```

#### Topics covered

- Image library
- Math library

### [Fetch](internal/ch1/fetch.go)

#### How to run

```bash
go run cmd/main/main.go http://gopl.io
# URL
```

#### Topics covered

- HTTP
- Copy

### [Fetchall](internal/ch1/fetchall.go)

#### How to run

```bash
go run cmd/main/main.go golang.org gopl.io godoc.org
# URLs
```

#### Topics covered

- HTTP
- Goroutines
- Channels
