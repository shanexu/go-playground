package main

import (
	"context"
	"fmt"
)

func hello(ctx context.Context) {
	fmt.Println("hello world")
}

func main() {
	hello(context.Background())
}

