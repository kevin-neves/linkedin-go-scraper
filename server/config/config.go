package config

import (
	"fmt"
	"strings"
)

type Configuration struct {
	MainLink      string   `json:"main_link"`
	SearchKey     string   `json:"search_key"`
	Filters       []string `json:"filters"`
	PagesScraped  int      `json:"pages_scraped"`
	ScrapeTimeout int      `json:"scrape_timeout"`
}

const link = "https://www.linkedin.com/jobs/search/?f_TPR=r86400&f_WT=2&geoId=106057199&keywords=%s&location=Brasil&refresh=true&sortBy=DD&start="
const searchKey = "desenvolvedor"

var configuration Configuration

func parseSearchKey(sk string) string {
	sk = strings.ReplaceAll(sk, " ", "%20")
	return sk
}

func InitConfig() {
	configuration.MainLink = fmt.Sprintf(link, searchKey)
	configuration.SearchKey = searchKey
	configuration.Filters = []string{}
	configuration.PagesScraped = 4
	configuration.ScrapeTimeout = 15
}

func SetConfiguration(config Configuration) {
	config.SearchKey = parseSearchKey(config.SearchKey)
	configuration = config
}

func GetConfiguration() Configuration {
	return configuration
}

func GetLink() string {
	return configuration.MainLink
}

func GetFilters() []string {
	return configuration.Filters
}

func GetPagesScraped() int {
	return configuration.PagesScraped
}

func GetScrapeTimeout() int {
	if configuration.ScrapeTimeout < 1 {
		return 1
	}
	return configuration.ScrapeTimeout
}
