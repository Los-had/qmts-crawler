package utils

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
    Lang        string    `json:"lang"`
    Title       string    `json:"title"`
    Description string    `json:"description"`
    Keywords    []string  `json:"keywords"`
    SitePages   SiteData  `json:"pages"`
    Images      []Image   `json:"images"`
    Visited     bool      `json:"visited"`
    VisitedTime string    `json:"time"`
    Hash        string    `json:"hash"`
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

type Suggestion struct {
    Text string `json:"text"`
    From string `json:"from"`
}

type StackOverFlowQuestion struct {
    URL     string `json:"url"`
    Title   string `json:"title"`
    Summary string `json:"summary"`
    Votes   string `json:"votes"`
    Date    string `json:"date"`
}
