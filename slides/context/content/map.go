package main

import (
	"context"
	"fmt"
)

func main() {
	basic := context.Background()
	example := map[string]string{"key": "value"}
	ctx := context.WithValue(basic, "meta", example)
	example["key"] = "not value"
	fmt.Printf("%+v", ctx.Value("meta"))
}
