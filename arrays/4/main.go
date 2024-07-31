package main

import (
	"strings"
)

func main(){
	input := "Mr John Smith    "
	urlify(&input)
}

func urlify(s *string) {
	length := len(strings.Trim(*s, " "))
	newLength := len(*s)
	r := []rune(*s)

	for i := length-1; i >=0; i-- {
		if r[i] == ' ' {
			r[newLength - 1] = '0'
			r[newLength - 2] = '2'
			r[newLength - 3] = '%'
			newLength -= 3
		} else {
			r[newLength - 1] = r[i]
			newLength -= 1
		}
	}
	*s = string(r)
}

