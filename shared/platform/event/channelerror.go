package event

import (
	"log"

	"github.com/bperezgo/admin_franchise/shared/domain"
)

type ChannelError struct {
	internalChan chan domain.WrapError
}

func NewChannelError() ChannelError {
	internalChan := make(chan domain.WrapError)

	ce := ChannelError{
		internalChan: internalChan,
	}
	ce.listen()
	return ce
}

func (c ChannelError) Publish(err error) {
	c.internalChan <- domain.WrapError{
		Err: err,
	}
}

func (c ChannelError) listen() {
	// TODO: Handle all the errors continuously
	go func() {
		for {
			select {
			case err := <-c.internalChan:
				if err.Err != nil {
					log.Println(err)
				}
			}
		}
	}()
}
