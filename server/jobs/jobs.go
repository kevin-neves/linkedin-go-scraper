package jobs

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/kevin-neves/linkedin-go-scraper/server/config"

	"github.com/kevin-neves/linkedin-go-scraper/server/scraper"
)

var newJobs = make(scraper.JobsMap)

type JobsResponse struct {
	JobTitle scraper.JobTitle `json:"job_title"`
	JobLink  scraper.JobLink  `json:"job_link"`
}

func mapToJsonResponse(m scraper.JobsMap) []JobsResponse {
	r := []JobsResponse{}
	for k, v := range m {
		r = append(r, JobsResponse{
			JobTitle: v,
			JobLink:  k,
		})
	}
	return r
}

func GetNewJobs() []JobsResponse {
	link := config.GetLink()
	var wg sync.WaitGroup
	for i := 0; i < config.GetPagesScraped(); i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			scraper.StartScrapper(fmt.Sprintf("%s%d", link, i*25), &newJobs)
		}(i)
	}

	wg.Wait()

	listJobs := mapToJsonResponse(newJobs)
	wg.Add(1)
	go func() {
		defer wg.Done()
		scraper.DescriptionScraper(string(listJobs[0].JobLink))
	}()

	wg.Wait()

	log.Println(len(newJobs))
	return listJobs
}

func UpdateJobs() {
	var wg sync.WaitGroup
	go func() {
		for {
			wg.Add(1)
			GetNewJobs()
			wg.Done()
			time.Sleep(time.Minute * time.Duration(config.GetScrapeTimeout()))
		}
	}()
	wg.Wait()
}
