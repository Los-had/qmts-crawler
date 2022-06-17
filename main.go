package main

import (
    "flag"
	"fmt"
	"github.com/Los-had/qmts-crawler/crawler"
    "github.com/Los-had/qmts-crawler/engines"
)

func main() {
    var seed string
    var suggest string

    // <-seed=> command line argument
    flag.StringVar(&seed, "seed", "", "Crawler seed list")
    // <-autocomplete=> command line argument
    flag.StringVar(&suggest, "autocomplete", "", "Auto complete your query")
    flag.Parse()

    if suggest != "" {
        fmt.Println(engines.AutoComplete(suggest))

        return
    } else if seed != "" {
        fmt.Println("Crawling", seed)
        crawler.StartCrawling(seed)

        return
    } else {
        fmt.Println("Invalid, no action specified.")
    }
}
 