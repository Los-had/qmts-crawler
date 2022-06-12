package crawler

import (
    "net/url"
    "github.com/gocolly/colly"
    "github.com/Los-had/qmts-crawler/utils"
)

// Find all the images in a web page
func FindAllImages(seed string) []utils.Image {
    crawler := colly.NewCollector(
        colly.Async(true),
        colly.IgnoreRobotsTxt(),
    )
    var imgs []utils.Image

    crawler.OnHTML("img", func (e *colly.HTMLElement) {
        alt := e.Attr("alt")
        host, err := url.Parse(e.Attr("src"))
        if err != nil {
            host.Host = "Cannot find the image host."
        }

        if alt == "" {
            alt = "Description not provided"
        }
        
        imgs = append(imgs, utils.Image{
            URL: e.Attr("src"),
            Alt: alt,
            Host: host.Host,
        })
    })

    crawler.Wait()
    crawler.Visit(seed)

    return imgs
}