package CoroutinePool

// Dispatcher 分配器
type Dispatcher struct {
	workerPool chan chan Job
	maxWorkers int
	workers    []*Worker
	jobQueue   chan Job
}

// NewDispatcher -- maxQueue 一级channel缓冲最大值， maxWorkers 二级channel缓冲最大值
func NewDispatcher(maxQueue, maxWorkers int) *Dispatcher {
	return &Dispatcher{
		workerPool: make(chan chan Job, maxWorkers),
		maxWorkers: maxWorkers,
		workers:    []*Worker{},
		jobQueue:   make(chan Job, maxQueue),
	}
}

// Run -- run workers
func (d *Dispatcher) Run() {
	if len(d.workers) != d.maxWorkers {
		count := d.maxWorkers - len(d.workers)
		d.createWorkers(count)
	}
	for _, worke := range d.workers {
		worke.Start()
	}
	go d.dispatcher()
}

// PushPayload -- Push a payload
func (d *Dispatcher) PushPayload(payload Payload) {
	d.jobQueue <- Job{
		Payload: payload,
	}
}

func (d *Dispatcher) prepareWorkes() {
	if len(d.workers) != d.maxWorkers {
		count := d.maxWorkers - len(d.workers)
		d.createWorkers(count)
	}
}

func (d *Dispatcher) createWorkers(count int) {
	if d.workers == nil {
		d.workers = []*Worker{}
	}
	for i := 0; i < count; i++ {
		worker := NewWorker(d.workerPool)
		d.workers = append(d.workers, worker)
	}
}

// Stop 停止所有worker
func (d *Dispatcher) Stop() {
	for _, worker := range d.workers {
		worker.Stop()
	}
}

func (d *Dispatcher) dispatcher() {
	for {
		select {
		case job := <-d.jobQueue:
			go func(job Job) {
				jobChannel := <-d.workerPool
				jobChannel <- job
			}(job)
		}
	}
}
