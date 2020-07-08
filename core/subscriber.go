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
	return husk.ValidateStruct(&s)
}

func AddSubscriber(subsc webpush.Subscription) error {
	defer ctx.Subscribers.Save()

	item := Subscriber{
		Subscription: subsc,
		State:        "ACTIVE",
	}
	set := ctx.Subscribers.Create(item)

	if set.Error != nil {
		return set.Error
	}

	return nil
}
