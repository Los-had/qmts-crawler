package utils

import (
    "fmt"
    "strings"
    "net/url"
    "crypto/sha256"
    "github.com/gocolly/colly"
)


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

// Generate urls
func MakeURLS(url string) []string {
    var links []string
    if strings.HasSuffix(url ,"/") {
        links = append(links, url + "about/")
        links = append(links, url + "contact/")
        links = append(links, url + "faq/")
        links = append(links, url + "download/")
    } else {
        links = append(links, url + "/about/")
        links = append(links, url + "/contact/")
        links = append(links, url + "/faq/")
        links = append(links, url + "/download/")
    }

    return links
}

// Get the about page
func GetAboutPage(site string) PageData {
    var res PageData
    toBeCrawled := MakeURLS(site)
    c := colly.NewCollector(
        colly.Async(true),
        colly.UserAgent(UserAgent),
    )

    c.OnHTML("title", func (e *colly.HTMLElement) {
        res.Title = strings.TrimSpace(e.Text)
    })

    c.OnHTML("meta[name=description]", func (e *colly.HTMLElement) {
        res.Description = strings.TrimSpace(e.Attr("content"))
    })
    
    res.URL = toBeCrawled[0]

    c.Visit(toBeCrawled[0])
    c.Wait()

    return res
}

// Get the contacts page
func GetContactsPage(site string) PageData {
    var res PageData
    toBeCrawled := MakeURLS(site)
    c := colly.NewCollector(
        colly.Async(true),
        colly.UserAgent(UserAgent),
    )

    c.OnHTML("title", func (e *colly.HTMLElement) {
        res.Title = strings.TrimSpace(e.Text)
    })

    c.OnHTML("meta[name=description]", func (e *colly.HTMLElement) {
        res.Description = strings.TrimSpace(e.Attr("content"))
    })
    
    res.URL = toBeCrawled[1]

    c.Visit(toBeCrawled[1])
    c.Wait()

    return res
}

// Get the FAQ page
func GetFAQtPage(site string) PageData {
    var res PageData
    toBeCrawled := MakeURLS(site)
    c := colly.NewCollector(
        colly.Async(true),
        colly.UserAgent(UserAgent),
    )

    c.OnHTML("title", func (e *colly.HTMLElement) {
        res.Title = strings.TrimSpace(e.Text)
    })

    c.OnHTML("meta[name=description]", func (e *colly.HTMLElement) {
        res.Description = strings.TrimSpace(e.Attr("content"))
    })
    
    res.URL = toBeCrawled[2]

    c.Visit(toBeCrawled[2])
    c.Wait()

    return res
}

// Get the download page
func GetDownloadPage(site string) PageData {
    var res PageData
    toBeCrawled := MakeURLS(site)
    c := colly.NewCollector(
        colly.Async(true),
        colly.UserAgent(UserAgent),
    )

    c.OnHTML("title", func (e *colly.HTMLElement) {
        res.Title = strings.TrimSpace(e.Text)
    })

    c.OnHTML("meta[name=description]", func (e *colly.HTMLElement) {
        res.Description = strings.TrimSpace(e.Attr("content"))
    })
    
    res.URL = toBeCrawled[3]

    c.Visit(toBeCrawled[3])
    c.Wait()

    return res
}

// Generate a unique hash by passing a url
func CreateHash(link string) string {
    hash := sha256.Sum256([]byte(link))
	return string(hash[:])
}
