package main

import "testing"

func TestSerial(t *testing.T) {
	fetched := make(map[string]bool)
	Serial("http://golang.org/", fetcher, fetched)
}

func TestConcurrentMutex(t *testing.T) {
	state := makeState()
	ConcurrentMutex("http://golang.org/", fetcher, state)
}

func TestConcurrentChannel(t *testing.T) {
	ConcurrentChannel("http://golang.org/", fetcher)
}

func BenchmarkSerial(t *testing.B) {
	fetched := make(map[string]bool)
	for i := 0; i < t.N; i++ {
		Serial("http://golang.org/", fetcher, fetched)
	}
}

func BenchmarkConcurrentMutex(t *testing.B) {
	state := makeState()
	for i := 0; i < t.N; i++ {
		ConcurrentMutex("http://golang.org/", fetcher, state)
	}
}

func BenchmarkConcurrentChannel(t *testing.B) {
	for i := 0; i < t.N; i++ {
		ConcurrentChannel("http://golang.org/", fetcher)
	}
}
