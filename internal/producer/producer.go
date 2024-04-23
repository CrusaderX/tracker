package producer

import (
	"github.com/IBM/sarama"
	"log"
)

type Config struct {
	Topic   string
	Brokers []string
}

type Producer struct {
	producer sarama.AsyncProducer
	config   Config
}

func New(brokers []string) (Producer, error) {
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForLocal
	config.Producer.Return.Successes = true

	p, err := sarama.NewAsyncProducer(brokers, config)

	if err != nil {
		return Producer{}, err
	}

	return Producer{producer: p}, nil
}

func (p Producer) Produce(topic, message string) error {
	m := &sarama.ProducerMessage{Topic: p.config.Topic, Value: sarama.StringEncoder(message)}

	select {
	case p.producer.Input() <- m:
		log.Println("New Message produced")
	}
	// case timeout -> error

	return nil
}
