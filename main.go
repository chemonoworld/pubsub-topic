package main

import (
	"context"
	"fmt"
	"github.com/chemonoworld/pubsub-topic/message"
	"github.com/chemonoworld/pubsub-topic/pubsub"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(4)
	publisher := pubsub.NewPublisher(ctx)
	subscriber1 := pubsub.NewSubscriber("subscriber1", "topicA", ctx)
	subscriber2 := pubsub.NewSubscriber("subscriber2", "topicA", ctx)
	subscriber3 := pubsub.NewSubscriber("subscriber3", "topicC", ctx)
	subscriber4 := pubsub.NewSubscriber("subscriber4", "topicC", ctx)
	subscriber5 := pubsub.NewSubscriber("subscriber5", "topicC", ctx)
	subscriber6 := pubsub.NewSubscriber("subscriber6", "topicC", ctx)
	subscriber7 := pubsub.NewSubscriber("subscriber7", "topicC", ctx)

	go publisher.Update(&wg)
	go subscriber1.Update(&wg)
	go subscriber2.Update(&wg)
	go subscriber3.Update(&wg)
	go subscriber4.Update(&wg)
	go subscriber5.Update(&wg)
	go subscriber6.Update(&wg)
	go subscriber7.Update(&wg)

	subscriber1.Subscribe(publisher)
	subscriber2.Subscribe(publisher)
	subscriber3.Subscribe(publisher)
	subscriber4.Subscribe(publisher)
	subscriber5.Subscribe(publisher)
	subscriber6.Subscribe(publisher)
	subscriber7.Subscribe(publisher)

	go func() {
		tick := time.Tick(time.Second * 2)
		for {
			select {
			case <-tick:
				//publisher.Publish(message.NewPubMessage("topicA", "Hello, topicA"))
				publisher.Publish(message.NewPubMessage("topicC", "Hello, topicC"))
			case <-ctx.Done():
				wg.Done()
				return
			}
		}
	}()

	fmt.Scanln()
	cancel()

	wg.Wait()
}
