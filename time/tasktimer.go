package main

import (
	"fmt"
	"time"
)

type TaskTimer struct {
	Interval int64     //执行间隔，单位秒
	ExecFunc func()    //调度执行方法
	ExecTime string    //执行时间，格式yyyy-MM-dd hh-mm-ss
	wait     chan bool //阻塞chan，在start中阻塞，防止主线程结束导致程序结束
	IsStop   bool      //结束标记
}

/**
 * 创建TaskTimer
 */
func CreateTaskTimer(execTime string, interval int64, execFunc func()) *TaskTimer {
	t := &TaskTimer{
		Interval: interval,
		ExecTime: execTime,
		ExecFunc: execFunc,
		wait:     make(chan bool, 1),
	}
	return t
}

/**
 * 启动任务
 */
func (t *TaskTimer) Start() {
	//计算首次执行时间
	tnow := time.Now()
	exeTime, _ := time.ParseInLocation("2006-01-02 15:04:05", t.ExecTime, tnow.Location())
	diff := exeTime.Unix() - tnow.Unix()
	for diff < 0 { //小于0则加一个间隔时间,直到大于0为止
		diff += t.Interval
	}

	time.AfterFunc(time.Duration(diff)*time.Second, func() {
		if !t.IsStop {
			t.ExecFunc() //第一次执行
		} else {
			return
		}
		for {
			timer := time.NewTimer(time.Duration(t.Interval) * time.Second) //每隔
			select {
			case <-timer.C:
				if !t.IsStop {
					t.ExecFunc()
				} else {
					return
				}
			}
		}
	})

	//阻塞
	// <-t.wait
}

func (t *TaskTimer) Stop() {
	t.IsStop = true
}

/**
 * 用来阻塞，在start之后调用，防止主线程结束
 */
func (t *TaskTimer) Wait() {
	<-t.wait
}

func Test() {
	fmt.Println(time.Now())
}

func main() {
	tt := CreateTaskTimer("2014-06-13 17:14:00", 1, Test)
	tt.Start()
	time.Sleep(10 * time.Second)
	tt.Stop()
	time.Sleep(10 * time.Second)
}
