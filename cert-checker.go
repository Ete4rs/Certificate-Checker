package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"os"
	"regexp"
	"strings"
)

type CertData struct {
	CrtShID            string
	LoggedAt           string
	NotBefore          string
	NotAfter           string
	CommonName         string
	MatchingIdentities string
	IssuerName         string
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("please just enter a domain")
		return
	}
	checkDomain(os.Args[1])
	url := "https://crt.sh/?q=" + os.Args[1]
	var certs []CertData
	c := colly.NewCollector()
	c.OnHTML("table table tr", func(e *colly.HTMLElement) {
		var cert CertData
		e.ForEach("td", func(i int, td *colly.HTMLElement) {
			columnData := strings.TrimSpace(td.Text)
			switch i {
			case 0:
				cert.CrtShID = columnData
			case 1:
				cert.LoggedAt = columnData
			case 2:
				cert.NotBefore = columnData
			case 3:
				cert.NotAfter = columnData
			case 4:
				cert.CommonName = columnData
			case 5:
				cert.MatchingIdentities = columnData
			case 6:
				cert.IssuerName = columnData
			}
		})
		certs = append(certs, cert)
	})
	err := c.Visit(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	for _, c := range certs {
		fmt.Println(c)
	}
}

func checkDomain(d string) {
	regex := "([a-zA-Z0-9-]+)\\.[a-z]{2,}$"
	r, err := regexp.Compile(regex)
	if err != nil {
		panic(err)
	}
	if !r.MatchString(d) {
		fmt.Println("please provide a correct domain")
		return
	}
}
