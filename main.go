package main

import (
	"flag"
	"log"
)

func main() {
	port := flag.Int("port", 8080, "port to listen on")
	flag.Parse()

	_, err := GetClient(*port)
	if err != nil {
		log.Fatalf("Failure creating oauth2 client: %v", err)
	}

}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
