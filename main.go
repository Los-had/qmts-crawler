package main

import (
    "flag"
	"fmt"
	"github.com/Los-had/qmts-crawler/crawler"
)

func main() {
    var seed string

    // <-seed=> command line argument
    flag.StringVar(&seed, "seed", "", "Crawler seed list")
    flag.Parse()

    fmt.Println("Crawling", seed)

    crawler.StartCrawling(seed)
}
 