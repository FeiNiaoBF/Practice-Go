package array

import "errors"

type Darray struct {
	array []int
	len   int
	cap   int
}

func NewDarr(size int) *Darray {
	if size <= 0 {
		size = 16
	}
	return &Darray{
		array: make([]int, 0, size),
		len:   0,
		cap:   size,
	}
}

func (darr *Darray) Size() int {
	return darr.len
}

func (darr *Darray) Capacity() int {
	return darr.cap
}

func (darr *Darray) IsEmpty() bool {
	return darr.len == 0
}

func (darr *Darray) At(index int) (int, error) {
	if index >= darr.len {
		return 0, errors.New("invalid index")
	}

	return darr.array[index], nil
}

func (darr *Darray) Push(item int) error {
	if darr.len == darr.cap {
		darr.resize(darr.cap * 2)
	}
	darr.len++
	darr.array = append(darr.array, item)
	return nil
}

func (darr *Darray) Insert(index, item int) error {

	if index >= darr.len {
		return errors.New("invalid index")
	}

	if darr.len == darr.cap {
		darr.resize(darr.cap * 2)
	}

	tmp := make([]int, darr.len)
	copy(tmp, darr.array[index:])

	darr.array = append(append(darr.array[:index], item), tmp...)
	darr.len++
	return nil
}

func (darr *Darray) Prepend(item int) error {
	darr.len++
	if darr.len == darr.cap {
		darr.resize(darr.cap * 2)
	}
	return darr.Insert(0, item)
}

func (darr *Darray) Pop() (int, error) {
	val := darr.array[darr.len - 1]
	if darr.cap / 3 >= darr.len {
		darr.resize(darr.cap / 2)
	}
	darr.len--
	return val, nil
}

func (darr *Darray) Delete(index int) error {
	if index >= darr.len {
		return errors.New("invalid index")
	}
	if darr.cap / 3 >= darr.len {
		darr.resize(darr.cap / 2)
	}
	darr.array = append(darr.array[:index], darr.array[index+1:]...)
	darr.len--
	return nil
}

// Remove  return the index of the removed item
func (darr *Darray) Remove(item int) ([]int, error) {
	result := make([]int, 0)
	for i, v := range darr.array {
		if v == item {
			result = append(result, i)
		}
	}
	if len(result) == 0 {
		return nil, errors.New("not found")
	}
	for i := len(result) - 1; i >= 0; i-- {
		darr.Delete(result[i])
	}
	return result, nil
}

func (darr *Darray) Find(item int) (int, error) {
	for i, v := range darr.array {
		if v == item {
			return i, nil
		}
	}
	return 0, errors.New("not found")
}

func (darr *Darray) resize(newCap int) {
	newArray := make([]int, darr.len, newCap)
	copy(newArray, darr.array)
	darr.array = newArray
	darr.cap = newCap
}
