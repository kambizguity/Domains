package domain

import (
	"fmt"
)

type Domain struct {
	prefix  string
	name    string
	pastfix string
	Status  int
	domain  string
}

var Domains map[string]Domain = make(map[string]Domain)

func (d *Domain) NewDomain(prefix string, name string, pastfix string) {
	d.prefix = prefix
	d.name = name
	d.pastfix = pastfix
	d.domain = fmt.Sprint(prefix, ".", name, ".", pastfix)
	d.Status = 0
}
func (d Domain) GetKey() string {
	return fmt.Sprint(d.name, ".", d.pastfix)
}

// func makeDomain(content string) []string {
// 	var domains []string
// 	for _, v := range alphabetic.Alphabet {
// 		var domain string = fmt.Sprint(content, v)
// 		domain = fmt.Sprint("http://www.", domain, ".com")
// 		domains = append(domains, domain)
// 	}
// 	return domains
// }

// func AppendDomains(content string) {
// 	var tempDomains = makeDomain(content)
// 	fmt.Print(content, "\r")
// 	Domains = append(Domains, tempDomains...)
// }
