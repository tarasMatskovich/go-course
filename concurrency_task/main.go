//////////////////////////////////////////////////////////////////////
//
// Given is a producer-consumer scenario, where a producer reads in
// tweets from a mockstream and a consumer is processing the
// data. Your task is to change the code so that the producer as well
// as the consumer can run concurrently
//

package main

import (
	"concurrency_task/mockstream"
	"fmt"
	"sync"
	"time"
)

func producer(stream mockstream.Stream, link chan<- *mockstream.Tweet, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		tweet, err := stream.Next()
		if err == mockstream.ErrEOF {
			link <- nil
			break
		}

		link <- tweet
	}
}

func consumer(link <-chan *mockstream.Tweet, wg *sync.WaitGroup) {
	i := 0
	for t := range link {
		i++
		if t == nil {
			break
		}
		wg.Add(1)
		go func(tw *mockstream.Tweet, wg *sync.WaitGroup, i int) {
			defer wg.Done()
			if tw.IsTalkingAboutGo() {
				fmt.Println(tw.Username, "\ttweets about golang")
			} else {
				fmt.Println(tw.Username, "\tdoes not tweet about golang")
			}
		}(t, wg, i)
	}
	wg.Wait()
}

func main() {
	start := time.Now()
	stream := mockstream.GetMockStream()
	link := make(chan *mockstream.Tweet)

	wp := &sync.WaitGroup{}
	wc := &sync.WaitGroup{}

	// Producer
	wp.Add(1)
	go producer(stream, link, wp)

	// Consumer
	consumer(link, wc)

	wp.Wait()
	close(link)
	wc.Wait()

	fmt.Printf("Process took %s\n", time.Since(start))
}
