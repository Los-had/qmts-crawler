package engines

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"github.com/Los-had/qmts-crawler/utils"
	"github.com/gocolly/colly"
)

func AutoComplete(query string) []string{
    var suggestions []string
    for _, v := range GetGoogleSuggestions(query) {
        suggestions = append(suggestions, v)
    }

    return suggestions
}

// Get brave search query suggestions
func GetBraveSuggestions(query string) []string {
    resp, err := http.Get(utils.BraveApiURL + query)
    if err != nil {
        return []string{}
    }

    data, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return []string{}
    }

    var suggestionsRawArray [][]string

    json.Unmarshal([]byte(data), &suggestionsRawArray)

    return suggestionsRawArray[0]
}

// Get google search suggestions
func GetGoogleSuggestions(query string) []string {
    c := colly.NewCollector(
        colly.Async(true),
        colly.IgnoreRobotsTxt(),
        colly.UserAgent(utils.UserAgent),
    )
    var suggestionsList []string

    c.OnHTML("suggestion", func (e *colly.HTMLElement) {
        suggestionsList = append(
            suggestionsList,
            e.Attr("data"),
        )
    })

    c.Wait()
    c.Visit(utils.GoogleApiURL + query)

    return suggestionsList
}

// Get wikipedia articles suggestions
func GetWikipediaSuggestions(query string) []string {
    resp, err := http.Get(utils.WikipediaApiURL + query)
    if err != nil {
        return []string{}
    }

    data, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return []string{}
    }

    var suggestionsRawArray [][]string

    json.Unmarshal([]byte(data), &suggestionsRawArray)

    return suggestionsRawArray[0]
}