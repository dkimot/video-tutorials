package pkg

type MessageStore interface {
	Close() error
	Write(streamName string, msg Message, expVersion string) error
}
