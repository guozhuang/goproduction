package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main () {
	//先顺序模型，然后counter操作使用原子操作
	var wg sync.WaitGroup
	wg.Add(10000)

	var counter int32

	for i:= 0; i< 10000; i++ {
		go func () {
			//而加锁是为了将一组操作原子化
			//虽然原子化操作已经涵盖了add，交换，以及配置的处理
			atomic.AddInt32(&counter, 1)

			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(counter)
}
