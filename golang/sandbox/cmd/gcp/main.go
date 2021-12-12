package main

import (
	"flag"
	"sandbox/gcp"
)

func main() {
	orgID := flag.String("org", "", "input organization id")
	flag.Parse()
	// gcp.MainResourceManager(*orgID)
	gcp.MainScc(*orgID)
}
