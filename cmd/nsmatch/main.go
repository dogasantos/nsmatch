package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"

	nsmatch "github.com/dogasantos/nsmatch/pkg/runner"
)

// This is a nice comment to make lint happy. hello lint, i'm here!
type Options struct {
	TargetFile			string
	ResolverFile		string
	TrustedFile			string
	Version				bool
	Verbose				bool
}


var version = "0.1"

func parseOptions() *Options {
	options := &Options{}
	flag.StringVar(&options.TargetFile, 		"l", "", "List of domains we should test")
	flag.StringVar(&options.ResolverFile, 		"r", "", "List of dns servers we should test")
	flag.StringVar(&options.TrustedFile, 		"t", "", "List of ns hosts we should look for")
	flag.BoolVar(&options.Version, 				"i", false, "Version info")
	flag.BoolVar(&options.Verbose, 				"v", false, "Verbose mode")
	flag.Parse()
	return options
}

func main() {

	options := parseOptions()
	if options.Version {
		fmt.Println(version)
	}
	
	if options.ResolverFile != "" {
		if options.Verbose == true {
			fmt.Printf("[+] NSMATCH v%s\n",version)
		}
		TargetFilestream, _ := ioutil.ReadFile(options.TargetFile)
		targetContent := string(TargetFilestream)
		targets := strings.Split(targetContent, "\n")

		ResolverFilestream, _ := ioutil.ReadFile(options.ResolverFile)
		resolversContent := string(ResolverFilestream)
		resolvers := strings.Split(resolversContent, "\n")

		trustedFilestream, _ := ioutil.ReadFile(options.TrustedFile)
		trustedContent := string(trustedFilestream)
		trustedns := strings.Split(trustedContent, "\n")
		
		if options.Verbose == true {
			fmt.Printf("  + Resolvers loaded: %d\n",len(resolvers))
			fmt.Printf("  + Targets loaded: %d\n",len(targets))
			fmt.Printf("  + Trusted NS servers: %d\n",len(trustedns))
			fmt.Printf("  + Starting routines\n")
		}
		
		for _, target := range targets {
			if len(target) > 1 {
				fmt.Println(target)
				go nsmatch.Start(resolvers, target, trustedns, options.Verbose)
				fmt.Println("======")
			}
		}

	}
}




