package main

import (
	"fmt"
	"github.com/ignite-laboratories/core/impulse"
	"time"
)

var pulseDuration = time.Second / 2
var period = 100000
var clock = impulse.NewClock(period)

func main() {
	clock.OnDownbeat(Observe)
	clock.Start()
}

var lastNow = time.Now()

func Observe(ctx impulse.Context) {
	now := time.Now()
	lastDuration := now.Sub(lastNow)

	if lastDuration < pulseDuration {
		period += 5000
	} else if lastDuration > pulseDuration {
		period -= 5000
	}
	clock.Period = period

	fmt.Printf("%v | %v\n", lastDuration, clock.Period)
	lastNow = now
}
