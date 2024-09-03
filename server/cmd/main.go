package main

import (
	"fmt"
	"server/config"
)

func main() {
	cfg := config.MustLoad()
	fmt.Printf("%#v\n", cfg)
}
