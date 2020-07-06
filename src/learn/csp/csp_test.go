package csp

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 500)
	return "Done"
}

func AsyncService() chan string {
	//retCh := make(chan string)
	//构建buffer channel
	retCh := make(chan string, 1)
	go func() {
		//调用service
		ret := service()
		fmt.Println("returned result.")
		retCh <- ret
		fmt.Println("service exited.")
	}()
	return retCh
}

func otherTask() {
	fmt.Println("working on something else")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Task is done.")
}

//同步调用
func TestService(t *testing.T) {
	fmt.Println(service())
	otherTask()
}

//CSP调用
func TestAsyncService(t *testing.T) {
	retCh := AsyncService()
	//立即执行
	otherTask()
	fmt.Println(<-retCh)
}

//select 实现超时控制
func TestSelect(t *testing.T) {
	select {
	case ret := <-AsyncService():
		t.Log(ret)
	case <-time.After(time.Millisecond * 100):
		t.Error("time out")
	}
}
