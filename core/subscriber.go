package core

import (
	webpush "github.com/SherClockHolmes/webpush-go"
	"github.com/louisevanderlith/husk"
)

type Subscriber struct {
	Subscription webpush.Subscription
	State        string
}

func (s Subscriber) Valid() error {
	return husk.ValidateStruct(s)
}

func AddSubscriber(subsc webpush.Subscription) error {
	item := Subscriber{
		Subscription: subsc,
		State:        "ACTIVE",
	}

	_, err := ctx.Subscribers.Create(item)

	if err != nil {
		return err
	}

	return ctx.Subscribers.Save()
}
