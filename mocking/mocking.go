package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (cs ConfigurableSleeper) Sleep() {
	cs.sleep(cs.duration)
}

func Countdown(w io.Writer, sleeper Sleeper) {
	for i := 3; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(w, i)
	}

	sleeper.Sleep()
	fmt.Fprint(w, "Go!")
}

func main() {
	realSleeper := ConfigurableSleeper{
		duration: 1 * time.Second,
		sleep:    time.Sleep,
	}

	Countdown(os.Stdout, realSleeper)
}
