package main

import (
	"fmt"
	"log"

	"github.com/natebrennand/concurrency_and_go/cat_util"
)

func main() {
	catList, err := cat_util.GetList("cats.txt")
	if err != nil {
		log.Fatalf("Could not generate list of catList => %s", err.Error())
	}

	for _, cat := range catList {
		if cat == "" {
			continue
		}
		err := cat_util.GetCat(cat)
		if err != nil {
			log.Printf("Failed to download cat, %s => %s", cat, err.Error())
		}
		fmt.Print(".")
	}
}
