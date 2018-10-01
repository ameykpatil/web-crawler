package crawler

import (
	"fmt"
	"github.com/ameykpatil/web-crawler/utils/html"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

// SiteMap contains url & other site-maps of urls referenced in the current url
type SiteMap struct {
	URL      string     `json:"url"`
	SiteMaps []*SiteMap `json:"siteMap,omitempty"`
}

// NewSiteMap is a constructor for creating SiteMap instance
func NewSiteMap(url string) *SiteMap {
	return &SiteMap{
		URL:      url,
		SiteMaps: make([]*SiteMap, 0),
	}
}

// Crawler is a struct which encapsulate all the properties related to crawling
type Crawler struct {
	seedURL   string            // the base URL of the website being crawled
	host      string            // host of the seed URL
	urls      map[string]string // map which contains urls processed to avoid duplicate processing
	siteMap   *SiteMap          // siteMap to show the url tree
	urlsMutex *sync.Mutex       // mutex to avoid concurrent access to urls map
}

// NewCrawler is a constructor for creating Crawler instance
func NewCrawler(seedURL string) *Crawler {
	parsedURL, err := url.Parse(seedURL)

	if err == nil {
		newSiteMap := NewSiteMap(seedURL)
		return &Crawler{
			seedURL:   seedURL,
			host:      parsedURL.Host,
			urls:      map[string]string{seedURL: "processing"},
			siteMap:   newSiteMap,
			urlsMutex: &sync.Mutex{},
		}
	}

	return &Crawler{}
}

// GetSiteMap returns a SiteMap of the crawler
func (crawler *Crawler) GetSiteMap() *SiteMap {
	return crawler.siteMap
}

// Start function starts crawling of the seedURL
func (crawler *Crawler) Start() {
	var wg sync.WaitGroup
	wg.Add(1)
	crawler.crawl(crawler.seedURL, &crawler.siteMap.SiteMaps, &wg)
	fmt.Printf("completed crawling root URL %s \n", crawler.seedURL)
	wg.Wait()
}

// crawl recursively fetch urls within a page & add it to siteMap
func (crawler *Crawler) crawl(pageURL string, siteMaps *[]*SiteMap, wg *sync.WaitGroup) {
	fmt.Printf("%v\n", pageURL)
	//send http request
	resp, err := http.Get(pageURL)
	if err != nil {
		fmt.Println("An error has occured", pageURL, err)
		wg.Done()
	} else {
		defer resp.Body.Close()
		linkMap := html.GetLinks(resp.Body)

		var childWaitGroup sync.WaitGroup

		for link := range linkMap {
			// handle relative link
			if strings.HasPrefix(link, "/") {
				if err == nil {
					link = crawler.seedURL + link
				}
			}
			// ignore link if it is from external domain
			parsedURL, err := url.Parse(link)
			if err == nil && parsedURL.Host != crawler.host {
				continue
			}
			linkSiteMap := NewSiteMap(link)
			*siteMaps = append(*siteMaps, linkSiteMap)
			crawler.urlsMutex.Lock()
			if crawler.urls[link] != "processing" && crawler.urls[link] != "processed" {
				crawler.urls[link] = "processing"
				childWaitGroup.Add(1)
				go crawler.crawl(link, &linkSiteMap.SiteMaps, &childWaitGroup)
			}
			crawler.urlsMutex.Unlock()
		}

		childWaitGroup.Wait()
		wg.Done()
	}
}
