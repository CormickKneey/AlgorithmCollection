package Records

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func TestGoroutinePool(t *testing.T) {
	goPool := new(GoroutinePool)
	goPool.Init(3,5)
	targets := []int{1,4,24,24,100,71,29,54,12,23,44,51,532,99,91,671,82}
	for i:=0;i<len(targets);i++{
		goPool.AddTask(func() error {
			return TroubleMaker(targets[i])
		})
	}
	isFinished := false
	goPool.SetFinishCallback(func() {
		func(isFinished *bool){
			*isFinished = true
		}(&isFinished)
	})
	goPool.Start()

	for !isFinished{
		time.Sleep(time.Millisecond * 100)
	}
	goPool.Stop()
	fmt.Printf("Done!")
}



func TroubleMaker(num int)error{
	fmt.Printf("The number get : %d", num)
	if num > 99{
		return errors.New("The number is too big! ")
	}
	return nil
}
