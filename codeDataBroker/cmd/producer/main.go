package main

import (
	"flag"
	"fmt"
	"github.com/Shopify/sarama"
	"myGitCode/codeDataBroker/httptransport"
	"myGitCode/codeDataBroker/kafka"
	log "myGitCode/mylog"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

var (
	listenAddrApi string

	// kafka
	kafkaBrokerUrl string
	kafkaVerbose   bool
	kafkaClientId  string
	kafkaTopic     string
)

func main() {
	flag.StringVar(&listenAddrApi, "listen-address", "0.0.0.0:9000", "Listen address for api")
	flag.StringVar(&kafkaBrokerUrl, "kafka-brokers", "localhost:19092,localhost:29092,localhost:39092", "Kafka brokers in comma separated value")
	flag.BoolVar(&kafkaVerbose, "kafka-verbose", true, "Kafka verbose logging")
	flag.StringVar(&kafkaClientId, "kafka-client-id", "my-kafka-client", "Kafka client id to connect")
	flag.StringVar(&kafkaTopic, "kafka-topic", "foo", "Kafka topic to push")

	flag.Parse()

	log.Debug(fmt.Sprintf("Equivalent command is: ./producer -listen-address %v -kafka-brokers %v -kafka-verbose %v -kafka-client-id %v -kafka-topic %v",listenAddrApi, kafkaBrokerUrl, kafkaVerbose, kafkaClientId, kafkaTopic))

	// connect to kafka
	kafkaProducer, err := kafka.ConfigureProducer(strings.Split(kafkaBrokerUrl, ","), kafkaClientId)
	if err != nil {
		log.Error("unable to configure kafka. " + err.Error())
		return
	}
	defer kafkaProducer.Close()

	var errChan = make(chan error, 1)

	go func() {
		log.Info(fmt.Sprintf("starting server at %s", listenAddrApi))
		errChan <- startAndServeProducerApi(listenAddrApi, kafkaProducer, kafkaTopic)
	}()

	var signalChan = make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	select {
	case <-signalChan:
		log.Info("got an interrupt, exiting...")
	case err := <-errChan:
		if err != nil {
			log.Error(err.Error() + "error while running api, exiting...")
		}
	}
}

func startAndServeProducerApi(listenAddrApi string, kafkaProducer sarama.SyncProducer, topic string ) error{

	err := httptransport.ServeProducerApi(listenAddrApi, kafkaProducer, topic)
	if err != nil{
		return err
	}
	return nil
}
