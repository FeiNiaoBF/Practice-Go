package leetcode

import (
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	nums := []int{1, 1, 1, 2, 2, 3}
	rel := removeTwoDuplicates(nums)
	if rel != 5 {
		t.Errorf("Expected 3, but got %d", rel)
	}
}
