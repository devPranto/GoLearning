package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	s := strings.Split(time.Now().String(), " ")
	fmt.Println(s[0])
	t := time.Now()
	i := int(t.Day())
	fmt.Println(i)
}
