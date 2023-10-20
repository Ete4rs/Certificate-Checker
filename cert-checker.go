package main

import (
	"cert-checker/argsparse"
	"fmt"
	"github.com/gocolly/colly/v2"
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
	Scan := argsparse.ArgumentParser()
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
	err := c.Visit(Scan.Target)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	for _, c := range certs {
		if Scan.Base {
			fmt.Println(c.CrtShID, "    ", c.LoggedAt, "    ", c.NotBefore, "    ", c.NotAfter, "    ",
				c.CommonName, "    ", c.MatchingIdentities, "    ", c.IssuerName)
			continue
		}
		if Scan.CrtShID {
			fmt.Print(c.CrtShID, "    ")
		}
		if Scan.LoggedAt {
			fmt.Print(c.LoggedAt, "    ")
		}
		if Scan.NotBefore {
			fmt.Print(c.NotAfter, "    ")
		}
		if Scan.NotAfter {
			fmt.Print(c.NotAfter, "    ")
		}
		if Scan.CommonName {
			fmt.Print(c.CommonName, "    ")
		}
		if Scan.MatchingIdentities {
			fmt.Print(c.MatchingIdentities, "    ")
		}
		if Scan.IssuerName {
			fmt.Print(c.IssuerName, "\n")
		}
	}
}
