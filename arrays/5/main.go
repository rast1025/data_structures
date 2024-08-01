package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main(){
	s := "aabcccccaaa"
	compressed := compress(s)
	fmt.Println(compressed)
	decompressed := decompress(compressed)
	fmt.Println(decompressed)
	fmt.Println(s == decompressed)
}

func compress(s string) string {
	if len(s) == 0 {
		return s
	}
	count := 1
	builder := strings.Builder{}
	builder.Grow(len(s))
	builder.WriteByte(s[0])
	
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			count++
		} else {
			countStr := strconv.Itoa(count)
			builder.WriteString(countStr)
			count = 1
			builder.WriteByte(s[i])
		}
	}
	countStr := strconv.Itoa(count)
	builder.WriteString(countStr)
	
	result := builder.String()

	if len(result) >= len(s) {
		return s
	}

	return result
}

func decompress(s string) string {
	var result strings.Builder
	
	for i := 0; i < len(s); {
		char := s[i]
		i++
		
		countStart := i
		for i < len(s) && isDigit(s[i]) {
			i++
		}
		count, _ := strconv.Atoi(s[countStart:i])
		
		for j := 0; j < count; j++ {
			result.WriteByte(char)
		}
	}
	
	return result.String()
}


func isDigit(b byte) bool {
    return b >= '0' && b <= '9'
}
