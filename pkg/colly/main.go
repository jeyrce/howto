package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	collector := colly.NewCollector(
		colly.UserAgent("jeyrce/colly"),
	)
	collector.OnRequest(func(r *colly.Request) {
		fmt.Println(r.URL.String())
	})
	collector.OnHTML("a[href]", func(e *colly.HTMLElement) {
		fmt.Println(e.Text, e.Attr("href"))
	})
	err := collector.Visit("https://ioseek.cn/")
	if err != nil {
		fmt.Errorf("x: %v", err)
		return
	}
}
