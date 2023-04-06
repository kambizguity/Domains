package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var domains []string
var alphabet []string = alphabetGenerator()

func main() {
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	} else {
		defer f.Close()

		appendDomains("")
		appendIntoFile(f)

		for _, valueA := range alphabet {
			appendDomains(valueA)
		}
		appendIntoFile(f)

		for _, valueA := range alphabet {
			for _, valueB := range alphabet {
				tempString := fmt.Sprint(valueA, valueB)
				appendDomains(tempString)
			}
		}
		appendIntoFile(f)

		for _, valueA := range alphabet {
			for _, valueB := range alphabet {
				for _, valueC := range alphabet {
					tempString := fmt.Sprint(valueA, valueB, valueC)
					appendDomains(tempString)
				}
			}
		}
		appendIntoFile(f)

		for _, valueA := range alphabet {
			fmt.Printf("valueA: %v\r", valueA)
			for _, valueB := range alphabet {
				fmt.Printf("valueB: %v\r", valueB)

				for _, valueC := range alphabet {
					fmt.Printf("valueC: %v\r", valueC)
					for _, valueD := range alphabet {
						fmt.Printf("valueD: %v\r", valueD)
						tempString := fmt.Sprint(valueA, valueB, valueC, valueD)
						appendDomains(tempString)
					}
				}
			}
		}
		appendIntoFile(f)

		for _, valueA := range alphabet {
			for _, valueB := range alphabet {
				for _, valueC := range alphabet {
					for _, valueD := range alphabet {
						for _, valueE := range alphabet {
							tempString := fmt.Sprint(valueA, valueB, valueC, valueD, valueE)
							appendDomains(tempString)
						}
					}
				}
			}
		}
		appendIntoFile(f)

		for _, valueA := range alphabet {
			for _, valueB := range alphabet {
				for _, valueC := range alphabet {
					for _, valueD := range alphabet {
						for _, valueE := range alphabet {
							for _, valueF := range alphabet {
								tempString := fmt.Sprint(valueA, valueB, valueC, valueD, valueE, valueF)
								appendDomains(tempString)
							}
						}
					}
				}
			}
		}
		appendIntoFile(f)

		for _, valueA := range alphabet {
			for _, valueB := range alphabet {
				for _, valueC := range alphabet {
					for _, valueD := range alphabet {
						for _, valueE := range alphabet {
							for _, valueF := range alphabet {
								for _, valueG := range alphabet {
									tempString := fmt.Sprint(valueA, valueB, valueC, valueD, valueE, valueF, valueG)
									appendDomains(tempString)
								}
							}
						}
					}
				}
			}
		}
		appendIntoFile(f)
	}
}

func makeDomain(content string) []string {
	var domains []string
	for _, v := range alphabet {
		var domain string = fmt.Sprint(content, v)
		domain = fmt.Sprint("http://www.", domain, ".com")
		domains = append(domains, domain)
	}
	return domains
}

func appendDomains(content string) {
	var tempDomains = makeDomain(content)

	domains = append(domains, tempDomains...)
}

func alphabetGenerator() []string {
	var alphabet []string

	for ch := 'a'; ch <= 'z'; ch++ {
		var char string = strconv.QuoteRune(ch)
		char = strings.ReplaceAll(char, "'", "")
		alphabet = append(alphabet, char)
	}
	for i := 0; i <= 9; i++ {
		alphabet = append(alphabet, fmt.Sprint(i))
	}

	return alphabet
}
func appendIntoFile(f *os.File) {
	var content string = ""
	for _, v := range domains {
		content = fmt.Sprint(content, v, "\n")
	}

	_, err := f.WriteString(content)
	if err != nil {
		fmt.Println(err)
		log.Panic(err)
	} else {
		domains = nil
	}
}
