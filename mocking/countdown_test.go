package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	// spySleeper := &SpySleeper{}

	Countdown(buffer, &SpyCountdownOperations{})

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	// if spySleeper.Calls != 3 {
	// 	t.Errorf("not enough calls to sleeper, want 3 got %d", spySleeper.Calls)
	// }

	t.Run("sleep before every print", func(t *testing.T) {
		spySleeperPrinter := &SpyCountdownOperations{}
		Countdown(spySleeperPrinter, spySleeperPrinter)

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleeperPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleeperPrinter.Calls)
		}
	})

}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}

// Final implementation of the spy to check the order of operation
type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const write = "write"
const sleep = "sleep"

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

// First implementation to track number of operations
// type SpySleeper struct {
// 	Calls int
// }

// func (s *SpySleeper) Sleep() {
// 	s.Calls++
// }
