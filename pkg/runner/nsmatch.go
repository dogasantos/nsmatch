package nsmatch

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"

	"github.com/miekg/dns"
	"github.com/projectdiscovery/retryabledns"
)

// This is a nice comment to make lint happy. hello lint, i'm here!
func DoResolve(target string, resolver string, trustedns []string) {
	var resolvers []string
	resolvers = append(resolvers,resolver+":53")

	dnsClient := retryabledns.New(resolvers, 2)
	dnsResponses, _ := dnsClient.Query(target, dns.TypeNS)

	if len(trustedns) > 0 {
		for _, nsfound := range dnsResponses.NS {
			for _, trusted := range trustedns {
				trusted = strings.ReplaceAll(trusted, " ", "")
				nsfound = strings.ReplaceAll(nsfound, " ", "")
				if len(trusted) > 2 {
					if strings.Contains(nsfound,trusted) {
						fmt.Printf("%s:%s,%s\n",target,nsfound,trusted)
					} else if nsfound == trusted {
						fmt.Printf("%s:%s,%s\n",target,nsfound,trusted)
					}
				}
			}
		}
	}
	return
}

func getRandomResolver(resolvers []string) string {
	rand.Seed(time.Now().UnixNano())
	randIdx := rand.Intn(len(resolvers))

	return resolvers[randIdx]
}

func Start(resolvers []string, target string, trustedns []string, verbose bool, wg *sync.WaitGroup) {
	var resolver string

	defer wg.Done()
	for {
		resolver = getRandomResolver(resolvers)
		if len(resolver) > 2{ 
			break
		}
	}

	if verbose {
		fmt.Printf("  + Testing: %s using %s looking for any of %s\n",target, resolver, trustedns)
	}
	
	DoResolve(target,resolver,trustedns)
	
}
	