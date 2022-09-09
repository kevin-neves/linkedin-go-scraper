package scraper

import (
	"log"
	"math/rand"
	"net/url"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type JobTitle string
type JobLink string
type JobsMap map[JobLink]JobTitle

func randomString() string {
	b := make([]byte, rand.Intn(10)+10)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func addCallbacks(c *colly.Collector) {
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", randomString())
		log.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	c.OnResponse(func(r *colly.Response) {
		log.Println("Visited", r.Request.URL)
	})
}

func getTitles(c *colly.Collector, jobsMap JobsMap) {
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		if strings.Contains(e.Attr("href"), "/jobs/view/") && !strings.Contains(e.Text, "Be an early applicant") {
			u, err := url.Parse(e.Attr("href"))
			if err != nil {
				log.Println(err)
			}
			jobLink := strings.Split(u.EscapedPath(), "?")[0]
			jobsMap[JobLink(jobLink)] = JobTitle(strings.Trim(e.Text, " \n"))
			// titles = append(titles, strings.Trim(e.Text, " \n"))
		}
	})
}

func StartScrapper(link string, jobsMap *JobsMap) {
	c := colly.NewCollector(
		colly.Async(true),
		// colly.Debugger(&debug.LogDebugger{}),
	)
	c.Limit(&colly.LimitRule{
		DomainGlob:  "*linkedin.*",
		Parallelism: 1,
		Delay:       3 * time.Second,
	})

	addCallbacks(c)
	getTitles(c, *jobsMap)

	c.Visit(link)

	c.Wait()
}
