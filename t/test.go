package main

import "fmt"

func main() {
	/* ar := [][]int{
		{1, 2},
	}
	ar = append(ar, []int{3, 4})
	ar = append(ar, []int{5, 6})
	fmt.Println(ar) */
	newHead := make([]int, 1)
	copy(newHead, []int{5, 5})
	//copy(dst, src)
	fmt.Println(newHead[1:])
}
