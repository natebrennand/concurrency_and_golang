package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/natebrennand/concurrency_and_go/cat_util"
)

func main() {
	catList, err := cat_util.GetList("cats.txt")
	if err != nil {
		log.Fatalf("Could not generate list of catList => %s", err.Error())
	}

	wg := &sync.WaitGroup{}
	for _, cat := range catList {
		wg.Add(1)
		go func(cat string, wg *sync.WaitGroup) {
			defer wg.Done()
			if cat == "" {
				return
			}

			err := cat_util.GetCat(cat)
			if err != nil {
				log.Printf("Failed to download cat, %s => %s", cat, err.Error())
			}
			fmt.Print(".")
		}(cat, wg)
	}
	wg.Wait()
}
