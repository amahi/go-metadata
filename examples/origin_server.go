package main

import (
	"fmt"
	"github.com/amahi/go-metadata"
)

func main() {
	Lib,err := metadata.Init(1000000,"metdata.db")
	if err!=nil {
	        fmt.Println("Error:",err)
	} else {
	        data, err := Lib.GetMetadata("MythBusters - 7x14 - Dirty vs. Clean Car","tv")
	        if err != nil {
		        fmt.Println("Error:",err)
	        } else {
                	fmt.Println(data)
	        }
        }
}
