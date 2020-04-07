package main

import (
	"fmt"
	"os"
	"os/signal"
	"testing"
	"time"
)

func TestStop(t *testing.T) {
	var count int
	count = 0
	for {
		fmt.Println("执行", count)
		count++
		time.Sleep(time.Second)
	}
}

func TestInterrupt(t *testing.T) {
	signals := make(chan os.Signal)
	signal.Notify(signals)
	s := <- signals
	fmt.Println(s)
}
