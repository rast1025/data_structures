package main

import (
	"fmt"
)

func main(){
	s1 := "hello"
	s2 := "ollhe"
	fmt.Println(isPermutation(s1,s2))
}



func isPermutation(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	hashMap := make(map[byte]int, 256)

	for i := 0; i < len(s1); i++ {
		hashMap[s1[i]]++
		hashMap[s2[i]]--
	}

	for _, v := range hashMap {
		if v != 0 {
			return false
		}
	}

	return true
}
