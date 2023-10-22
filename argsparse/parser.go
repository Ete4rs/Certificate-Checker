package argsparse

import (
	"flag"
	"fmt"
	"os"
	"regexp"
)

type ResultFLag struct {
	CrtShID            bool
	LoggedAt           bool
	NotBefore          bool
	NotAfter           bool
	CommonName         bool
	MatchingIdentities bool
	IssuerName         bool
	Target             string
	Base               bool
}

func ArgumentParser() *ResultFLag {
	certID := flag.Bool("c", false, "cert ID")
	loggedAt := flag.Bool("l", false, "logged At")
	notBefore := flag.Bool("b", false, "not exist before")
	notAfter := flag.Bool("a", false, "not exist after")
	commonName := flag.Bool("n", false, "common name")
	matchingIdentities := flag.Bool("m", false, "Matching Identities")
	issuerName := flag.Bool("i", false, "issuer")
	target := flag.String("t", "", "domain")
	help := flag.Bool("h", false, "Show help")
	flag.Parse()
	if *help {
		flag.PrintDefaults()
		os.Exit(0)
	}
	checkDomain(*target)

	flags := ResultFLag{}
	flags.CrtShID = *certID
	flags.LoggedAt = *loggedAt
	flags.NotBefore = *notBefore
	flags.NotAfter = *notAfter
	flags.CommonName = *commonName
	flags.MatchingIdentities = *matchingIdentities
	flags.IssuerName = *issuerName
	flags.Target = "https://crt.sh/?q=" + *target

	if flags.IssuerName || flags.CommonName || flags.MatchingIdentities || flags.NotBefore || flags.LoggedAt ||
		flags.CrtShID || flags.NotAfter {
		flags.Base = false
	} else {
		flags.Base = true
	}
	return &flags
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
