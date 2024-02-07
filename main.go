package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	fmt.Println("Hello Akash")

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("domain, hasMX, hasSPF, sprRecord, hasDMARC, dmarcRecord")
	fmt.Print("Enter your Domain	:	")
	scanner.Scan()

	checkDomainMx(scanner.Text())

}

func checkDomainMx(domain string) {
	var hasMx, hasSPF, hasDMARC bool // hasSPF, hasDMARC bool
	//var domain, sprRecord, dmarcRecord string
	var sprRecord, dmarcRecord string
	res, _ := net.LookupMX(domain)

	if len(res) > 0 {
		hasMx = true
	}
	r, _ := net.LookupTXT(domain)

	for _, val := range r {
		if strings.HasPrefix(val, "v=spf") {
			hasSPF = true
			sprRecord = val
			break
		}
	}
	dmarcsRecords, _ := net.LookupTXT("_dmarc." + domain)

	for _, val := range dmarcsRecords {

		if strings.HasPrefix(val, "v=DMARC1") {
			dmarcRecord = val
			hasDMARC = true
		}
	}

	fmt.Println("domain", domain)
	fmt.Println("HasMX", hasMx) 
	fmt.Println("HasSPF", hasSPF)
	fmt.Println("HasDMARC", hasDMARC)
	fmt.Println("SPRecord", sprRecord)
	fmt.Println("DMARCRecord", dmarcRecord)

}
