package scrapfranquise

import (
	"fmt"

	"github.com/likexian/whois"
)

func ScrapFranquise() {
	result, err := whois.Whois("marriott.com")
	if err == nil {
		fmt.Println("result", result)
	}

	fmt.Println("End")
}
