package main

import "fmt"

func main() {
	colors := make(map[int]string, 0)
	colors[0] = "hex code of red"
	colors[1] = "hex code of blue"
	fmt.Println(colors)
}