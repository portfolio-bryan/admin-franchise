package concurrency

import (
	"sync"
)

func Fanin(
	done <-chan interface{},
	channels ...<-chan ChannelData,
) <-chan ChannelData {
	var wg sync.WaitGroup
	multiplexedStream := make(chan ChannelData)

	multiplex := func(stream <-chan ChannelData) {
		defer wg.Done()
		for v := range stream {
			select {
			case <-done:
				return
			case multiplexedStream <- v:
			}
		}
	}

	wg.Add(len(channels))

	for _, c := range channels {
		go multiplex(c)
	}

	go func() {
		wg.Wait()
		close(multiplexedStream)
	}()

	return multiplexedStream
}
