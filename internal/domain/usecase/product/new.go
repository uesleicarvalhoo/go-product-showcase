package product

type UseCase struct {
	eventTopic string
	repository Repository
	broker     Broker
}

func New(r Repository, b Broker, eventTopic string) UseCase {
	return UseCase{
		eventTopic: eventTopic,
		broker:     b,
		repository: r,
	}
}
