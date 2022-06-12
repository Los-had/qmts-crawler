package utils

import (
    "fmt"
    "strings"
    "net/url"
    "github.com/gocolly/colly"
)

type PageData struct {
    URL         string `json:"url"`
    Title       string `json:"title"`
    Description string `json:"description"`
}

type SiteData struct {
    AboutPage    PageData `json:"about"`
    ContactsPage PageData `json:"contacts"`
    FAQPage      PageData `json:"FAQ"`
    DownloadPage PageData `json:"download"`
}

type Result struct {
    Favicon     string    `json:"favicon"`
    URL         string    `json:"url"`
    Title       string    `json:"title"`
    Description string    `json:"description"`
    Keywords    []string  `json:"keywords"`
    SitePages   SiteData  `json:"pages"`
    Images      []Image   `json:"images"`
    Visited     bool      `json:"visited"`
    VisitedTime string    `json:"time"`
}

type Seed struct {
    Host   string `json:"host"`
    Scheme string `json:"scheme"`
    Params string `json:"params"`
    Port   string `json:"port"`
}

type Image struct {
    URL  string `json:"url"`
    Alt  string `json:"alt"`
    Host string `json:"host"`
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
