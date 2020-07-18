package Records

import "fmt"

type GoroutinePool struct {
	Queue chan func() error    // 将任务(函数)存入channel
	Number int
	Total int
	result chan error //结果存入另一个管道
	finishCallback func()
}

func (goPool *GoroutinePool)Init (number ,total int){
	goPool.Number = number
	goPool.Total = total
	goPool.Queue = make(chan func() error,total)
	goPool.result = make(chan error,total)
}

func (goPool *GoroutinePool)Start(){
	// 跑起来
	for i:=0;i<goPool.Number;i++{
		go func() {
			for {
				task,ok := <-goPool.Queue
				if !ok{ // Queue中没有task了
					break
				}
				err := task()
				goPool.result <- err
			}
		}()
	}
	//获取结果
	for j:=0;j<goPool.Total;j++{
		res,ok := <-goPool.result
		if !ok{   // result里没结果了
			break
		}
		if res != nil{ //res = error
			fmt.Print(res)
		}
	}
	    // 所有任务都执行完成，回调函数
	if goPool.finishCallback != nil {
		goPool.finishCallback()
	}
}

func (goPool *GoroutinePool)Stop(){
	close(goPool.Queue)
	close(goPool.result)
}

func (goPool *GoroutinePool)AddTask(task func() error){
	goPool.Queue <- task
}

func (goPool *GoroutinePool)SetFinishCallback(callback func()){
	goPool.finishCallback = callback
}

