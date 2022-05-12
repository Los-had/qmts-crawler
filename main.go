package main

import (
	"fmt"
	"github.com/Los-had/qmts-crawler/crawler"
)


func main() {
    var seed string
    fmt.Println("--- QMTS crawler ---")
    fmt.Scanln(&seed)
    seedinfo := crawler.GetSeedInfo(seed)
    fmt.Println("--- SEED INFO ---")
    fmt.Printf(
        "Host: %v \nScheme: %v \nPort: %v \nParams: %v \n",
        seedinfo.Host,
        seedinfo.Scheme,
        seedinfo.Port,
        seedinfo.Params,
    )
    fmt.Println("--- RESULTS ---")
    r := crawler.Crawl(seed)
    for _, i := range r {
        fmt.Println(crawler.Scrape(i))
    }
}
