package main

import (
	"fmt"
	"gross-store/models"
	"sync"
)

//TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.

func main() {
	store := &models.Store{
		Shorts:  20000,
		Jackets: 20000,
		Mu:      sync.Mutex{},
	}
	wg := &sync.WaitGroup{}
	shortsCH := make(chan int, store.Shorts)
	jacketsCH := make(chan int, store.Jackets)
	for i := 0; i < 25000; i++ {
		wg.Add(1)
		go func() {
			switch i % 2 {
			case 0:
				shortsCH <- store.RestShorts(i)
				wg.Done()
			case 1:
				jacketsCH <- store.RestJackets(i)
				wg.Done()
			}
		}()
	}

	wg.Wait()
	fmt.Println(store)
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
