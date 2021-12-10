package workerpool

type Dispatcher interface {
	LaunchWorker(worker Worker)
	ReceiveTask(task Task)
	Stop()
}
