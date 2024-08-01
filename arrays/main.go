package main

import (
	"fmt"
	"strings"
)

func main(){
	s1 := "waterbottle"
	s2 := "tlewaterbot"

	fmt.Println(isRotation(s1,s2))
}

func isRotation(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	doubleString := s1 + s1

	return strings.Contains(doubleString, s2)
	 
}
