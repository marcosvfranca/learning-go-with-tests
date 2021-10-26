package main

import (
	"fmt"
	"io"
	"time"
)

const startCount = 3
const finalMessage = "Vai!"
const write = "write"
const pause = "pause"

// func main() {
// 	sleeper := &SleeperConfigurable{1 * time.Second, time.Sleep}
// 	Score(os.Stdout, sleeper)
// }

func Score(buffer io.Writer, sleeper Sleeper) {
	for i := startCount; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(buffer, i)
	}

	for i := startCount; i > 0; i-- {
		fmt.Fprintln(buffer, i)
	}

	sleeper.Sleep()
	fmt.Fprint(buffer, finalMessage)
}

type Sleeper interface {
	Sleep()
}

type SleeperSpy struct {
	Calls int
}

func (s *SleeperSpy) Sleep() {
	s.Calls++
}

type SpyOperationCount struct {
	Calls []string
}

func (s *SpyOperationCount) Pause() {
	s.Calls = append(s.Calls, pause)
}

func (s *SpyOperationCount) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

type SleeperConfigurable struct {
	duration time.Duration
	pause    func(time.Duration)
}

type TimeSpy struct {
	durationPause time.Duration
}

func (t *TimeSpy) Pause(duration time.Duration) {
	t.durationPause = duration
}

func (s *SleeperConfigurable) Pause() {
	s.pause(s.duration)
}
