package main

import (
	"fmt"
	"log"

	"github.com/natebrennand/concurrency_and_go/cat_util"
)

func main() {
	// Does parallelizing make it faster?
	// runtime.GOMAXPROCS(runtime.NumCPU())

	catList, err := cat_util.GetList("cats.txt")
	if err != nil {
		log.Fatalf("Could not generate list of catList => %s", err.Error())
	}

	catPipe := make(chan string)
	for i := 0; i < 100; i++ {
		go func(pipe chan string) {
			for {
				cat := <-pipe
				if cat == "" {
					continue
				}
				err := cat_util.GetCat(cat)
				if err != nil {
					log.Printf("Failed to download cat, %s => %s", cat, err.Error())
				}
				fmt.Print(".")
			}
		}(catPipe)
	}

	for _, cat := range catList {
		catPipe <- cat
	}
}
