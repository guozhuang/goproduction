package main

import "fmt"

func main() {
	//test slice extension
	arr := [...]int{0,1,2,3,4,5,6,7}
	s1 := arr[2:6]
	s2 := s1[3:6]

	fmt.Println("s1:", s1)//s1: [2 3 4 5]
	fmt.Println("s2", s2)//s2 [5 6 7]
}
