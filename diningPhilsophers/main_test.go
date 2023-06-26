package main

import (
	"testing"
	"time"
)

func TestMain(t *testing.T) {
	sleepTime = 0 * time.Second
	eatTime = 0 * time.Second
	thinkTime = 0 * time.Second

	for i := 0; i < 1000; i++ {
		main()
		if len(finishingList) != 5 {
			t.Error("wrong number of finishing")
		}

		finishingList = []string{}
	}
}
