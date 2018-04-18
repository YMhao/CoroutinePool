# 概述

Coroutine Pool

这是一个协程池， 用户处理大负载情况，降低延时

主要基于以下两点考虑：
1、要有一个缓冲的队列。
2、控制产生的协程的数量。

# 如何处理大负载，降低延时

简单地使用goroutine能满足中小型负载， 代码如下

```golang

   go func(){
       // to do
   }()

```

但是这样我们没法控制产生的go routine的数量，当请求量很大的时候，这样的代码很快就崩溃了。  

改进（但问题依然存在，不能很好地控制任务并行的数量）
```golang

var JobQueue = make(chan Payload, MAX_JOB_QUEUE)

func StartProcessor() {
    for {
        select {
        case job := <-JobQueue:
            //to do
        }
    }
}
```
这样做有一个存放任务的队列，但是不能控制任务队列的并发量。   

那么如何解决呢？

这里使用了二级channel的方法来解决这个问题。
第一级：有一个存放任务的缓冲队列。
第二级：控制任务队列的并发数（对worker的数量进行控制）。

#  该协程池特点

1、使用简单， 扩展方便 .  
2、负载的结构体可以进行自定义， 实现Call()方法即可。

# 接口说明

接口有4个：
1、新建二级缓存。
2、开始工作。
2、推送负载。
3、停止工作。

新建二级缓存(一个分配器):
```golang
// NewDispatcher -- maxQueue 任务队列的缓存大小， maxWorkers 工作携程的数量（控制并行数）
func NewDispatcher(maxQueue, maxWorkers int) *Dispatcher
```
开始工作：
```golang
func (d *Dispatcher) Run()
```

推送负载:
```golang
// Payload -- Payload interface
// worker 将会从 Job 中获取 payload, 执行该payload的call方法
type Payload interface {
	Call()
}
// PushPayload -- Push a payload
func (d *Dispatcher) PushPayload(payload Payload)

```

停止工作:
```golang
func (d *Dispatcher) Stop()
```
