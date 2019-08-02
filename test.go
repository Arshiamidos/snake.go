package main

import "fmt"

func main() {
	ar := [][]int{
		{1, 2},
	}
	ar = append(ar, []int{3, 4})
	ar = append(ar, []int{5, 6})
	fmt.Println(ar)
}
