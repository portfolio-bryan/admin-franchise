package scrapfranquise

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	urlLib "net/url"

	"github.com/likexian/whois"
	whoisparser "github.com/likexian/whois-parser"
	"golang.org/x/net/html"
)

type ScrapResponse struct {
	HTMLMetaData HTMLMeta
	Protocol     string
	Jumps        int
	WhoisData    whoisparser.WhoisInfo
}

type WhoisData struct{}

type HTMLMeta struct {
	Title       string
	Description string
	Image       string
	SiteName    string
}

type ProtocolAndJumps struct {
	Protocol string
	Jumps    int
}

type Endpoint struct {
	IpAddress string `json:"ipAddress"`
	Grade     string `json:"grade"`
	SeverName string `json:"serverName"`
}

type SSLLabsResponse struct {
	Host      string     `json:"host"`
	Port      int        `json:"port"`
	Protocol  string     `json:"protocol"`
	Endpoints []Endpoint `json:"endpoints"`
}

func ScrapFranquise(ctx context.Context, url string) (ScrapResponse, error) {
	// Get HTML meta data
	htmlMetaData, err := GetHTMLMetaData(url)

	if err != nil {
		return ScrapResponse{}, err
	}

	// Get information of the protocol and jumps
	protocolAndJumps, err := GetProtocolAndJumps(url)
	if err != nil {
		return ScrapResponse{}, err
	}

	// Get whois data
	whoisData, err := GetWhoisData(url)
	if err != nil {
		return ScrapResponse{}, err
	}

	return ScrapResponse{
		HTMLMetaData: htmlMetaData,
		Protocol:     protocolAndJumps.Protocol,
		Jumps:        protocolAndJumps.Jumps,
		WhoisData:    whoisData,
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

func GetProtocolAndJumps(url string) (ProtocolAndJumps, error) {
	ssllabsUrl := fmt.Sprintf("https://api.ssllabs.com/api/v3/analyze?host=%s", url)

	resp, err := http.Get(ssllabsUrl)
	if err != nil {
		return ProtocolAndJumps{}, err
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return ProtocolAndJumps{}, err
	}

	sslLabsResponse := SSLLabsResponse{}

	err = json.Unmarshal(b, &sslLabsResponse)
	if err != nil {
		return ProtocolAndJumps{}, err
	}

	return ProtocolAndJumps{
		Protocol: sslLabsResponse.Protocol,
		Jumps:    len(sslLabsResponse.Endpoints),
	}, nil
}

func GetWhoisData(url string) (whoisparser.WhoisInfo, error) {
	r, err := urlLib.Parse(url)
	if err != nil {
		return whoisparser.WhoisInfo{}, err
	}

	whois_raw, err := whois.Whois(r.Host)
	if err != nil {
		return whoisparser.WhoisInfo{}, err
	}

	result, err := whoisparser.Parse(whois_raw)

	if err != nil {
		return whoisparser.WhoisInfo{}, err
	}

	return result, nil

}
