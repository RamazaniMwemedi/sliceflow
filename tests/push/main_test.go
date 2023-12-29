package push_test

import (
	"reflect"
	"testing"

	sliceflow "github.com/RamazaniMwemedi/sliceflow"
)

// TestPush_AppendSlices tests the Push function by appending slices to the original slice.
func TestPush_AppendSlices(t *testing.T) {
	originalSlice := []int{1, 2, 3}
	expectedSlice := []int{1, 2, 3, 4, 5, 6}
	updatedSlice := sliceflow.PushMethod(originalSlice, []int{4, 5, 6})
	if !reflect.DeepEqual(updatedSlice, expectedSlice) {
		t.Errorf("Test failed. Expected: %v, Got: %v", expectedSlice, updatedSlice)
	}
}

// TestPush_AppendMultipleSources tests the Push function by appending slices and individual elements to the original slice.
func TestPush_AppendMultipleSources(t *testing.T) {
	originalSlice := []int{1, 2, 3}
	expectedSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	updatedSlice := sliceflow.PushMethod(originalSlice, []int{4, 5, 6}, 7, 8, 9)
	if !reflect.DeepEqual(updatedSlice, expectedSlice) {
		t.Errorf("Test failed. Expected: %v, Got: %v", expectedSlice, updatedSlice)
	}
}

// TestPush_AppendEmptySlice tests the Push function by appending an empty slice to the original slice.
func TestPush_AppendEmptySlice(t *testing.T) {
	originalSlice := []int{1, 2, 3}
	expectedSlice := []int{1, 2, 3}
	updatedSlice := sliceflow.PushMethod(originalSlice, []int{})
	if !reflect.DeepEqual(updatedSlice, expectedSlice) {
		t.Errorf("Test failed. Expected: %v, Got: %v", expectedSlice, updatedSlice)
	}
}

// TestPush_AppendSingleElement tests the Push function by appending a single element to the original slice.
func TestPush_AppendSingleElement(t *testing.T) {
	originalSlice := []int{1, 2, 3}
	expectedSlice := []int{1, 2, 3, 4}
	updatedSlice := sliceflow.PushMethod(originalSlice, 4)
	if !reflect.DeepEqual(updatedSlice, expectedSlice) {
		t.Errorf("Test failed. Expected: %v, Got: %v", expectedSlice, updatedSlice)
	}
}
