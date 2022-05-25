package crawler

import (
	"fmt"
	"net/url"
	"strings"
	"time"
	"github.com/gocolly/colly"
)

var results []Result
var userAgent = "QMTSbot/0.1.1"

type Seed struct {
    Host   string `json:"host"`
    Scheme string `json:"scheme"`
    Params string `json:"params"`
    Port   string `json:"port"`
}

type Result struct {
    Favicon     string    `json:"favicon"`
    URL         string    `json:"url"`
    Title       string    `json:"title"`
    Description string    `json:"description"`
    Keywords    []string  `json:"keywords"`
    Visited     bool      `json:"visited"`
    VisitedTime string    `json:"time"`
}

// Check if is a valid URL
func CheckURL(rawURL string) bool {
    parsedURL, err := url.Parse(rawURL) 
    if err != nil {
        return false
    }
    
    if parsedURL.Scheme != "http" && parsedURL.Scheme != "https"  {
        return false
    }

    return true
}

// Get information of an especific domain
func GetSeedInfo(seed string) Seed {
    si, err := url.Parse(seed)
    if err != nil {
        panic(err)
    }
    return Seed{
        Host: si.Hostname(),
        Scheme: si.Scheme,
        Port: si.Port(),
        Params: si.RawQuery,
    }
}

// Parse the Result struct
func ParseResult(r Result) Result {
    if r.Favicon == "" {
        r.Favicon = ""
    } else if r.Keywords == nil {
        r.Keywords = []string{r.URL, r.Title}
    } else if r.Description == "" {
        r.Description = fmt.Sprintf("Description not provided, %v", r.VisitedTime)
    } else if r.Title == "" {
        r.Title = r.URL
    }

    return r
}

// Get data(favicon, title, description and etc) from a website
func Scrape(url string) Result {
    if !CheckURL(url) {
        fmt.Println("Invalid URL")
        return Result{}
    }
    var result Result
    result.URL = url
    c := colly.NewCollector(
        colly.IgnoreRobotsTxt(),
        colly.UserAgent(userAgent),
        colly.Async(true),
    )
    
    c.OnHTML("title", func (e *colly.HTMLElement) {
        result.Title = strings.TrimSpace(e.Text)
    })

    c.OnHTML("meta[name=keywords]", func (e *colly.HTMLElement) {
        result.Keywords = strings.Split(e.Attr("content"), ", ")
    })
    
    c.OnHTML("meta[name=description]", func (e *colly.HTMLElement) {
        result.Description = strings.TrimSpace(e.Attr("content"))
    })

    c.OnHTML("link[rel=\"shortcut icon\"]", func (e *colly.HTMLElement) {
        favicon := e.Attr("href")
        if !strings.HasPrefix(favicon, "/") {
            result.Favicon = favicon
        } else {
            favicon = url + e.Attr("href")
            result.Favicon = favicon
        }
    })

    c.OnScraped(func (r *colly.Response) {
        result.Visited = true
        result.VisitedTime = time.Now().String()
    })

    c.OnRequest(func (r *colly.Request) {
        fmt.Println("[GET] ->", r.URL)
    })

    c.OnResponse(func (resp *colly.Response) {
        if resp.StatusCode != 200 {
            fmt.Println("Request error, status code:", resp.StatusCode)
            return
        }
        contentType := resp.Headers.Get("Content-Type")
        if !strings.Contains(contentType, "text/html") {
            fmt.Println("Invalid content-type, Content-Type:", contentType)
            return
        }
    })

    c.OnError(func (r *colly.Response, err error) {
        fmt.Println("Request failed:", url, "\nError:", err)
        return
    })
    
    c.Visit(url)
    c.Wait()

    return ParseResult(result)
}

// Get all the links in the webpage
func Crawl(seedlist string) []string {
    var urls []string
    crawler := colly.NewCollector(
        colly.UserAgent(userAgent),
        colly.Async(true),
    )
    
    crawler.OnHTML("a", func (e *colly.HTMLElement) {
        C_URL := e.Attr("href")
        _, err := url.ParseRequestURI(C_URL)
        if err == nil {
            if !strings.HasPrefix(C_URL, "/") {
                urls = append(urls, C_URL)
            } else if strings.HasPrefix(C_URL, "//") {
                C_URL = "http:" + C_URL
                urls = append(urls, C_URL)
            } else {
                C_URL = e.Request.AbsoluteURL(e.Attr("href"))
                if C_URL != "" {
                    urls = append(urls, C_URL)
                }
            }
        }
    })

    crawler.Visit(seedlist)
    crawler.Wait()

    return urls
}
