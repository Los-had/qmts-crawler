package engines

import (
    "github.com/gocolly/colly"
    "github.com/Los-had/qmts-crawler/utils"
)

func GetStackOverflowQuetions(query string) []utils.StackOverFlowQuestion {
    var questions []utils.StackOverFlowQuestion
    c := colly.NewCollector(
        colly.IgnoreRobotsTxt(),
        colly.Async(true),
        colly.UserAgent(utils.UserAgent),
    )

    c.OnHTML("div.s-post-summary", func (e *colly.HTMLElement) {
        url := "https://stackoverflow.com" + e.ChildAttr("a.s-link", "href")
        title := e.ChildText("a.s-link")
        summary := e.ChildText("div.s-post-summary--content-excerpt")
        votes := e.ChildAttr("div.s-post-summary--stats-item", "title")[9:]
        date := e.ChildText("span.relativetime")

        questions = append(questions, utils.StackOverFlowQuestion{
            URL: url,
            Title: title,
            Summary: summary,
            Votes: votes,
            Date: date,
        })
    })

    c.Wait()
    c.Visit(utils.StackOverflowSearchURL + query)

    return questions
}
