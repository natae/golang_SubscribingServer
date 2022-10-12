package subscription

type Subscription interface {
	Updates() <-chan Item
	Close() error
}

type Item struct {
	Title, Channel, GUID string
}
