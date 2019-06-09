package main

import "fmt"

func main() {
	var s1 []int
	s2 := []int{2,3,4}

	s1 = append(s1, 1)
	s1 = append(s1, s2[0:]...)

	fmt.Println(s1)//[1 2 3 4]
}