package main

import (
	"fmt"
	"os"

	"github.com/oze4/golas"
	"github.com/pkg/profile"
)

func main() {
	defer profile.Start(profile.MemProfileHeap()).Stop()
	reader, err := os.Open("samples/unwrapped.las")
	if err != nil {
		panic(err)
	}
	las := golas.Parse(reader)
	fmt.Println(las.Version())
}
