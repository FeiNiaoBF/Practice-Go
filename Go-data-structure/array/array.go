package array

import "fmt"

/// Array in Golang

// Array Data Structure
type array struct {
	data   []int
	length uint
}

// NewArray create a array
func NewArray(capacity uint) *array {
	if capacity == 0 {
		return nil
	}
	return &array{
		data:   make([]int, capacity),
		length: 0,
	}
}

/// Public function of array

// Len get the array length
func (arr *array) Len() uint {
	return arr.length
}

// Insert insert the value at the index
func (arr *array) Insert(val int, ind uint) (*array, error) {
	if arr.isIndexOutOfRange(ind) {
		return nil, fmt.Errorf("out of range")
	}
	if arr.Len() == uint(cap(arr.data)) {
		return nil, fmt.Errorf("full array")
	}
	for i := arr.Len(); i > ind; i-- {
		arr.data[i] = arr.data[i-1]
	}
	arr.data[ind] = val
	arr.length++
	return arr, nil
}

// Delete delete the value at the index
func (arr *array) Delete(ind uint) (*array, error) {
	if arr.isIndexOutOfRange(ind) {
		return nil, fmt.Errorf("out of range")
	}
	for i := ind; i < arr.Len()-1; i++ {
		arr.data[i] = arr.data[i+1]
	}
	arr.length--
	return arr, nil
}

// Find find the value at the index
func (arr *array) Find(ind uint) (int, error) {
	if arr.isIndexOutOfRange(ind) {
		return 0, fmt.Errorf("out of range")
	}
	return arr.data[ind], nil
}

// Print print the array
func (arr *array) Print() {
	var format string
	for i := uint(0); i < arr.Len(); i++ {
		format += fmt.Sprintf("%+v|", arr.data[i])
	}
	fmt.Println(format)
}

/// Private function of array

// isIndexOutOfRange check if the index is out of range
func (arr *array) isIndexOutOfRange(index uint) bool {
	return index >= uint(cap(arr.data))
}

/// Algorithm of array

// BobbleSort bobble sort
func (arr *array) Bobble() {
	if arr.Len() <= 1 {
		return
	}
	for i := uint(0); i < arr.Len(); i++ {
		for j := uint(0); j < arr.Len()-i-1; j++ {
			if arr.data[j] > arr.data[j+1] {
				arr.data[j], arr.data[j+1] = arr.data[j+1], arr.data[j]
			}
		}
	}
}

// BinarySearch
func (arr *array) BinarySearch(tar int) int {
	var (
		low  = 0
		high = int(arr.Len()) - 1
	)
	for low <= high {
		mid := low + (high-low)/2
		if arr.data[mid] > tar {
			high = mid - 1
		}else if arr.data[mid] < tar {
			low = mid + 1
		}else{
			return mid
		}
	}
	return -1
}
