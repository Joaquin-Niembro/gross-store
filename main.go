package main

import (
	"crypto/sha256"
	"fmt"
	"gross-store/models"
	"gross-store/utils"
	"sync"
)

// TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.
func generateOrderHash(ch <-chan int) <-chan string {
	output := make(chan string)
	wg := sync.WaitGroup{}
	for order := range ch {
		wg.Add(1)
		go func(o int) {
			h := sha256.New()
			h.Write([]byte(string(rune(o))))
			bs := h.Sum(nil)
			output <- fmt.Sprintf("%x", bs)
			wg.Done()
		}(order)
	}
	go func() {
		wg.Wait()
		close(output)
	}()
	return output
}

func simulateBuys(wg *sync.WaitGroup, store *models.Store) (<-chan int, <-chan int) {
	shortsCH := make(chan int, 20000)
	jacketsCH := make(chan int, 20000)
	for i := 0; i < 40000; i++ {
		wg.Add(1)
		go func(n int) {
			switch i % 2 {
			case 0:
				shortsCH <- store.RestShorts(n)
				wg.Done()
			case 1:
				jacketsCH <- store.RestJackets(n)
				wg.Done()
			}
		}(i)
	}
	go func() {
		wg.Wait()
		close(shortsCH)
		close(jacketsCH)
	}()
	return shortsCH, jacketsCH
}

func main() {
	store := &models.Store{
		Shorts:  20000,
		Jackets: 20000,
		Mu:      sync.Mutex{},
	}
	wg := &sync.WaitGroup{}
	shortsCH, jacketsCH := simulateBuys(wg, store)
	channels := make([]<-chan string, 2)
	channels[0], channels[1] = generateOrderHash(shortsCH), generateOrderHash(jacketsCH)
	results := utils.FanIn(channels...)
	for r := range results {
		fmt.Println("order: ", r)
	}
}
