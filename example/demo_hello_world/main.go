package main

import (
	"fmt"
	"time"

	"github.com/YMhao/CoroutinePool"
)

// PayLoadType1 -- PayLoad Type 1
type PayLoadType1 struct {
	Data string
}

// Call -- method Call
func (d *PayLoadType1) Call() {
	fmt.Println("Type1:", "Data =", d.Data)
}

// PayLoadType2 -- PayLoad Type 1
type PayLoadType2 struct {
	UserName string
}

// Call -- method Call
func (d *PayLoadType2) Call() {
	fmt.Println("Type2:", "Hello", d.UserName)
}

func main() {
	d := CoroutinePool.NewDispatcher(10)
	d.Run()

	go func() {
		for {
			payload := &PayLoadType1{
				Data: "abc",
			}
			CoroutinePool.JobQueue <- CoroutinePool.Job{
				Payload: payload,
			}
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for {
			payload := &PayLoadType2{
				UserName: "world",
			}
			CoroutinePool.JobQueue <- CoroutinePool.Job{
				Payload: payload,
			}
			time.Sleep(2 * time.Second)
		}
	}()
	time.Sleep(5 * time.Second)
}
