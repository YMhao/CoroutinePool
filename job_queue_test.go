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
	d := NewDispatcher(10)
	d.Run()

	go func() {
		for {
			payload := &PayloadType1{
				Data: "abc",
			}
			JobQueue <- Job{
				Payload: payload,
			}
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for {
			payload := &PayloadType2{
				UserName: "hao",
			}
			JobQueue <- Job{
				Payload: payload,
			}
			time.Sleep(2 * time.Second)
		}
	}()

	time.Sleep(5 * time.Second)
}
