package main

import (
	"fmt"
	"github.com/Los-had/qmts-crawler/crawler"
)


func main() {
    var count int64
    var seed string
    //var queue []crawler.Result
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
        count++
    }
    fmt.Printf("%v results found!\n", count)
}
 