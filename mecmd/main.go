package main

import (
	"flag"
	"fmt"

	"github.com/kaepa3/mecaber"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		panic("args error")
	}
	err, m := mecaber.CreateNew()
	if err != nil {
		panic(err)
	}
	defer m.Destroy()

	done := make(chan struct{})
	errCh, node := m.ParseToNode(args[0], done)
	if err != nil {
		panic(err)
	}

LOOP:
	for {
		select {
		case text := <-node:
			fmt.Println(text)
		case err := <-errCh:
			fmt.Println(err)
			break LOOP
		case <-done:
			break LOOP
		}
	}
}
