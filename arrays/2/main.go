package main

import (
	"fmt"
	"runtime"
	"time"
	"strings"
)

func main(){
	fmt.Println(reverse("hellosd;lkfjdsklfgdjslfdsfkldsjlkfdsjklfdjklsdfkljsklfdjklasjkldjklasds"))
	fmt.Println(reverseStringBuilder("hellosd;lkfjdsklfgdjslfdsfkldsjlkfdsjklfdjklsdfkljsklfdjklasjkldjklasds"))
	b := "hello"
	reverseInPlace(&b)
	fmt.Println(b)
}

func reverse(s string) string {
	defer timer()()
	var result string

	for i := len(s) -1; i >= 0; i-- {
		result += string(s[i])
	}

	return result
}

func reverseStringBuilder(s string) string {
	defer timer()()
	b := strings.Builder{}
	b.Grow(len(s))
	for i := len(s)-1; i >= 0; i-- {
		b.WriteByte(s[i])
	}

	return b.String()
}

func reverseInPlace(s *string) {
	r := []rune(*s)
	for i, j := 0, len(r) -1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	*s = string(r)
}

func callerName(skip int) string {
    const unknown = "unknown"
    pcs := make([]uintptr, 1)
    n := runtime.Callers(skip+2, pcs)
    if n < 1 {
        return unknown
    }
    frame, _ := runtime.CallersFrames(pcs).Next()
    if frame.Function == "" {
        return unknown
    }
    return frame.Function
}

func timer() func() {
    name := callerName(1)
    start := time.Now()
    return func() {
        fmt.Printf("%s took %v\n", name, time.Since(start))
    }
}
