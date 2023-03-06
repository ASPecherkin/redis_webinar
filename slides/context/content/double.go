package main

import (
	"context"
	"fmt"
)

func main() {
	basic := context.Background()
	first := context.WithValue(basic, "userID", "1")
	second := context.WithValue(first, "userID", "2")
	fmt.Printf("%+v", second.Value("userID"))
}
