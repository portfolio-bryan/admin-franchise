package event

import (
	"log"

	"github.com/bperezgo/admin_franchise/shared/domain"
)

type ChannelError struct {
	internalChan chan domain.Error
}

func NewChannelError() ChannelError {
	internalChan := make(chan domain.Error)

	ce := ChannelError{
		internalChan: internalChan,
	}
	ce.listen()
	return ce
}

func (c ChannelError) Publish(err error) {
	c.internalChan <- domain.Error{
		Err: err,
	}
}

func (c ChannelError) listen() {
	// TODO: Handle all the errors continuously
	go func() {
		for err := range c.internalChan {
			if err.Err != nil {
				log.Println(err)
			}
		}
	}()
}
