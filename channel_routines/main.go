package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"https://www.facebook.com/",
		"https://pkg.go.dev/std",
		"https://google.com",
	}
	for _, link := range links {
		checkLink(link)

	}

}
func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " link might be down")
		return
	}
	fmt.Println(link, "link is up ")

}
