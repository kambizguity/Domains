package main

import (
	"fmt"
	"sort"
	"sync"
	"whois/alphabetic"
	"whois/domain"
	"whois/file"
)

type mutexDomain struct {
	Domains map[string]domain.Domain
	sync.Mutex
}

func main() {
	var wg sync.WaitGroup
	var domains mutexDomain

	f := file.MakeCSVFile()
	if f == nil {
		fmt.Println("Error")
		return
	} else {
		defer f.Close()
		domains.Domains = make(map[string]domain.Domain)
		wg.Add(1)
		fmt.Println("Three chars")
		go func() {
			for _, valueA := range alphabetic.Alphabet {
				for _, valueB := range alphabetic.Alphabet {
					for _, valueC := range alphabetic.Alphabet {
						tempString := fmt.Sprint(valueA, valueB, valueC)
						addDomain(&domains, tempString)
					}
				}
			}
			wg.Done()
		}()
		fmt.Println("Four chars")
		wg.Add(1)
		go func() {
			for _, valueA := range alphabetic.Alphabet {
				for _, valueB := range alphabetic.Alphabet {
					for _, valueC := range alphabetic.Alphabet {
						for _, valueD := range alphabetic.Alphabet {
							tempString := fmt.Sprint(valueA, valueB, valueC, valueD)
							addDomain(&domains, tempString)
						}
					}
				}
			}
			wg.Done()
		}()
		wg.Wait()
		fmt.Println("Coping...")
		domain.Domains = domains.Domains
		fmt.Printf("len(domain.Domains): %v\n", len(domain.Domains))

		keys := make([]string, 0, len(domain.Domains))
		for k := range domain.Domains {
			keys = append(keys, k)
		}

		sort.Strings(keys)
		fmt.Printf("len(keys): %v\n", len(keys))
		fmt.Println(keys[len(keys)-1])
		fmt.Printf("keys[0]: %v\n", keys[0])

	}
}
func addDomain(to *mutexDomain, with string) {
	to.Lock()
	var newDomain domain.Domain
	newDomain.NewDomain("www", with, "com")
	to.Domains[newDomain.GetKey()] = newDomain
	to.Unlock()
}

// for _, valueA := range alphabetic.Alphabet {
// 	for _, valueB := range alphabetic.Alphabet {
// 		for _, valueC := range alphabetic.Alphabet {
// 			for _, valueD := range alphabetic.Alphabet {
// 				for _, valueE := range alphabetic.Alphabet {
// 					for _, valueF := range alphabetic.Alphabet {
// 						for _, valueG := range alphabetic.Alphabet {
// 							tempString := fmt.Sprint(valueA, valueB, valueC, valueD, valueE, valueF, valueG)
// 							domain.AppendDomains(tempString)
// 						}
// 					}
// 				}
// 			}
// 		}
// 	}
// }
