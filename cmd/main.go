package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"proto-benchmark-value-vs-pointers/generator"
	"proto-benchmark-value-vs-pointers/proto"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	var (
		optional = make(chan []*proto.MessageOptional)
		value    = make(chan []*proto.MessageValue)
	)
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))

	}()
	go func() {
		for {
			go func() {
				optional <- generator.GenerateMessageOptional(10)
			}()
			go func() {
				value <- generator.GenerateMessageValue(10)
			}()
			time.Sleep(time.Millisecond * 10)
		}
	}()

	var optionalCount, valueCount int
	ticker := time.NewTicker(time.Second * 3)

	go func() {
		for {
			<-ticker.C
			fmt.Printf("optional msgs: %d\n", optionalCount)
			fmt.Printf("value msgs:    %d\n", valueCount)
		}
	}()

	for {
		select {
		case msg := <-optional:
			optionalCount += len(msg)
		case msg := <-value:
			valueCount += len(msg)
		}
	}
}
