package main

import (
	"fmt"
	"os"
)

func main() {

	for _, s := range os.Environ() {
		fmt.Printf("child1_1:%s\n", s)
	}
}
