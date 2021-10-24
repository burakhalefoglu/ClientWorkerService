package Ikafka

type IKafka interface {
	Produce(key *[]byte, value *[]byte, topic string) (err error)
	Consume(topic string, groupId string, callback func(topic string, data []byte) error)
}
