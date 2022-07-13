package nscheck

import (
	"fmt"

	"github.com/miekg/dns"
	"github.com/projectdiscovery/retryabledns"
)

// This is a nice comment to make lint happy. hello lint, i'm here!
func DoResolve(hostname string, resolver string) []string{
	var responseValue []string
	var resolvers []string
	resolvers = append(resolvers,resolver+":53")

	dnsClient := retryabledns.New(resolvers, 2)
	dnsResponses, _ := dnsClient.Query(hostname, dns.TypeNS)
	
	fmt.Printf("%s:%s",dnsResponses.NS)
	return 
}


// This is a nice comment to make lint happy. hello lint, i'm here!
/*
func CheckResolverOld(resolver string, trusted_rdata string, wg * sync.WaitGroup, verbose bool) {
	defer wg.Done()

	test := DoResolve("nscheck.pingback.me",resolver)[0] 
	if strings.Contains("8.8.8.8",test) {
		if verbose {
			fmt.Printf("%s:INCORRECT\n",resolver)
		}
	} else {
		fmt.Printf("%s:CORRECT\n",resolver)
	}

}
*/

// resolver, targets, trustedns, options.Verbose
func Start(resolver string, target string, trustedns []string, verbose bool) {

		test := DoResolve(target,resolver)[0]

		
	}
	