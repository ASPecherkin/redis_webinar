package main

import (
	"context"
	"fmt"
)

type Obj struct {
	Name string
}

func main() {
	basic := context.Background()
	example := Obj{Name: "name"}
	ctx := context.WithValue(basic, example, "value")
	example.Name = "not name"
	fmt.Printf("%+v", ctx.Value(example))
}
