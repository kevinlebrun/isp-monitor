package main

import (
	"flag"
	"fmt"
	"log"
)

func DoPing(count int, destination string, result chan<- string) {
	ping, err := ExecPing(count, destination)
	if err != nil {
		// TODO make this less deadly
		log.Fatal(err)
	}

	result <- ping
}

func main() {
	var withHeaders = flag.Bool("with-headers", false, "Show headers")
	flag.Parse()

	result := make(chan string, 3)

	go DoPing(5, "www.google.fr", result)
	go DoPing(5, "www.meetic.fr", result)
	go DoPing(5, "192.168.1.1", result)

	if *withHeaders {
		fmt.Println(PingHeaders())
	}

	for i := 0; i < 3; i++ {
		fmt.Println(<-result)
	}
}
