package crawler

import (
    "fmt"
    "time"
    "github.com/Los-had/qmts-crawler/utils"
)

var visited map[string]bool = map[string]bool{} // map with the visited links
var crawled []utils.Result

// TODO: improve this shit

// Executes the crawler
func StartCrawling(seed string) {
    var queue []string
    startUrls := Crawl(seed)
    
    for _, i := range startUrls {
        currentPage := Crawl(i)
        for _, u := range currentPage {
            queue = append(queue, u)
        }
    }

    // copy the startUrls array to the queue]
    /*
    for _, i := range startUrls {
        queue = append(queue, i)
    }
    */

    time.Sleep(time.Second * 30)

    for _, url := range queue {
        if ok := visited[url]; ok {
            continue
        }

        time.Sleep(time.Second * 1)
        
        crawled = append(crawled, Scrape(url))
        visited[url] = true
    }

    fmt.Println("Crawled info:\n", crawled)
    fmt.Println("Already crawled sites:", visited)
}