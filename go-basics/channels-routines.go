package main

import (
	"fmt"
	"net/http"
	"time"
)

func concurrency() {

	links := []string {
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	//loopPing(c)
	//loopPing2(c)
	loopPingSleepWithFunctionLiteral(c)
}

func loopPingSleepWithFunctionLiteral(c chan string) {
	for l := range c { //for each message got from channel
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func loopPing2(c chan string) {
	for l := range c { //for each message got from channel
		go checkLink(l, c)
	}
}

func loopPing(c chan string) {
	//for i := 0; i < len(links); i++ {
	for { //infinite loop
		//fmt.Println(<-c)
		go checkLink(<-c, c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if(err != nil){
		fmt.Println(link, " might be down!")
		//c <- "ERROR - " + err.Error()
		c <- link
		return
	}

	fmt.Println(link, " is up!")
	//c <- "OK"
	c <- link
}