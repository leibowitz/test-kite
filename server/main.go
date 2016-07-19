package main

import (
	"github.com/koding/kite"
	"github.com/leibowitz/test-kite/bla"
	"time"
	"fmt"
)

func main() {
	k := kite.New("first", "1.0.0")
	k.HandleFunc("square", func(r *kite.Request) (interface{}, error) {
		tStart := time.Now()
		a := bla.Bla{}
		err := r.Args.One().Unmarshal(a)
		fmt.Printf("%v done in %s\n", err, time.Now().Sub(tStart))
		a.A = "done"
		return a, nil // send back the result
	}).DisableAuthentication()
	k.Config.Port = 6000
	k.Run()
}
