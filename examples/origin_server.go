package main

import (
	"fmt"
	"github.com/nileshjagnik/metadata"
)

func main() {
	data, err := metadata.GetMetadata("Comedy Central Presents - 5x09 - Tom Papa (1).avi","tv")
	if err != nil {
		fmt.Println(err)
	} else {
        	fmt.Println(data)
	}
}
