package main

import (
	"fmt"
	"whois/alphabetic"
	"whois/domain"
	"whois/file"
)

func main() {
	f := file.MakeCSVFile()
	if f == nil {
		return
	} else {
		defer f.Close()
		for _, valueA := range alphabetic.Alphabet {
			for _, valueB := range alphabetic.Alphabet {
				for _, valueC := range alphabetic.Alphabet {
					tempString := fmt.Sprint(valueA, valueB, valueC)
					domain.AppendDomains(tempString)
				}
			}
		}
		file.AppendIntoFile(f)

		// for _, valueA := range alphabetic.Alphabet {
		// 	for _, valueB := range alphabetic.Alphabet {
		// 		for _, valueC := range alphabetic.Alphabet {
		// 			for _, valueD := range alphabetic.Alphabet {
		// 				for _, valueE := range alphabetic.Alphabet {
		// 					tempString := fmt.Sprint(valueA, valueB, valueC, valueD, valueE)
		// 					domain.AppendDomains(tempString)
		// 				}
		// 			}
		// 		}
		// 	}
		// }
		// file.AppendIntoFile(f)

		// 	for _, valueA := range alphabetic.Alphabet {
		// 		for _, valueB := range alphabetic.Alphabet {
		// 			for _, valueC := range alphabetic.Alphabet {
		// 				for _, valueD := range alphabetic.Alphabet {
		// 					for _, valueE := range alphabetic.Alphabet {
		// 						for _, valueF := range alphabetic.Alphabet {
		// 							tempString := fmt.Sprint(valueA, valueB, valueC, valueD, valueE, valueF)
		// 							appendDomains(tempString)
		// 						}
		// 					}
		// 				}
		// 			}
		// 		}
		// 	}
		// 	appendIntoFile(f)

		// 	for _, valueA := range alphabetic.Alphabet {
		// 		for _, valueB := range alphabetic.Alphabet {
		// 			for _, valueC := range alphabetic.Alphabet {
		// 				for _, valueD := range alphabetic.Alphabet {
		// 					for _, valueE := range alphabetic.Alphabet {
		// 						for _, valueF := range alphabetic.Alphabet {
		// 							for _, valueG := range alphabetic.Alphabet {
		// 								tempString := fmt.Sprint(valueA, valueB, valueC, valueD, valueE, valueF, valueG)
		// 								appendDomains(tempString)
		// 							}
		// 						}
		// 					}
		// 				}
		// 			}
		// 		}
		// 	}
		// 	appendIntoFile(f)
	}

}
