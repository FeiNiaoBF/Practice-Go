package singleflight

import (
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestDo(t *testing.T) {
	var g Group

	var v any
	var wg sync.WaitGroup

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(val any, i int) {
			var err error
			v, err = g.Do("key", func() (any, error) {
				return "bar:" + strconv.Itoa(i), nil
			})
			if err != nil {
				t.Errorf("Do error = %v", err)
			}
			wg.Done()
		}(v, i)
	}

	// if v != "bar" || err != nil {
	// 	t.Errorf("Do v = %v, error = %v", v, err)
	// }
	wg.Wait()
	// want vale is random
	// TODO: fix it???
	if get, want := fmt.Sprintf("%v (%T)", v, v), "bar:0 (string)"; get != want {
		t.Errorf("Do v = %v T = %T", v, v)
	}
}

func TestDoErr(t *testing.T) {
	var g Group

	someErr := fmt.Errorf("some error")
	v, err := g.Do("key", func() (interface{}, error) {
		return nil, someErr
	})
	if err != someErr {
		t.Errorf("Do error = %v; want someErr", err)
	}
	if v != nil {
		t.Errorf("unexpected non-nil value %#v", v)
	}
}

func TestDoDupSuppress(t *testing.T) {
	var g Group

	c := make(chan string)
	var calls int32

	fn := func() (any, error) {
		atomic.AddInt32(&calls, 1)
		return <-c, nil
	}

	const n = 10
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			v, err := g.Do("key", fn)
			if err != nil {
				t.Errorf("Do error = %v; want nil", err)
			}
			if v.(string) != "bar" {
				t.Errorf("Do value = %q; want %q", v, "bar")
			}
			wg.Done()
		}()
	}
	time.Sleep(100 * time.Millisecond)
	c <- "bar"
	wg.Wait()
	if got := atomic.LoadInt32(&calls); got != 1 {
		t.Errorf("number of calls = %d; want 1", got)
	}
}
