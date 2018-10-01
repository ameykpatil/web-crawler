package html

import (
	"golang.org/x/net/html"
	"io"
	"net/url"
	"strings"
)

// GetLinks parse & return links in the body of the page provided
func GetLinks(body io.Reader) map[string]bool {
	links := make(map[string]bool)
	tokenizer := html.NewTokenizer(body)
	for {
		tokenType := tokenizer.Next()
		switch tokenType {
		case html.ErrorToken:
			return links
		case html.StartTagToken, html.EndTagToken:
			token := tokenizer.Token()
			if "a" == token.Data {
				for _, attr := range token.Attr {
					if attr.Key == "href" {
						link := attr.Val
						parsedURL, err := url.Parse(link)
						if err == nil {
							query := parsedURL.RawQuery
							link = strings.Replace(link, query, "", 1)
							frag := parsedURL.Fragment
							link = strings.Replace(link, frag, "", 1)
							links[link] = true
						}
					}
				}
			}
		}
	}
}
