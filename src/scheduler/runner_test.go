package scheduler

import "testing"

func TestRunner(t *testing.T) {
	t1()
}

func t1() {
	HotSpotRunner{}.Run()
}
