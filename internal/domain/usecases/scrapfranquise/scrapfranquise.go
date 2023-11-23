package scrapfranquise

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/likexian/whois"
	"golang.org/x/net/html"
)

type ScrapResponse struct {
	HTMLMetaData HTMLMeta
}

type HTMLMeta struct {
	Title       string
	Description string
	Image       string
	SiteName    string
}

func ScrapFranquise(ctx context.Context, url string) (ScrapResponse, error) {
	// Get HTML meta data
	htmlMetaData, err := GetHTMLMetaData(url)

	if err != nil {
		return ScrapResponse{}, err
	}

	// Get information of the protocol and jumps
	GetProtocolAndJumps()

	// Get whois data
	GetWhoisData(url)

	return ScrapResponse{
		HTMLMetaData: htmlMetaData,
	}, nil
}

func GetHTMLMetaData(url string) (HTMLMeta, error) {
	resp, err := http.Get(url)
	if err != nil {
		return HTMLMeta{}, err
	}

	defer resp.Body.Close()

	meta := extract(resp.Body)

	return *meta, nil
}

func extract(resp io.Reader) *HTMLMeta {
	z := html.NewTokenizer(resp)

	titleFound := false

	hm := new(HTMLMeta)

	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			return hm
		case html.StartTagToken, html.SelfClosingTagToken:
			t := z.Token()
			if t.Data == `body` {
				return hm
			}
			if t.Data == "title" {
				titleFound = true
			}
			if t.Data == "meta" {
				desc, ok := extractMetaProperty(t, "description")
				if ok {
					hm.Description = desc
				}

				ogTitle, ok := extractMetaProperty(t, "og:title")
				if ok {
					hm.Title = ogTitle
				}

				ogDesc, ok := extractMetaProperty(t, "og:description")
				if ok {
					hm.Description = ogDesc
				}

				ogImage, ok := extractMetaProperty(t, "og:image")
				if ok {
					hm.Image = ogImage
				}

				ogSiteName, ok := extractMetaProperty(t, "og:site_name")
				if ok {
					hm.SiteName = ogSiteName
				}
			}
		case html.TextToken:
			if titleFound {
				t := z.Token()
				hm.Title = t.Data
				titleFound = false
			}
		}
	}
}

func extractMetaProperty(t html.Token, prop string) (content string, ok bool) {
	for _, attr := range t.Attr {
		if attr.Key == "property" && attr.Val == prop {
			ok = true
		}

		if attr.Key == "content" {
			content = attr.Val
		}
	}

	return
}

func GetProtocolAndJumps() {

}

func GetWhoisData(url string) {
	result, err := whois.Whois(url)
	if err == nil {
		fmt.Println("result", result)
	}
}
