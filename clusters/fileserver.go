package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	directory := "/Users/nithin.reddy/Documents/GitHub/Flux-tutorial/clusters/helm-charts"
	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir(directory)))

	log.Printf("Serving %s on HTTP port: 8100\n", directory)
	log.Fatal(http.ListenAndServe(":8100", nil))
}
