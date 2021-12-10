package main

import (
	"fmt"
	"pattern-golang/concurrency/pubsub/pubsub"
	"time"
)

type MovieNews struct {
	MovieName   string
	Description string
}

type Netflix struct {
	subscribers []pubsub.Subscriber
	subCh       chan pubsub.Subscriber
	unsubCh     chan pubsub.Subscriber
	msgCh       chan interface{}
	stopCh      chan struct{}
}

func NewNetflix() *Netflix {
	return &Netflix{
		subscribers: []pubsub.Subscriber{},
		subCh:       make(chan pubsub.Subscriber, 20),
		unsubCh:     make(chan pubsub.Subscriber, 20),
		msgCh:       make(chan interface{}, 20),
		stopCh:      make(chan struct{}, 20),
	}
}

func (n *Netflix) Start() {
	for {
		select {
		case msg := <-n.msgCh:
			for _, s := range n.subscribers {
				s.Notify(msg)
			}
		case add := <-n.subCh:
			n.subscribers = append(n.subscribers, add)
		case remove := <-n.unsubCh:
			for i, s := range n.subscribers {
				if s == remove {
					n.subscribers = append(n.subscribers[:i], n.subscribers[i+1:]...)
					remove.Close()
					break
				}
			}
		case <-n.stopCh:
			for _, s := range n.subscribers {
				s.Close()
			}
			close(n.subCh)
			close(n.msgCh)
			close(n.unsubCh)
			return
		}
	}
}

func (n *Netflix) SubscribeChan() chan<- pubsub.Subscriber {
	return n.subCh
}

func (n *Netflix) UnsubscribeChan() chan<- pubsub.Subscriber {
	return n.unsubCh
}

func (n *Netflix) PublishChan() chan<- interface{} {
	return n.msgCh
}

func (n *Netflix) Close() {
	close(n.stopCh)
}

type Audience struct {
	id    int
	name  string
	email string
	msgCh chan interface{}
}

func NewAudience(id int, name, email string) *Audience {
	msgCh := make(chan interface{})
	a := &Audience{
		id:    id,
		name:  name,
		email: email,
		msgCh: msgCh,
	}
	go func() {
		for m := range a.msgCh {
			fmt.Printf("audience(id:%d,name:%s:%s)\n", a.id, a.name, m)
			fmt.Printf("send email(%s):%s\n", a.email, m)
		}
	}()
	return a
}

func (a *Audience) Notify(msg interface{}) (err error) {
	defer func() {
		if rec := recover(); rec != nil {
			err = fmt.Errorf("%#v", rec)
		}
	}()
	select {
	case a.msgCh <- msg:
	case <-time.After(time.Second * 5):
		err = fmt.Errorf("timeout\n")
		return
	}
	return
}

func (a *Audience) Close() {
	close(a.msgCh)
}
