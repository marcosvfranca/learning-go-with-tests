package main

import (
	"testing"
	"time"
)

// func TestScore(t *testing.T) {
// 	t.Run("test a counting", func(t *testing.T) {
// 		buffer := &bytes.Buffer{}

// 		Score(buffer, &SpyOperationCount{})

// 		result := buffer.String()
// 		expected := `3
// 	2
// 	1
// 	Vai!`

// 		if result != expected {
// 			t.Errorf("result '%s', expected '%s'", result, expected)
// 		}
// 	})

// 	t.Run("pause before print", func(t *testing.T) {
// 		buffer := &bytes.Buffer{}
// 		spyPrintSleep := &SpyOperationCount{}
// 		Score(spyPrintSleep, spyPrintSleep)

// 		expected := []string{
// 			pause,
// 			write,
// 			pause,
// 			write,
// 			pause,
// 			write,
// 			pause,
// 			write,
// 		}

// 		if !reflect.DeepEqual(expected, spyPrintSleep.Calls) {
// 			t.Errorf("expected '%v' calss, result '%v'", expected, spyPrintSleep.Calls)
// 		}
// 	})
// }

func TestSleeperConfigurable(t *testing.T) {
	timePause := 5 * time.Second

	timeSpy := &TimeSpy{}
	sleeper := SleeperConfigurable{timePause, timeSpy.Pause}
	sleeper.Pause()

	if timeSpy.durationPause != timePause {
		t.Errorf("has been paused for %v, but it should paused for %v", timePause, timeSpy.durationPause)
	}
}
