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
