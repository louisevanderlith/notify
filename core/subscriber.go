package core

import "github.com/louisevanderlith/husk"

type Subscriber struct {
}

func (s Subscriber) Valid() (bool, error) {
	return husk.ValidateStruct(&s)
}
