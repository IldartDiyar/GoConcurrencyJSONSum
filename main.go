package main

import (
	"encoding/json"
	"flag"
	"io"
	"log"
	"os"
	"sync"
	"time"
)

type Data struct {
	A int `json:"a"`
	B int `json:"b"`
}

func main() {
	var maxGr int
	flag.IntVar(&maxGr, "rt", 5, "max goroutines")
	flag.Parse()

	if maxGr < 0 {
		log.Fatal("Inccorect input for max goroutines ")
	}

	data, err := getData()
	if err != nil {
		log.Fatalf("Can not get data due to %v", err)
	}

	sem := NewSemaphore(maxGr)
	var wg sync.WaitGroup
	resu := make(chan int)
	st := time.Now()
	go func() {
		for _, dt := range data {
			wg.Add(1)
			sem.Acquire()

			go func(dt Data) {
				defer wg.Done()
				defer sem.Release()
				sum := dt.A + dt.B
				resu <- sum
			}(dt)
		}
		wg.Wait()
		close(resu)
	}()
	res := 0
	for r := range resu {
		res += r
	}

	log.Printf("Total sum: %d\n", res)
	log.Printf("Total time: %s\n", time.Since(st))
}

func getData() ([]Data, error) {
	file, err := os.Open("jj.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var data []Data
	if err := json.Unmarshal(byteValue, &data); err != nil {
		return nil, err
	}
	return data, nil
}

type Semaphore struct {
	gg chan struct{}
}

func NewSemaphore(maxGoroutines int) *Semaphore {
	return &Semaphore{gg: make(chan struct{}, maxGoroutines)}
}
func (s *Semaphore) Acquire() {
	s.gg <- struct{}{}
}
func (s *Semaphore) Release() {
	<-s.gg
}
