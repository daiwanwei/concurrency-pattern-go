package workerpool

type Task struct {
	Data    interface{}
	Handler TaskHandler
}

type TaskHandler func(interface{})
