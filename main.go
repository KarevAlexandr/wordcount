package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	flag.Parse()
	if flag.Args()[0] == "" {
		fmt.Println(0)
	} else {
		slice := strings.Split(flag.Args()[0], " ")
		fmt.Println(len(slice))
	}

}
