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
            host.Scheme = "Cannot find the image host."
        }

        if alt == "" {
            alt = "Description not provided"
        }
        
        imgs = append(imgs, Image{
            URL: e.Attr("src"),
            Alt: alt,
            Host: host.Scheme,
        })
    })

    crawler.Wait()
    crawler.Visit(seed)

    return imgs
}