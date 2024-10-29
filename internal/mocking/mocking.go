package mocking

import (
	"fmt"
	"io"
	"time"
)

const FINAL_WORD = "Go!"
const COUNTDOWN_STARTING = 3

func Countdown(writer io.Writer, sleeper Sleeper) {
	for i := COUNTDOWN_STARTING; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(writer, i)
	}
	sleeper.Sleep()
	fmt.Fprintf(writer, FINAL_WORD)
}

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}
