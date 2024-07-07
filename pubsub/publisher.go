package pubsub

import (
	"context"
	"github.com/chemonoworld/pubsub-topic/message"
	"sync"
)

type Publisher struct {
	ctx                context.Context
	subscribe          chan *message.SubMessageChan
	pubMsgCh           chan *message.PubMessage
	topicToSubscribers map[string][]chan<- *message.PubMessage
}

func NewPublisher(ctx context.Context) *Publisher {
	return &Publisher{
		ctx:                ctx,
		subscribe:          make(chan *message.SubMessageChan),
		pubMsgCh:           make(chan *message.PubMessage),
		topicToSubscribers: make(map[string][]chan<- *message.PubMessage),
	}
}

func (p *Publisher) Subscribe(subMsgCh *message.SubMessageChan) {
	p.subscribe <- subMsgCh
}

func (p *Publisher) Publish(msg *message.PubMessage) {
	p.pubMsgCh <- msg
}

func (p *Publisher) Update(wg *sync.WaitGroup) {
	for {
		select {
		case sub := <-p.subscribe:
			topic := sub.GetTopic()
			p.topicToSubscribers[topic] = append(p.topicToSubscribers[topic], sub.GetMsgCh())
		case msg := <-p.pubMsgCh:
			subscribers := p.topicToSubscribers[msg.GetTopic()]
			for _, subscriber := range subscribers {
				subscriber <- msg
			}
		case <-p.ctx.Done():
			wg.Done()
			return
		}
	}
}
