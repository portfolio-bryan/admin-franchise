package scrapfranquise

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	urlLib "net/url"

	"github.com/bperezgo/admin_franchise/shared/platform/concurrency"
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
	done := make(chan interface{})

	defer close(done)
	// Get HTML meta data
	channGetHTMLMetaData := GetHTMLMetaData(url)
	// Get information of the protocol and jumps
	channProtocolAndJumps := GetProtocolAndJumps(url)
	// Get whois data
	chanWhois := GetWhoisData(url)

	multiplexedStream := concurrency.Fanin(done, channGetHTMLMetaData, channProtocolAndJumps, chanWhois)

	var htmlMetaData HTMLMeta
	var protocolAndJumps ProtocolAndJumps
	var whoisData whoisparser.WhoisInfo

	errors := []error{}

	for v := range multiplexedStream {
		if v.Error != nil {
			errors = append(errors, v.Error)
			continue
		}

		switch v.Data.(type) {
		case HTMLMeta:
			htmlMetaData = v.Data.(HTMLMeta)
		case ProtocolAndJumps:
			protocolAndJumps = v.Data.(ProtocolAndJumps)
		case whoisparser.WhoisInfo:
			whoisData = v.Data.(whoisparser.WhoisInfo)
		}
	}

	if len(errors) > 0 {
		return ScrapResponse{}, fmt.Errorf("errors: %v", errors)
	}

	return ScrapResponse{
		HTMLMetaData: htmlMetaData,
		Protocol:     protocolAndJumps.Protocol,
		Jumps:        protocolAndJumps.Jumps,
		WhoisData:    whoisData,
	}, nil
}

func GetHTMLMetaData(url string) <-chan concurrency.ChannelData {
	chann := make(chan concurrency.ChannelData)

	go func() {
		defer close(chann)
		resp, err := http.Get(url)
		if err != nil {
			chann <- concurrency.ChannelData{
				Error: err,
			}
			return
		}

		defer resp.Body.Close()

		meta := extract(resp.Body)

		chann <- concurrency.ChannelData{
			Data: meta,
		}
	}()

	return chann
}

func extract(resp io.Reader) HTMLMeta {
	z := html.NewTokenizer(resp)

	titleFound := false

	hm := new(HTMLMeta)

	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			return *hm
		case html.StartTagToken, html.SelfClosingTagToken:
			t := z.Token()
			if t.Data == `body` {
				return *hm
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

func GetProtocolAndJumps(url string) <-chan concurrency.ChannelData {
	chann := make(chan concurrency.ChannelData)

	go func() {
		defer close(chann)

		ssllabsUrl := fmt.Sprintf("https://api.ssllabs.com/api/v3/analyze?host=%s", url)

		resp, err := http.Get(ssllabsUrl)
		if err != nil {
			chann <- concurrency.ChannelData{
				Error: err,
			}
			return
		}

		defer resp.Body.Close()

		b, err := io.ReadAll(resp.Body)
		if err != nil {
			chann <- concurrency.ChannelData{
				Error: err,
			}
			return
		}

		sslLabsResponse := SSLLabsResponse{}

		err = json.Unmarshal(b, &sslLabsResponse)
		if err != nil {
			chann <- concurrency.ChannelData{
				Error: err,
			}
			return
		}

		chann <- concurrency.ChannelData{
			Data: ProtocolAndJumps{
				Protocol: sslLabsResponse.Protocol,
				Jumps:    len(sslLabsResponse.Endpoints),
			},
		}
	}()

	return chann
}

func GetWhoisData(url string) <-chan concurrency.ChannelData {
	chann := make(chan concurrency.ChannelData)

	go func() {
		defer close(chann)

		r, err := urlLib.Parse(url)
		if err != nil {
			chann <- concurrency.ChannelData{
				Error: err,
			}
			return
		}

		whois_raw, err := whois.Whois(r.Host)
		if err != nil {
			chann <- concurrency.ChannelData{
				Error: err,
			}
			return
		}

		result, err := whoisparser.Parse(whois_raw)

		if err != nil {
			chann <- concurrency.ChannelData{
				Error: err,
			}
			return
		}

		chann <- concurrency.ChannelData{
			Data: result,
		}
	}()

	return chann
}
