package main

import (
	"math/rand"
	"sync"
	"time"
)

func benchmarkUnBufferedChannel() {
	ch := make(chan int64)
	t := time.NewTicker(1 * time.Minute)

	var wg sync.WaitGroup
	wg.Add(2)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for {
			select {
			case _ = <-t.C:
				close(ch)
				return
			default:
				ch <- rand.Int63()
			}
		}
	}(&wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for range ch {
		}
	}(&wg)

	wg.Wait()
}

func benchmarkBufferedChannel() {
	ch := make(chan int64, 1)
	t := time.NewTicker(1 * time.Minute)

	var wg sync.WaitGroup
	wg.Add(2)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for {
			select {
			case _ = <-t.C:
				close(ch)
				return
			default:
				ch <- rand.Int63()
			}
		}
	}(&wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for range ch {
		}
	}(&wg)

	wg.Wait()
}

func benchmarkBufferedChannelNConsumers(n int) {
	ch := make(chan int64, 1)
	t := time.NewTicker(1 * time.Minute)

	var wg sync.WaitGroup

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for {
			select {
			case _ = <- t.C:
				close(ch)
				return
			default:
				ch <- rand.Int63()
			}
		}
	}(&wg)

	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			for range ch {
			}
		}(&wg)
	}

	wg.Wait()
}

func benchmarkNBufferedChannelNConsumers(nBuf, n int) {
	ch := make(chan int64, nBuf)
	t := time.NewTicker(1 * time.Minute)

	var wg sync.WaitGroup

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for {
			select {
			case _ = <- t.C:
				close(ch)
				return
			default:
				ch <- rand.Int63()
			}
		}
	}(&wg)

	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			for range ch {
			}
		}(&wg)
	}

	wg.Wait()
}