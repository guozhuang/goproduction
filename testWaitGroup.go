package main

import (
	"fmt"
	"sync"
)

func main () {

	var wg sync.WaitGroup
	wg.Add(10000)//升级成1w时，得到的结果就不一致起来

	var counter int

	for i:=0; i < 10000; i++ {
		go func() {
			counter ++//原因就是counter++操作在cpu看来不是原子的
			//fmt.Println(counter)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(counter)
}
