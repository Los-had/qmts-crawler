package main

import (
	"fmt"
	"net/url"
	"strings"
	"github.com/gocolly/colly"
	//"github.com/gocolly/colly/proxy"
)

var proxyList []string = []string{"http://192.155.107.214:1080", "http://213.230.97.10:3128", "http://170.239.255.2:55443"}
var results []Result
var seed string = "https://youtube.com/"
var userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.54 Safari/537.36"

type Seed struct {
    Host   string `json:"host"`
    Scheme string `json:"scheme"`
    Params string `json:"params"`
    Port   string `json:"port"`
}

type Result struct {
    Favicon     string   `json:"favicon"`
    URL         string   `json:"url"`
    Title       string   `json:"title"`
    Description string   `json:"description"`
    Keywords    []string `json:"keywords"`
    Info        *Seed    `json:"info"`
}

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

func Scrape(url string) Result {
    var result Result
    result.URL = url
    c := colly.NewCollector(
        colly.IgnoreRobotsTxt(),
        colly.UserAgent(userAgent),
        colly.Async(true),
    )

    /*
    if py, err := proxy.RoundRobinProxySwitcher(proxyList...); err != nil {
        fmt.Println("Error occurred:", err)
    } else {
        c.SetProxyFunc(py)
    }
    */
    
    c.OnHTML("title", func (e *colly.HTMLElement) {
        result.Title = e.Text
    })

    c.OnHTML("meta[name=keywords]", func (e *colly.HTMLElement) {
        result.Keywords = strings.Split(e.Attr("content"), ", ")
    })
    
    c.OnHTML("meta[name=description]", func (e *colly.HTMLElement) {
        result.Description = e.Attr("content")
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

    c.OnRequest(func (r *colly.Request) {
        fmt.Println("[GET] ->", r.URL)
    })

    c.OnResponse(func (resp *colly.Response) {
        if resp.StatusCode != 200  || resp.Headers.Get("Content-Type") != "text/html" {
            fmt.Println("Request error, status code:", resp.StatusCode)
            return
        }
    })

    c.OnError(func (r *colly.Response, err error) {
        fmt.Println("Request failed:", url, "\nError:", err)
        return
    })
    
    c.Visit(url)
    c.Wait()

    return result
}

func Crawl(seedlist string) []string {
    var urls []string
    crawler := colly.NewCollector(
        colly.UserAgent(userAgent),
        colly.Async(true),
    )
    crawler.OnHTML("a", func (e *colly.HTMLElement) {
        C_URL := e.Attr("href")
        if !strings.HasPrefix(C_URL, "/") {
            _, err := url.ParseRequestURI(C_URL)
            if err == nil {
                urls = append(urls, C_URL)
            }   
        } else {
            C_URL = e.Request.AbsoluteURL(e.Attr("href"))
            if C_URL != "" {
                urls = append(urls, C_URL)
            }
        }
    })

    crawler.Visit(seedlist)
    crawler.Wait()

    return urls
}

func main() {
    //seedinfo := GetSeedInfo(seed)
    fmt.Println("--- QMTS crawler ---")
    r := Crawl(seed)
    for _, i := range r {
        fmt.Println(Scrape(i))
    }
}
