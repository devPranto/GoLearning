package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	Scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Domain,hasMX,hasSPF,sprRecord,hasDMARC,dmarcRecord")
	for Scanner.Scan() {
		checkDomain(Scanner.Text())
		//break
	}
	if err := Scanner.Err(); err != nil {
		log.Fatal("couldn't read from the user input")
	}
}
func checkDomain(domain string) {
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	mxRecords, err := net.LookupMX(domain)

	if err != nil {
		log.Printf("erroor : %v \n", err)
	}
	if len(mxRecords) > 0 {
		hasMX = true
	}
	txtRecords, err := net.LookupTXT(domain)

	if err != nil {
		log.Printf("erroor : %v \n", err)
	}
	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}
	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("error : %v \n", err)
	}
	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}
	fmt.Printf("hasMX = %v , hasSPF = %v, hasDMARC = %v , spfRecord = %v , dmarcRecord = %v \n", hasMX, hasSPF, hasDMARC, spfRecord, dmarcRecord)
}
