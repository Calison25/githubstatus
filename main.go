package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"strings"
)

type Githubstats struct {
	SystemName   string
	Systemstatus bool
}

func main() {
	collyCollector := colly.NewCollector()
	var githubStatus []Githubstats
	collyCollector.OnHTML("div[data-component-id]", func(element *colly.HTMLElement) {
		systemNameFromHtml := element.DOM.Find("span.name").Text()
		systemStatusFromHtml := element.DOM.Find("span.component-status").Text()
		if len(systemNameFromHtml) > 0 && !strings.Contains(systemNameFromHtml, "Visit") {
			newGithubStatus := Githubstats{
				SystemName:   strings.TrimSpace(systemNameFromHtml),
				Systemstatus: strings.Contains(systemStatusFromHtml, "Operational"),
			}
			githubStatus = append(githubStatus, newGithubStatus)
		}
	})
	collyCollector.Visit("https://www.githubstatus.com")
	fmt.Println(githubStatus)
}
