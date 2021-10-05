package main

import (
	"github.com/gocolly/colly"
	"strings"
)

type Githubstats struct {
	SystemName   string
	Systemstatus bool
}

func main() {
	collyCollector := colly.NewCollector(colly.Async(false))
	var githubStatus []Githubstats
	collyCollector.OnHTML("div[data-component-id]", func(element *colly.HTMLElement) {
		element.ForEach("span", func(_ int, element *colly.HTMLElement) {
			systemNameFromHtml := element.DOM.Filter(".name").Text()
			systemStatusFromHtml := element.DOM.Filter(".component-status").Text()
			if len(systemNameFromHtml) > 0 && !strings.Contains(systemNameFromHtml, "Visit") {
				newGithubStatus := Githubstats{
					SystemName:   strings.TrimSpace(systemNameFromHtml),
					Systemstatus: strings.Contains(systemStatusFromHtml, "Operational"),
				}
				githubStatus = append(githubStatus, newGithubStatus)
			}
		})
	})
	collyCollector.Visit("https://www.githubstatus.com")
}
