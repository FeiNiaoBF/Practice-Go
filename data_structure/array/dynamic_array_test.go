package array

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewDarr(t *testing.T) {
	darr := NewDarr(10)
	if darr.Capacity() != 10 {
		t.Errorf("Expected capacity to be 10, got %d", darr.Capacity())
	}
	if darr.Size() != 0 {
		t.Errorf("Expected size to be 0, got %d", darr.Size())
	}
}

func TestPush(t *testing.T) {
	darr := NewDarr(10)
	darr.Push(1)
	if darr.Size() != 1 {
		t.Errorf("Expected size to be 1, got %d", darr.Size())
	}
	val, err := darr.At(0)
	if err != nil || val != 1 {
		t.Errorf("Expected value at index 0 to be 1, got %d", val)
	}
}

func TestInsert(t *testing.T) {
	darr := NewDarr(10)
	darr.Push(1)
	darr.Push(2)
	err := darr.Insert(1, 3)
	if err != nil {
		t.Errorf("Insert failed: %v", err)
	}
	val, err := darr.At(1)
	if err != nil || val != 3 {
		t.Errorf("Expected value at index 1 to be 3, got %d", val)
	}
}

func TestPrepend(t *testing.T) {
	darr := NewDarr(10)
	darr.Push(1)
	err := darr.Prepend(0)
	if err != nil {
		t.Errorf("Prepend failed: %v", err)
	}
	val, err := darr.At(0)
	if err != nil || val != 0 {
		t.Errorf("Expected value at index 0 to be 0, got %d", val)
	}
}

func TestPop(t *testing.T) {
	darr := NewDarr(10)
	darr.Push(1)
	val, err := darr.Pop()
	if err != nil || val != 1 {
		t.Errorf("Expected popped value to be 1, got %d", val)
	}
	if darr.Size() != 0 {
		t.Errorf("Expected size to be 0, got %d", darr.Size())
	}
}

func TestDelete(t *testing.T) {
	darr := NewDarr(10)
	darr.Push(1)
	darr.Push(2)
	err := darr.Delete(0)
	if err != nil {
		t.Errorf("Delete failed: %v", err)
	}
	val, err := darr.At(0)
	if err != nil || val != 2 {
		t.Errorf("Expected value at index 0 to be 2, got %d", val)
	}
}

func TestRemove(t *testing.T) {
	darr := NewDarr(10)
	darr.Push(1)
	darr.Push(2)
	darr.Push(1)
	darr.Push(2)
	darr.Push(1)
	removed, err := darr.Remove(1)
	fmt.Print(removed)
	if err != nil {
		t.Errorf("Remove failed: %v", err)
	}
	if len(removed) != 3 || !reflect.DeepEqual(removed, []int{0, 2, 4}) {
		t.Errorf("Expected removed item index to be [0, 2, 4], got %v", removed)
	}
	val, err := darr.At(1)
	if err != nil || val != 2 {
		t.Errorf("Expected value at index 1 to be 2, got %d", val)
	}
}

func TestFind(t *testing.T) {
	darr := NewDarr(10)
	darr.Push(1)
	darr.Push(2)
	index, err := darr.Find(2)
	if err != nil || index != 1 {
		t.Errorf("Expected index of 2 to be 1, got %d", index)
	}
}

func TestResize(t *testing.T) {
	darr := NewDarr(4)
	for i := 0; i < 8; i++ {
		darr.Push(i)
	}
	if darr.Capacity() != 8 {
		t.Errorf("Expected capacity to be 4 after resize, got %d", darr.Capacity())
	}
	for i := 0; i < 8; i++ {
		darr.Pop()
	}
	if darr.Capacity() != 2 {
		t.Errorf("Expected capacity to be 2 after pop, got %d", darr.Capacity())
	}
}
