package main

import (
	"fmt"
	"github.com/amahi/go-metadata"
)

func main() {
	data, err := metadata.GetMetadata("The.Prince.and.the.Pauper.avi","movie")
	if err != nil {
		fmt.Println(err)
	} else {
        	fmt.Println(data)
	}
}
