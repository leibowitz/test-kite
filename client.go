package main

import (
	"fmt"
	"runtime"
	"sync"

	"flag"
	"github.com/jeffail/tunny"
	"github.com/koding/kite"
	"github.com/leibowitz/test_kite/bla"
	"time"
)

func process(client *kite.Client) interface{} {
	tStart := time.Now()
	response, err := client.Tell("square", bla.Bla{})
	if err != nil {
		fmt.Printf("%s\n", err)
	} else {
		a := bla.Bla{}
		err := response.Unmarshal(&a)
		if err != nil {
			fmt.Printf("%s\n", err)
		} else {
			fmt.Printf("%v in %s\n", a, time.Now().Sub(tStart))
		}
	}
	return response
}

func main() {
	var count = flag.Int("c", 100, "count")
	flag.Parse()

	k := kite.New("second", "1.0.0")

	client := k.NewClient("http://localhost:6000/kite")
	client.Dial()

	wg := &sync.WaitGroup{}

	pool, _ := tunny.CreatePool(runtime.NumCPU(), func(object interface{}) interface{} {
		//input, _ := object.([]byte)
		response := process(client)

		//client.Tell("square", 4)
		//fmt.Printf("in %s\n", time.Now().Sub(tStart))
		// Do something that takes a lot of work
		//output := input
		defer wg.Done()

		return response
	}).Open()

	defer pool.Close()

	tStart := time.Now()
	for i := 0; i < *count; i++ {
		wg.Add(1)
		go pool.SendWork(nil)
	}

	wg.Wait()

	elapsed := time.Now().Sub(tStart)
	fmt.Printf("total: %s, avg: %s\n", elapsed, elapsed/time.Duration(*count))
}
