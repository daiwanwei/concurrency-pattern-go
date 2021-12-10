package pubsub

type Publisher interface {
	Start()
	SubscribeChan() chan<- Subscriber
	UnsubscribeChan() chan<- Subscriber
	PublishChan() chan<- interface{}
	Close()
}
