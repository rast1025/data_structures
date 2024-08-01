package main

import (
	"fmt"
)

func main(){
	arr := [][]int{
		{1, 2, 3, 0},
		{3, 5, 7, 9},
		{8, 0, 4, 2},
	}
	fmt.Println(nullifyArray(arr))
}

func nullifyArray(arr [][]int) [][]int {
	 row := make([]bool, len(arr))
	 column := make([]bool, len(arr[0]))

	 for i := 0; i < len(arr); i++ {
	 	for j := 0; j < len(arr[i]); j++ {
	 		if arr[i][j] == 0 {
	 			row[i] = true
	 			column[j] = true
	 		}
	 	}
	 }

	 for i := 0;  i < len(arr); i++ {
	 	for j := 0; j < len(arr[i]); j++ {
	 		if row[i] || column[j] {
	 			arr[i][j] = 0
	 		}
	 	}
	 }

	 return arr
}
