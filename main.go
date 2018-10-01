package main

import (
	"github.com/ameykpatil/web-crawler/crawler"
	"github.com/ameykpatil/web-crawler/utils/server"
	"net/http"
)

func main() {

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		server.SendResponse(w, map[string]string{"message": "pong"}, http.StatusOK)
	})

	http.HandleFunc("/crawl", CrawlHandler)

	err := http.ListenAndServe(":4000", nil)
	if err != nil {
		panic(err)
	}
}

// CrawlHandler is a handler function for crawl API
func CrawlHandler(w http.ResponseWriter, r *http.Request) {
	seedURL := r.URL.Query().Get("seedUrl")
	if seedURL != "" {
		crawlerInstance := crawler.NewCrawler(seedURL)
		crawlerInstance.Start()
		server.SendResponse(w, crawlerInstance.GetSiteMap(), http.StatusOK)
	} else {
		server.SendResponse(w, nil, http.StatusBadRequest)
	}
}
