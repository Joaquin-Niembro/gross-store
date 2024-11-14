package utils

import "sync"

func FanIn[k any](channels ...<-chan k) <-chan k {
	wg := sync.WaitGroup{}
	wg.Add(len(channels))
	output := make(chan k)

	for _, channel := range channels {
		go func(ch <-chan k) {
			defer wg.Done()
			for i := range ch {
				output <- i
			}
		}(channel)
	}

	go func() {
		wg.Wait()
		close(output)
	}()
	return output
}
