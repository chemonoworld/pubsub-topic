package main

import (
	"context"
	"fmt"
)

type Subscriber struct {
	ctx      context.Context
	name     string
	subMsgCh *SubMessageChan
}

func NewSubscriber(name, topic string, ctx context.Context) *Subscriber {
	return &Subscriber{
		ctx:      ctx,
		name:     name,
		subMsgCh: &SubMessageChan{topic: topic, msgCh: make(chan *PubMessage)},
	}
}

func (s *Subscriber) Subscribe(p *Publisher) {
	p.Subscribe(s.subMsgCh)
}

func (s *Subscriber) Update() {
	for {
		select {
		case msg := <-s.subMsgCh.msgCh:
			fmt.Printf("%s got Message:%s in %s topic\n", s.name, msg, s.subMsgCh.topic)
		case <-s.ctx.Done():
			wg.Done()
			return
		}
	}
}
