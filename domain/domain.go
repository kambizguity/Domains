package domain

import (
	"fmt"
	"whois/alphabetic"
)

var Domains []string

func makeDomain(content string) []string {
	var domains []string
	for _, v := range alphabetic.Alphabet {
		var domain string = fmt.Sprint(content, v)
		domain = fmt.Sprint("http://www.", domain, ".com")
		domains = append(domains, domain)
	}
	return domains
}

func AppendDomains(content string) {
	var tempDomains = makeDomain(content)
	fmt.Print(content, "\r")
	Domains = append(Domains, tempDomains...)
}
