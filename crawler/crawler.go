package crawler

import (
	"log"
	"net/url"
    "github.com/Los-had/qmts-crawler/utils"
	"strings"
	"time"
	"github.com/gocolly/colly"
)

var results []utils.Result

// Get data(favicon, title, description and etc) from a website
func Scrape(link string) utils.Result {
    if !utils.CheckURL(link) {
        log.Println("Invalid URL")
        return utils.Result{}
    }
    
    var result utils.Result
    result.URL = link
    result.Hash = utils.CreateHash(link)

    c := colly.NewCollector(
        colly.IgnoreRobotsTxt(),
        colly.UserAgent(utils.UserAgent),
        colly.Async(true),
    )
    
    c.OnHTML("title", func (e *colly.HTMLElement) {
        result.Title = strings.TrimSpace(e.Text)
    })

    c.OnHTML("html", func (e *colly.HTMLElement) {
        result.Lang = e.Attr("lang")
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
            favicon = link + e.Attr("href")
            result.Favicon = favicon
        }
    })

    c.OnScraped(func (r *colly.Response) {
        result.Visited = true
        result.VisitedTime = time.Now().String()
        host, err := url.ParseRequestURI(link)
        if err == nil {
            if host.Scheme + "://" + host.Host + "/" == link {
                result.SitePages.AboutPage = utils.GetAboutPage(link)
                result.SitePages.FAQPage = utils.GetFAQtPage(link)
                result.SitePages.DownloadPage = utils.GetDownloadPage(link)
                result.SitePages.ContactsPage = utils.GetContactsPage(link)
            }

            return
        } else {
            return
        }
    })

    c.OnRequest(func (r *colly.Request) {
        log.Println("[GET] ->", r.URL)
    })

    c.OnResponse(func (resp *colly.Response) {
        if resp.StatusCode != 200 {
            log.Println("Request error, status code:", resp.StatusCode)
            return
        }
        contentType := resp.Headers.Get("Content-Type")
        if !strings.Contains(contentType, "text/html") {
            log.Println("Invalid content-type, Content-Type:", contentType)
            return
        }
    })

    c.OnError(func (r *colly.Response, err error) {
        log.Println("Request failed:", link, "\nError:", err)
        return
    })

    result.Images = FindAllImages(link)
    
    c.Visit(link)
    c.Wait()

    return utils.ParseResult(result)
}

// Get all the links in the webpage
func Crawl(seedlist string) []string {
    var urls []string
    crawler := colly.NewCollector(
        colly.UserAgent(utils.UserAgent),           
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
