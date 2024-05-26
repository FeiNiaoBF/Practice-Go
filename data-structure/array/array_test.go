package array_test

import (
	"GoAlgo/array"
	"testing"
)
func TestMaint (t *testing.T) {
	TestArray_Insert(t)
	TestArray_Delete(t)
	TestArray_Find(t)
	TestPrint(t)
}


func TestArray_Insert(t *testing.T) {

	array := array.NewArray(10)
	array.Insert(1, 0)
	array.Insert(2, 1)
	array.Insert(3, 2)
	array.Insert(4, 3)
	array.Insert(5, 4)

	if array.Len() != 5 {
		t.Errorf("array length should be 5, got %d", array.Len())
	}
}

func TestArray_Delete(t *testing.T) {
	array := array.NewArray(10)
	array.Insert(1, 0)
	array.Insert(2, 1)
	array.Insert(3, 2)
	array.Insert(4, 3)
	array.Insert(5, 4)

	array.Delete(3)
	array.Delete(2)
	if array.Len() != 3 {
		t.Errorf("array length should be 3, got %d", array.Len())
	}
}

func TestArray_Find(t *testing.T) {
	array := array.NewArray(10)
	array.Insert(1, 0)
	array.Insert(2, 1)
	array.Insert(3, 2)
	array.Insert(4, 3)
	array.Insert(5, 4)

	val, err := array.Find(3)
	if err != nil {
		t.Errorf("Error occurred: %v", err)
	}
	if val != 4 {
		t.Errorf("Expected 4, got %d", val)
	}
}

func TestPrint(t *testing.T) {
    array := array.NewArray(10)
    array.Insert(1, 0)
    array.Insert(2, 1)
    array.Insert(3, 2)
    array.Insert(4, 3)
    array.Insert(5, 4)

	array.Print()
}
