package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"githubstatus/service"
	"strings"
)

type GithubStatusInfo struct {
	SystemName        string
	SystemSlugname    string
	SystemStatusLabel string
}

func main() {
	collyCollector := colly.NewCollector()
	var githubStatus []GithubStatusInfo
	collyCollector.OnHTML("div[data-component-id]", func(element *colly.HTMLElement) {
		classAttribute := element.Attr("class")
		systemNameFromHtml := element.DOM.Find("span.name").Text()
		if len(systemNameFromHtml) > 0 && !strings.Contains(systemNameFromHtml, "Visit") {
			slugnameFromClassAttribute := service.GetSystemSlugnameByClassAttribute(classAttribute)
			newGithubStatus := GithubStatusInfo{
				SystemName:        strings.TrimSpace(systemNameFromHtml),
				SystemSlugname:    slugnameFromClassAttribute,
				SystemStatusLabel: service.GetSystemLabelBySlugname(slugnameFromClassAttribute),
			}
			githubStatus = append(githubStatus, newGithubStatus)
		}
	})
	collyCollector.Visit("https://www.githubstatus.com")
	fmt.Println(githubStatus)
}
