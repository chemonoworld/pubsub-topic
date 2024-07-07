package pubsub

import (
	"context"
	"fmt"
	"github.com/chemonoworld/pubsub-topic/message"
	"sync"
)

type Subscriber struct {
	ctx      context.Context
	name     string
	subMsgCh *message.SubMessageChan
}

func NewSubscriber(name, topic string, ctx context.Context) *Subscriber {
	return &Subscriber{
		ctx:      ctx,
		name:     name,
		subMsgCh: message.NewSubMessageChan(topic),
	}
}

func (s *Subscriber) Subscribe(p *Publisher) {
	p.Subscribe(s.subMsgCh)
}

func (s *Subscriber) Update(wg *sync.WaitGroup) {
	for {
		select {
		case msg := <-s.subMsgCh.GetMsgCh():
			fmt.Printf("%s got Message:%#v in %s topic\n", s.name, msg.GetMsg(), s.subMsgCh.GetTopic())
		case <-s.ctx.Done():
			wg.Done()
			return
		}
	}
}
