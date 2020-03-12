package router

import (
	"github.com/ayaka-tajiri/web-crawler/crawler"
	"github.com/labstack/echo"
	"net/http"
)

func Api() *echo.Echo {
	e := echo.New()
	// Routes
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/web_crawler", crawler.StartCrawler())
	return e
}