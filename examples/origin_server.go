package main

import (
	"fmt"
	"github.com/amahi/go-metadata"
)

func main() {
	data, err := metadata.GetMetadata("Marx Brothers - At the Circus (1939).avi","movie")
	if err != nil {
		fmt.Println("Error: ",err)
	} else {
        	fmt.Println("Data: \n",data)
	}
}
