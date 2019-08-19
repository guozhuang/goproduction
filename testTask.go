package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//假设发送从1到1w的id，需要发送优惠卷，此时分为十个任务来处理。

type task struct {
	taskId int
}

func main () {
	//将task放到相应结构中，并且同步保证相应任务处理
	taskch := make (chan task, 10)
	var wg sync.WaitGroup
	wg.Add(10000)

	var counter int32

	//此时相当于抢占式的协程调度
	for i:= 0; i< 10; i ++ {
		go func(workerId int) {
			for {
				ch := <-taskch
				for j := 1; j < 1001; j++ {
					uid := ch.taskId*1000 + j
					ch.sendData(uid, workerId)

					atomic.AddInt32(&counter, 1)
					wg.Done()
				}
			}
		}(i)
	}

	for i := 0; i < 10; i++ {
		//将值塞入task
		tmptask := task{
			i}

		taskch <- tmptask
	}

	wg.Wait()
	fmt.Println(counter)
}



//假定这个就是发送的方法
func (task *task) sendData (uid int, workerId int) {
	fmt.Printf("%d输出%d,workerId:%d", task.taskId, uid, workerId)
	fmt.Println()
}
