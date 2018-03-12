package CoroutinePool

// Worker -- the worker that executes the job
type Worker struct {
	WorkerPool chan chan Job
	JobChannel chan Job
	quit       chan bool
}

// NewWorker -- new a worker
func NewWorker(WorkerPool chan chan Job) *Worker {
	return &Worker{
		WorkerPool: WorkerPool,
		JobChannel: make(chan Job),
		quit:       make(chan bool),
	}
}

// Start -- Start the run loop for the worker
func (w Worker) Start() {
	go func() {
		for {
			w.WorkerPool <- w.JobChannel

			select {
			case job := <-w.JobChannel:
				if job.Payload != nil {
					// executes Call()
					job.Payload.Call()
				}
			case <-w.quit:
				return
			}
		}
	}()
}

// Stop -- stop the run loop for the worker
func (w Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}
