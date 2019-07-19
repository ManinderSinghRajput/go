package kafka

import (
	"github.com/Shopify/sarama"
	cluster "gopkg.in/bsm/sarama-cluster.v2"
)

func Pull(consumer *cluster.Consumer, maxGoroutines int) chan *sarama.ConsumerMessage {
	am := make(chan *sarama.ConsumerMessage, maxGoroutines)

	go func() {
		for {
			t := <-consumer.Messages()

			am <- t
		}
	}()

	return am
}
