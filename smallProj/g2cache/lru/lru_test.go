package lru

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	key1 := "key1"
	lru := New(0)
	lru.Add(key1, 1)
	val, ok := lru.Get(key1)
	if !ok {
		t.Fatalf("%s: cache hit = %v; want %v", key1, ok, !ok)
	} else if val != 1 && ok {
		t.Fatalf("%s expected get to return 1 but got %v", key1, val)
	} else {
		t.Logf("%s: Ok test pass get key1=%v", key1, val)
	}
}

func TestRemove(t *testing.T) {
	lru := New(0)
	lru.Add("myKey", 1234)
	if val, ok := lru.Get("myKey"); !ok {
		t.Fatal("TestRemove returned no match")
	} else if val != 1234 {
		t.Fatalf("TestRemove failed.  Expected %d, got %v", 1234, val)
	}

	lru.Remove("myKey")
	if _, ok := lru.Get("myKey"); ok {
		t.Fatal("TestRemove returned a removed entry")
	}
}

func TestOnEvicted(t *testing.T) {
	keys := make([]Key, 0)
	callback := func(key Key, value any) {
		keys = append(keys, key)
	}
	lru := New(10)
	lru.OnEvicted = callback
	for i := 0; i < 12; i++ {
		lru.Add(fmt.Sprintf("myKey%d", i), i)
	}

	if len(keys) != 2 {
		t.Fatalf("Call OnEvicted failed, expect keys length equals to 2, but got %d", len(keys))
	}

	if keys[0] != Key("myKey0") {
		t.Fatalf("got %v in first evicted key; want %s", keys[0], "myKey0")
	}
	if keys[1] != Key("myKey1") {
		t.Fatalf("got %v in second evicted key; want %s", keys[1], "myKey1")
	}

}
