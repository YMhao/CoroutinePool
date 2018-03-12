package CoroutinePool

// Payload -- An Payload interface
// worker 将会从 Job 中获取 payload, 执行该payload的call方法
type Payload interface {
	Call()
}

// Job -- the job to be run
type Job struct {
	Payload Payload
}

// JobQueue -- A buffered channel that we can send work requests on.
// 工作任务发布到这里， worker将会从这里读取工作任务
var JobQueue = make(chan Job)
