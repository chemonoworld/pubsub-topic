package main

import (
	"context"
)

type Publisher struct {
	ctx       context.Context
	subscribe chan *SubMessageChan
	publicCh  chan *PubMessage
	// 이걸 토픽별로 나눠서 바꿔보자. ㅇㅇ
	topicToSubscribers map[string][]chan<- *PubMessage
}

func NewPublisher(ctx context.Context) *Publisher {
	return &Publisher{
		ctx:                ctx,
		subscribe:          make(chan *SubMessageChan),
		publicCh:           make(chan *PubMessage),
		topicToSubscribers: make(map[string][]chan<- *PubMessage),
	}
}

func (p *Publisher) Subscribe(subMsgCh *SubMessageChan) {
	p.subscribe <- subMsgCh
}

func (p *Publisher) Publish(msg *PubMessage) {
	p.publicCh <- msg
}

func (p *Publisher) Update() {
	for {
		select {
		case sub := <-p.subscribe:
			topic := sub.getTopic()
			p.topicToSubscribers[topic] = append(p.topicToSubscribers[topic], sub.getMsgCh())
		case msg := <-p.publicCh:
			subscribers := p.topicToSubscribers[msg.getTopic()]
			for _, subscriber := range subscribers {
				subscriber <- msg
			}
		case <-p.ctx.Done():
			wg.Done()
			return
		}
	}
}
