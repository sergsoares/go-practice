package main

import (
	"testing"
	"time"
)

func Test_slow1(t *testing.T) {
	t.Parallel()
	time.Sleep(1 * time.Second)
}

func Test_slow2(t *testing.T) {
	t.Parallel()
	time.Sleep(1 * time.Second)
}
func Test_slow3(t *testing.T) {
	t.Parallel()
	time.Sleep(1 * time.Second)
}
func Test_slow4(t *testing.T) {
	t.Parallel()
	time.Sleep(1 * time.Second)
}
func Test_slow5(t *testing.T) {
	t.Parallel()
	time.Sleep(1 * time.Second)
}
