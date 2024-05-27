package array

import "errors"

type Darray struct {
	array []int
	len   int
	cap   int
}

func NewDarr(size int) *Darray {
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
	if item >= darr.len {
		return errors.New("invalid item")
	}
	darr.len++
	darr.array = append(darr.array, item)
	return nil
}

func (darr *Darray) Insert(index, item int) error {
	if item >= darr.cap {
		return errors.New("invalid item")
	}
	if index >= darr.cap {
		return errors.New("invalid index")
	}

	darr.len++
	tmp := make([]int, darr.len)
	copy(tmp, darr.array[index:])

	darr.array = append(append(darr.array[:index], item), tmp...)

	return nil
}

func (darr *Darray) Prepend(item int) error {
	darr.len++
	return darr.Insert(0, item)
}

func (darr *Darray) Pop() (int, error) {
	item := darr.array[(darr.len - 1):]
	darr.len--
	return item[0], nil
}

func (darr *Darray) Delete(index int) error {
	if index >= darr.cap {
		return errors.New("invalid index")
	}
	darr.len--
	darr.array = append(darr.array[:index-1], darr.array[index:]...)
	return nil
}

func (darr *Darray) Remove(item int) ([]int, error) {
	result := make([]int, 0)
	for i, v := range darr.array {
		if v == item {
			darr.Delete(i)
			result = append(result, i)
		}
	}
	return result, nil
}

func (darr *Darray) Find(item int) (int, error) {
	for i, v := range darr.array {
		if v == item {
			return i, nil
		}
	}
	return 0, errors.New("not valut")
}

func (darr *Darray) resize(newCap int) {
}
