package confluent

import (
	"ClientWorkerService/pkg/logger"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"os"
)

type ConfluentKafkaGo struct {
	Log logger.ILog
}

func (k *ConfluentKafkaGo) Produce(key *[]byte, value *[]byte, topic string) (err error) {

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": os.Getenv("KAFKA_BROKER")})
	if err != nil {
		k.Log.SendPanicLog("ConfluentKafka", "Produce Connection Failed: ", err.Error())
		panic(err)
	}
	pErr := p.Produce(&kafka.Message{
		Key:            *key,
		TopicPartition: kafka.TopicPartition{Topic: &topic},
		Value:          *value,
	}, nil)
	if pErr != nil {
		return pErr
	}
	k.Log.SendInfoLog("ConfluentKafka", "Producer", topic, key)
	p.Flush(15 * 1000)
	return nil
}

func (k *ConfluentKafkaGo) Consume(topic string, groupId string, callback func(topic string, data []byte) error) {

	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":    os.Getenv("KAFKA_BROKER"),
		"group.id":             groupId,
		"auto.offset.reset":    "smallest"})
	if err != nil{
		k.Log.SendPanicLog("ConfluentKafka", "Consumer Connection Failed: ", err.Error())
		panic(err)
	}

	var run = true
	for run == true {
		ev := consumer.Poll(0)
		switch e := ev.(type) {
		case *kafka.Message:
			callErr := callback(topic, e.Value)
			k.Log.SendInfoLog("ConfluentKafka", "Consumer", topic, groupId)
			if callErr == nil{
				go func() {
					offsets, err := consumer.Commit()
					if err != nil{
						k.Log.SendErrorfLog("ConfluentKafka",
							"Consumer","%% Commit failed %v\n", offsets, err.Error())
					}
				}()
			}

		case kafka.PartitionEOF:
			k.Log.SendErrorfLog("ConfluentKafka",
				"Consumer","%% PartitionEOF %v\n", e, err.Error())
		case kafka.Error:
			k.Log.SendErrorfLog("ConfluentKafka",
				"Consumer","%% Kafka Error: %v\n", e, err.Error())
			run = false
		default:
		}
	}
}

