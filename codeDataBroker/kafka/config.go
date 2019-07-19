package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"gopkg.in/bsm/sarama-cluster.v2"
	"log"
	"myGitCode/mylog"
	"os"
	"strings"
)

var ready = make(chan bool, 0)

/*func Configure(kafkaBrokerUrls []string, clientId string, topic string) (w *kafka.Writer, err error) {
	dialer := &kafka.Dialer{
		Timeout:  10 * time.Second,
		ClientID: clientId,
	}

	config := kafka.WriterConfig{
		Brokers:          kafkaBrokerUrls,
		Topic:            topic,
		Balancer:         &kafka.LeastBytes{},
		Dialer:           dialer,
		WriteTimeout:     10 * time.Second,
		ReadTimeout:      10 * time.Second,
		CompressionCodec: snappy.NewCompressionCodec(),
	}
	w = kafka.NewWriter(config)
	writer = w
	return w, nil
}*/

func ConfigureProducer(kafkaBrokerUrls []string, clientId string)(sarama.SyncProducer, error) {
	// setup sarama log to stdout
	sarama.Logger = log.New(os.Stdout, "", log.Ltime)

	// producer config
	config := sarama.NewConfig()
	config.Producer.Retry.Max = 5
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	config.ClientID = clientId

	// async producer
	//prd, err := sarama.NewAsyncProducer([]string{kafkaConn}, config)

	// sync producer
	prd, err := sarama.NewSyncProducer(kafkaBrokerUrls, config)

	return prd, err
}

func ConfigureConsumer(brokers []string, topics []string, groupID string) *cluster.Consumer{
	config := cluster.NewConfig()
	config.Consumer.Return.Errors = true
	config.Group.Return.Notifications = true
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	consumer, err := cluster.NewConsumer(brokers, "testGroup", topics, config)
	if err != nil {
		mylog.Error("Cannot connect to kafka." + err.Error())
	}

	go func(){
		for err := range consumer.Errors(){
			mylog.Error("Kafka error" + err.Error())
			return
		}
	}()
	go func() {
		for note := range consumer.Notifications() {
			mylog.Debug(fmt.Sprintf("Rebalanced"))
			for topic, partitions := range note.Claimed {
				for partition := range partitions {
					mylog.Debug(fmt.Sprintf("Claimed partition [%s/%d]", topic, partition))
				}
			}
			for topic, partitions := range note.Released {
				for partition := range partitions {
					mylog.Debug(fmt.Sprintf("Released partition [%s/%d]", topic, partition))
				}
			}
		}
	}()

	mylog.Debug(fmt.Sprintf("Connected to kafka broker for %s with groupId %s and topic %s ", groupID, groupID, strings.Join(topics[:], ",")))

	return consumer
}