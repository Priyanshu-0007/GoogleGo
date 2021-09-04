package main

import (
	"fmt"
	"os"
	"strings"
)

func main(){

	msg := os.Args[1]
	l := len(msg)
	s := strings.Repeat("!",l) + msg
	s = strings.ToUpper(s)
	fmt.Println(s)

}







