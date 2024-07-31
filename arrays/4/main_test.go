package main

import(
	"testing"
	"fmt"
)

func TestUrlify(t *testing.T) {
	var tests = []struct {
		str string
		want string
	}{
		{"Mr John Smith    ", "Mr%20John%20Smith"},
		{"My test  ", "My%20test"},
	}

	for _, tt := range tests {
		str := tt.str
		testName := fmt.Sprintf("%s, %s", tt.str, tt.want)
		t.Run(testName, func(t *testing.T){
			urlify(&str)
			if str != tt.want {
				t.Errorf("got %s, want %s", str, tt.want)
			}
		})
	}
	
}
