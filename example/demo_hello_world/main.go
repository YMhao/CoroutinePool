package main

import (
	"fmt"
	"time"

	"github.com/YMhao/CoroutinePool"
)

// PayloadType1 -- PayLoad Type 1
type PayloadType1 struct {
	Data string
}

// Call -- method Call
func (d *PayloadType1) Call() {
	fmt.Println("Type1:", "Data =", d.Data)
}

// PayloadType2 -- PayLoad Type 1
type PayloadType2 struct {
	UserName string
}

// Call -- method Call
func (d *PayloadType2) Call() {
	fmt.Println("Type2:", "Hello", d.UserName)
}

func main() {
	d := CoroutinePool.NewDispatcher(100, 10)
	d.Run()

	go func() {
		for {
			payload := &PayloadType1{
				Data: "abc",
			}
			d.PushPayload(payload)
		}
	}()

	go func() {
		for {
			payload := &PayloadType2{
				UserName: "hao",
			}
			d.PushPayload(payload)
		}
	}()

	time.Sleep(1 * time.Second)
	d.Stop()
	time.Sleep(1 * time.Second)
}
