package confluent

import (
	"ClientWorkerService/pkg/helper"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"log"
)

type ConfluentKafkaGo struct {
}

func (k *ConfluentKafkaGo) Produce(key *[]byte, value *[]byte, topic string) (err error) {

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": helper.ResolvePath("KAFKA_HOST", "KAFKA_PORT")})
	if err != nil {
		log.Fatal("ConfluentKafka", "Produce Connection Failed: ", err.Error())
	}
	pErr := p.Produce(&kafka.Message{
		Key:            *key,
		TopicPartition: kafka.TopicPartition{Topic: &topic},
		Value:          *value,
	}, nil)
	if pErr != nil {
		return pErr
	}
	log.Print("ConfluentKafka", "Producer", topic, key)
	p.Flush(15 * 1000)
	return nil
}

func (k *ConfluentKafkaGo) Consume(topic string, groupId string, callback func(topic string, data []byte) error) {

	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": helper.ResolvePath("KAFKA_HOST", "KAFKA_PORT"),
		"group.id":          groupId,
		"auto.offset.reset": "smallest"})
	if err != nil {
		log.Fatal("ConfluentKafka", "Consumer Connection Failed: ", err.Error())
	}

	var run = true
	for run == true {
		ev := consumer.Poll(0)
		switch e := ev.(type) {
		case *kafka.Message:
			callErr := callback(topic, e.Value)
			log.Print("ConfluentKafka", "Consumer", topic, groupId)
			if callErr == nil {
				go func() {
					offsets, err := consumer.Commit()
					if err != nil {
						log.Fatal("ConfluentKafka",
							"Consumer", "%% Commit failed %v\n", offsets, err.Error())
					}
				}()
			}

		case kafka.PartitionEOF:
			log.Fatal("ConfluentKafka",
				"Consumer", "%% PartitionEOF %v\n", e, err.Error())
		case kafka.Error:
			log.Fatal("ConfluentKafka",
				"Consumer", "%% Kafka Error: %v\n", e, err.Error())
			run = false
		default:
		}
	}
}
