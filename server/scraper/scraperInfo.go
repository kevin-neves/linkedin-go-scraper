package scraper

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

const baseLink = "https://br.linkedin.com"

func getDescription(c *colly.Collector) {
	c.OnHTML("div[span]", func(e *colly.HTMLElement) {
		log.Println("SUCCESS!!!!")
		log.Println(e.Attr("class"))
	})
}

func DescriptionScraper(path string) {
	log.Println("enter description")

	c := colly.NewCollector(
		colly.Async(true),
	)

	addCallbacks(c)
	getDescription(c)

	c.Visit(fmt.Sprintf("%s%s", baseLink, path))

	c.Wait()
}
