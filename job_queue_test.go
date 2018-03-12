package CoroutinePool

import (
	"fmt"
	"testing"
	"time"
)

type PayloadType1 struct {
	Data string
}

func (d *PayloadType1) Call() {
	fmt.Println("type1:", d.Data)
}

type PayloadType2 struct {
	UserName string
}

func (d *PayloadType2) Call() {
	fmt.Println("Type2:", d.UserName)
}

func TestJobQueue(t *testing.T) {
	d := NewDispatcher(100, 10)
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
