package crawler

import (
	"encoding/json"
	"github.com/gocolly/colly/v2"
	"github.com/labstack/echo"
	"net/http"
)

func StartCrawler() func(c echo.Context) error {
	return func(c echo.Context) error {
		url := &url{
			Url: []string{
				"http://www.meaning666.com/schedule/",
				"https://shadowsjapan.com/live/",
			},
		}
		result := url.crawl()

		json.Marshal(result)
		return c.JSON(http.StatusOK, result)
	}
}

type url struct {
	Url []string
}

type crawlerContent struct {
	Url       string   `json:"url"`
	Schedules []string `json:"schedules"`
}

func (i *url) crawl() []*crawlerContent {

	var result []*crawlerContent
	for _, value := range i.Url {
		// meaningのライブ情報を取得する
		c := colly.NewCollector()

		var schedules []string
		c.OnHTML("article", func(e *colly.HTMLElement) {
			schedules = append(schedules, e.Text)
		})

		c.Visit(value)

		result = append(result, &crawlerContent{
			Url: value,
			Schedules: schedules,
		})
	}

	return result
}

