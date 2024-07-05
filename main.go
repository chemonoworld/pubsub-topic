package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(4)
	publisher := NewPublisher(ctx)
	subscriber1 := NewSubscriber("subscriber1", "topicA", ctx)
	subscriber2 := NewSubscriber("subscriber2", "topicA", ctx)
	subscriber3 := NewSubscriber("subscriber3", "topicC", ctx)

	go publisher.Update()
	go subscriber1.Update()
	go subscriber2.Update()
	go subscriber3.Update()

	subscriber1.Subscribe(publisher)
	subscriber2.Subscribe(publisher)
	subscriber3.Subscribe(publisher)

	go func() {
		tick := time.Tick(time.Second * 2)
		for {
			select {
			case <-tick:
				publisher.Publish(&PubMessage{topic: "topicC", msg: "Hello, World A!"})
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
