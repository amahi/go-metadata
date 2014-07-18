package main

import (
	"fmt"
	"github.com/amahi/go-metadata"
)

const size int = 1000000

func main() {
	Lib, err := metadata.Init(size, "metadata.db")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		data, err := Lib.GetMetadata("Breaking Bad", "tv")
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println(data)
		}
	}
}
