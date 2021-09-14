package main

import (
	"io/ioutil"
	"net/http"
	"github.com/gocolly/colly"
)

func main() {
	collyRes := colly.NewCollector()
	print(collyRes)
	res, _ := http.Get("https://www.githubstatus.com/")
	defer res.Body.Close()
	content, _ := ioutil.ReadAll(res.Body)
	contentAsString := string(content)
	print(contentAsString)
}
