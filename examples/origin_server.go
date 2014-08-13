package main

import (
	"fmt"
	"github.com/amahi/go-metadata"
)

const size int = 1000000

func main() {
	Lib, err := metadata.Init(size, "metadata.db", "TMDB-API-KEY", "TVRAGE-API-KEY", "TVDB-API-KEY")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		data, err := Lib.GetMetadata("Argo", "movie")
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println(data)
		}
	}
}
