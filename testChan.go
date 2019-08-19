package main

import (
	"fmt"
)

func main () {

	ch := make (chan int, 2)

	go func () {
		fmt.Println(<-ch)
	}()

	ch <- 3
	ch <- 4
	ch <- 5
	close(ch)
	fmt.Println("hello")

}
