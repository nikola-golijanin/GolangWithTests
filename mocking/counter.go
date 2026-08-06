package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)

	sleeper2 := &ConfigurableSleeper{1 * time.Millisecond, time.Sleep}
	Countdown(os.Stdout, sleeper2)
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprint(out, finalWord)
}

type Sleeper interface {
	Sleep()
}

// DefaultSleeper
type DefaultSleeper struct {
}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

// ConfigurableSleeper
type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}
