package workerpool

type Worker interface {
	Working(taskCh chan Task)
}
