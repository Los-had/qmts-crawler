package crawler

import "fmt"

var visited map[string]bool = map[string]bool{} // map with the visited links
var crawled []Result

// TODO: improve this shit

// Executes the crawler
func StartCrawling(seed string) {
    var queue []string
    startUrls := Crawl(seed)

    // copy the startUrls array to the queue
    for _, i := range startUrls {
        queue = append(queue, i)
    }

    for _, url := range startUrls {
        if ok := visited[url]; ok {
            continue
        }
        
         crawled = append(crawled, Scrape(url))
        visited[url] = true
        //queue[idx]
    }

    fmt.Println("Crawled info:\n", crawled)
    fmt.Println("Already crawled sites:", visited)
}