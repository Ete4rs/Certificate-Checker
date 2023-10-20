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
	flags := ResultFLag{}
	flags.CrtShID = *flag.Bool("c", false, "cert ID")
	flags.LoggedAt = *flag.Bool("l", false, "logged At")
	flags.NotBefore = *flag.Bool("b", false, "not exist before")
	flags.NotAfter = *flag.Bool("a", false, "not exist after")
	flags.CommonName = *flag.Bool("n", false, "common name")
	flags.MatchingIdentities = *flag.Bool("m", false, "Matching Identities")
	flags.IssuerName = *flag.Bool("i", false, "issuer")
	flags.Target = *flag.String("t", "", "domain")
	help := flag.Bool("h", false, "Show help")
	flag.Parse()
	if *help {
		flag.PrintDefaults()
		os.Exit(0)
	}
	checkDomain(flags.Target)
	flags.Target = "https://crt.sh/?q=" + flags.Target
	if flags.IssuerName || flags.CommonName || flags.MatchingIdentities || flags.NotBefore || flags.LoggedAt ||
		flags.CrtShID || flags.NotAfter {
		flags.Base = true
	} else {
		flags.Base = false
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
