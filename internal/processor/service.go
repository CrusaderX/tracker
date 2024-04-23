package processor

type Processor struct {
	Producer Producer
}

type Producer interface {
	Produce(message, topic string) error
}

func NewDefaultProcessor(producer Producer) *Processor {
	return &Processor{
		Producer: producer,
	}
}

func (p *Processor) Handle(message, topic string) error {
	return p.Producer.Produce(message, topic)
}
