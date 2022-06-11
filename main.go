package main

import (
    "flag"
	"fmt"
	"github.com/Los-had/qmts-crawler/crawler"
)


func main() {
    //var count int64
    var seed string

    flag.StringVar(&seed, "seed", "", "Crawler seed list")
    flag.Parse()

    fmt.Println("Crawling", seed)

    crawler.StartCrawling(seed)
    /*
    fmt.Println("--- RESULTS ---")
    fmt.Printf("%v results found!\n", count)
    */
}
 