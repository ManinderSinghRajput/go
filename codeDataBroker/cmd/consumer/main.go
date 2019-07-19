package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/namsral/flag"
	"myGitCode/codeDataBroker/kafka"
	log "myGitCode/mylog"
)

var (
	// kafka
	kafkaBrokerUrl     string
	kafkaVerbose       bool
	kafkaTopic         string
	kafkaConsumerGroup string
	kafkaClientId      string
	maxParallelMsg  *int
)

func main() {
	flag.StringVar(&kafkaBrokerUrl, "kafka-brokers", "localhost:19092,localhost:29092,localhost:39092", "Kafka brokers in comma separated value")
	flag.BoolVar(&kafkaVerbose, "kafka-verbose", true, "Kafka verbose logging")
	flag.StringVar(&kafkaTopic, "kafka-topic", "foo", "Kafka topic. Only one topic per worker.")
	flag.StringVar(&kafkaConsumerGroup, "kafka-consumer-group", "consumer-group", "Kafka consumer group")
	flag.StringVar(&kafkaClientId, "kafka-client-id", "my-client-id", "Kafka client id")
	maxParallelMsg = flag.Int("maxRoutines", 100000, "Number max of messages read from kafka that can be treated in parallel.")
	flag.Parse()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	brokers := strings.Split(kafkaBrokerUrl, ",")
	topics := strings.Split(kafkaTopic, ",")

	consumer := kafka.ConfigureConsumer(brokers, topics, "testGroup")

	defer consumer.Close()
	messages := kafka.Pull(consumer, *maxParallelMsg)

	for {
		select{
		case msg := <-messages:
			log.Info(fmt.Sprintf("Topic: %v, Partition: %v, Offset: %v, Message: %s", msg.Topic, msg.Partition, msg.Offset, msg.Value))
		case <-sigChan:
			log.Info(fmt.Sprintf("Stoping in progress...."))
		}

	}





}