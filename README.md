# 概述

Coroutine Pool

这是一个协程池， 用户处理大负载情况，降低延时

# 如何处理大负载，降低延时

简单地使用goroutine能满足中小型负载， 代码如下

```golang

   go func(){
       // to do
   }()

```

但是这样我们没法控制产生的go routine的数量， 当收到每分钟1百万个请求时，这样的代码很快就崩溃了。

如果用创建一个缓冲的channel来处理，也只是把问题推迟了而已。

那么如何解决呢？

这里使用了二级channel的方法来解决这个问题。

#  该协程池特点

使用简单， 扩展方便 

负载的结构体可以进行自定义， 实现Call()方法即可

如果是api接口，那么负载可以是各个api的请求

举个例子， 获取用户信息，匹配文本
```golang
type GetUserInfoRequest struct{
    SessionId string
    C *gin.Context
}

func (g *GetUserInfoRequest) Call() {
    // to do 
}

type MatchWordRequest struct {
    SessionId string
    Keyword   string
    C         *gin.Context
}

func (m *MatchUserRequest) Call() {
    // to do 
}

```


# 如何使用(三部曲)

例子请看 example/demo_hello_world

1. 定义负载

负载接口定义如下:
```golang
type Payload interface {
	Call()
}
```
例子：

```golang
// PayLoadType1 -- PayLoad Type 1
type PayLoadType1 struct {
	Data string
}

// Call -- method Call
func (d *PayLoadType1) Call() {
	fmt.Println("Type1:", "Data =", d.Data)
}
```

2. 设置池大小, 启动
```golang
func main() {
    ...
	d := CoroutinePool.NewDispatcher(10)
    d.Run()
    ...
}
```

3. 往 JobQueue 里推送工作任务即可

```golang
payload := &PayloadType1{
    Data: "abc",
}
JobQueue <- Job{
    Payload: payload,
}
```