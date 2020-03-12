package main

import "github.com/ayaka-tajiri/web-crawler/router"

func main() {
	api := router.Api()

	err := api.Start(":1323")
	if err != nil {
	}
}
