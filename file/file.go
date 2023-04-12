package file

import (
	"fmt"
	"log"
	"os"
	"time"
	"whois/domain"
)

func MakeCSVFile() *os.File {
	dateTime := time.Now()
	fileName := fmt.Sprint("csv/domains_", dateTime.Year(), "_", dateTime.Month(), "_", dateTime.Day(), ".csv")
	f, err := os.Create(fileName)
	if err != nil {
		log.Fatal("Error to make new csv file, Error:", err)
		return nil
	}
	return f
}
func AppendIntoFile(f *os.File) {
	fmt.Println("AppendIntoFile")
	var content string = ""
	fmt.Println("Iteration on array")
	for _, v := range domain.Domains {
		fmt.Printf("%v\r", v)
		content = fmt.Sprint(content, v, "\n")
	}
	fmt.Println("Writing into file")
	_, err := f.WriteString(content)
	if err != nil {
		fmt.Println(err)
		log.Panic(err)
	} else {
		domain.Domains = nil
	}
	fmt.Println("domains has appended into file")

}
