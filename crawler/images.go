package crawler

import (
    "net/url"
    "github.com/gocolly/colly"
)

type Image struct {
    URL  string `json:"url"`
    Alt  string `json:"alt"`
    Host string `json:"host"`
}

// Find all the images in a web page
func FindAllImages(seed string) []Image {
    crawler := colly.NewCollector(
        colly.Async(true),
        colly.IgnoreRobotsTxt(),
    )
    var imgs []Image

    crawler.OnHTML("img", func (e *colly.HTMLElement) {
        alt := e.Attr("alt")
        host, err := url.Parse(e.Attr("src"))
        if err != nil {
            host.Host = "Cannot find the image host."
        }

        if alt == "" {
            alt = "Description not provided"
        }
        
        imgs = append(imgs, Image{
            URL: e.Attr("src"),
            Alt: alt,
            Host: host.Host,
        })
    })

    crawler.Wait()
    crawler.Visit(seed)

    return imgs
}