# Web-Crawler
A simple web crawler in Go.  

## Problem Statement
write a simple web crawler in a programming language of your choice. The crawler should be limited to one domain, it would crawl all pages within a domain, but not follow external links. Given a URL, it should print a simple site map, showing the links between pages.  

## Instructions
Clone this repository.  
Go to the directory where repository is cloned.   
(`$GOPATH/src/github.com/ameykpatil/web-crawler`)

## Run Tests
`/bin/sh ./check.sh`  
`check.sh` is a file created which runs multiple checks such as `fmt`, `vet`, `lint` & `test`  

## Run Service
`go install`   
Check if the service is running by hitting `http://localhost:4000/ping`  
You should get `message` as `pong`  

## Crawl Monzo.com
Hit following url in browser  
`http://localhost:4000/crawl?seedUrl=https://monzo.com`

## Notes
`siteMap` is returned as a `json` having nested structure.  
Crawling of different links take place in a parallel fashion using `go routines`.  
`urls` map is synchronized using `mutex` to avoid concurrent access.  
Crawling of links found within a link is synchronized & handled using `waitGroup`.    
Tests are written for utilities (To be frank, I have never written tests involving concurrency).  
