package main

import (
	"bytes"
	"reflect"
	"testing"
)

const (
	sleep = "sleep"
	write = "write"
)

type SpySleeperAndWriter struct {
	Steps []string
}

func (s *SpySleeperAndWriter) Sleep() {
	s.Steps = append(s.Steps, sleep)
}

func (s *SpySleeperAndWriter) Write(p []byte) (n int, err error) {
	s.Steps = append(s.Steps, write)
	return 0, nil
}

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spy := &SpySleeperAndWriter{}

	Countdown(buffer, spy)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestCountdownSteps(t *testing.T) {
	spy := &SpySleeperAndWriter{}

	Countdown(spy, spy)

	want := []string{
		sleep,
		write,
		sleep,
		write,
		sleep,
		write,
		sleep,
		write,
	}

	if !reflect.DeepEqual(want, spy.Steps) {
		t.Errorf("got %q wanted %q", spy.Steps, want)
	}
}
