package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
)

func main() {
	arg := os.Args[1]

	u, err := url.Parse(arg)
	if err != nil {
		log.Fatal(err)
	}
	u.RawQuery = ""
	fmt.Println(u)
}
