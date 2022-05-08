package main

import (
	"fmt"
	"os"
)

func main() {

	for _, s := range os.Environ() {
		fmt.Printf("child2:%s\n", s)
	}
}
