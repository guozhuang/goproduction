package main

import "fmt"

func printfSliceCap(s []int){
	//可以看出，这里的打印本身就是一个interface{}
	fmt.Printf("s len(s)=%d, cap(s)=%d \n", len(s), cap(s))
}

func main() {
	//test slice cap
	var s []int

	for i:=0; i<100; i++ {
		printfSliceCap(s)
		s = append(s, i*2+1)//这里的i如果前面有一个1，则表示复数i
	}
}
