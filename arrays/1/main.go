package main

import(
	"fmt"
	"time"
	"strings"
	"sort"
	"runtime"
)

func main(){
	s := "abcdefghaijklmnopqrstuvwxyz"
	fmt.Println(isUniqueWithAlphabet(s))	
	fmt.Println(isUniqueWithMap(s))	
	fmt.Println(isUniqueInPlace(s))	
	fmt.Println(isUniqueSort(s))	
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

func isUniqueWithAlphabet(s string) bool {
	defer timer()()
	if len(s) > 26 {
		return false
	}
	alph := [26]bool{}

	for i := 0; i < len(s); i++ {
		el := int(s[i]) - int('a')
		if !alph[el] {
			alph[el] = true
		} else {
			return false
		}
	}
	
	return true
}

func isUniqueWithMap(s string) bool {
	defer timer()()
	s = strings.ToUpper(s)
	charMap := make(map[rune]struct{})
	for _, r := range s {
		if _, ok := charMap[r]; !ok {
			charMap[r] = struct{}{}
		} else {
			return false
		}
	}

	return true
}

func isUniqueInPlace(s string) bool {
	defer timer()()
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s) - 1; j++ {
			if s[i] == s[j] {
				return false
			}
		}
	}

	return true
}

func isUniqueSort(s string) bool {
	defer timer()()
	stringSlice := []rune(s)
	sort.Slice(stringSlice, func(i, j int) bool {return stringSlice[i] < stringSlice[j]})

	for i := 1; i < len(stringSlice); i++ {
		if stringSlice[i] == stringSlice[i-1] {
			return false
		}
	}
	
	return true
}
