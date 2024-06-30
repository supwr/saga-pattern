package kafka

import (
	confluentkafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
)

type Config struct {
	Brokers string
	Topic   string
}

type Message struct {
	Data []byte
	Key  []byte
}

type Producer struct {
	topic    string
	producer *confluentkafka.Producer
}

func NewProducer(cfg Config) (*Producer, error) {
	log.Printf("connecting to %s", cfg.Brokers)
	p, err := confluentkafka.NewProducer(&confluentkafka.ConfigMap{
		"bootstrap.servers": cfg.Brokers,
		"acks":              "all"})
	if err != nil {
		return nil, err
	}

	return &Producer{
		topic:    cfg.Topic,
		producer: p,
	}, nil
}

func (p *Producer) SendMessage(message Message) error {
	log.Printf("sending message to %s %s", p.topic, string(message.Data))
	return p.producer.Produce(&confluentkafka.Message{
		TopicPartition: confluentkafka.TopicPartition{Topic: &p.topic, Partition: confluentkafka.PartitionAny},
		Value:          message.Data,
		Key:            message.Key,
	}, nil)
}
