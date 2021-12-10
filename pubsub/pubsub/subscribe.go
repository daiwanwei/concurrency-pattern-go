package pubsub

type Subscriber interface {
	Notify(interface{}) error
	Close()
}
