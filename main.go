package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"sync"
    "flag"
)

func main() {
	var domain_input = flag.String("d", "", "Domain to generate subdomains with")
    flag.Parse()
    domain := *domain_input
    fmt.Println(domain)
	processed := make(chan string)
	subdomains := make(chan string)
	domains := make(map[string]struct{})
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func() {
			for subdomain := range subdomains {
				reg, _ := regexp.Compile("[^a-zA-Z0-9-.]+")
				record := reg.ReplaceAllString(strings.ToLower(subdomain+"."+domain), "")
				processed <- record

			}

			wg.Done()
		}()
	}

	go func() {
		sc := bufio.NewScanner(os.Stdin)
		for sc.Scan() {
			subdomains <- sc.Text()
		}
		close(subdomains)
	}()
	go func() {
		wg.Wait()
		close(processed)
	}()
	for domain := range processed {
		_, ok := domains[domain]
		if !ok {
			domains[domain] = struct{}{}
			fmt.Println(domain)
		}
	}

}
