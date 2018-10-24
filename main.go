package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	port := flag.Int("port", 8080, "port to listen on (default: 8080; use 0 to skip http listener)")
	flag.Parse()

	c, err := GetClient(*port)
	if err != nil {
		log.Fatalf("Failure creating oauth2 client: %v", err)
	}

	resp, err := c.Get("https://photoslibrary.googleapis.com/v1/albums")
	if err != nil {
		fmt.Printf("Error response from GET request. err:%v\n", err)
	} else {
		r, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Error Reading GET Response. err:%v\n", err)
		} else {
			fmt.Printf("GET Response: %s\n\n", r)
		}
	}
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
