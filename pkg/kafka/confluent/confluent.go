package confluent

import (
	"fmt"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"log"
	"os"
)



func Produce(key *[]byte, value *[]byte, topic string) (err error) {

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "0.0.0.0:19091"})
	if err != nil {
		panic(err)
	}
		p.Produce(&kafka.Message{
			Key: *key,
			TopicPartition: kafka.TopicPartition{Topic: &topic},
			Value:          *value,
		}, nil)
	// Wait for message deliveries before shutting down
	p.Flush(15 * 1000)
	return nil
}

func Consume(topic string, groupId string, callback func(topic string, data []byte) error) {

	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":    "192.168.1.1:9092",
		"group.id":             groupId,
		"auto.offset.reset":    "smallest"})
	if err != nil{
		panic(err)
	}

	var run = true
	for run == true {
		ev := consumer.Poll(0)
		switch e := ev.(type) {
		case *kafka.Message:
			callErr := callback(topic, e.Value)
			if callErr == nil{
				go func() {
					offsets, err := consumer.Commit()
					if err != nil{
						log.Printf("%% Reached %v\n", err)
					}
					log.Printf("%% Reached %v\n", offsets)
				}()
			}

		case kafka.PartitionEOF:
			log.Printf("%% Reached %v\n", e)
		case kafka.Error:
			fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
			run = false
		default:
			log.Printf("Ignored %v\n", e)
		}
	}
}

