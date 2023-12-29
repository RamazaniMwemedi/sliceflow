package filter_test

import (
	"fmt"
	"reflect"
	"testing"

	sliceflow "github.com/RamazaniMwemedi/sliceflow"
)

type User struct {
	Name string
	Age  int
}

func TestFilter_Numbers(t *testing.T) {
	// Test case: Filter even numbers
	numbers := []int{1, 2, 3, 4, 5, 6}
	filteredNumbers := sliceflow.Filter(numbers, func(x interface{}) bool {
		return x.(int)%2 == 0
	})

	expectedNumbers := []int{2, 4, 6}
	if !reflect.DeepEqual(filteredNumbers, expectedNumbers) {
		t.Errorf("Test failed. Expected: %v, Got: %v", expectedNumbers, filteredNumbers)
	}
}

func TestFilter_Structs(t *testing.T) {
	// Test case: Filter users with age between 10 and 15
	users := []User{
		{Name: "Alice", Age: 12},
		{Name: "Bob", Age: 18},
		{Name: "Charlie", Age: 9},
		{Name: "David", Age: 14},
	}
	filteredUsers := sliceflow.Filter(users, func(u interface{}) bool {
		return u.(User).Age >= 10 && u.(User).Age <= 15
	})
	fmt.Println(filteredUsers)
	expectedUsers := []User{
		{Name: "Alice", Age: 12},
		{Name: "David", Age: 14},
	}
	if !reflect.DeepEqual(filteredUsers, expectedUsers) {
		t.Errorf("Test failed. Expected: %v, Got: %v", expectedUsers, filteredUsers)
	}
}

func TestFilter_InvalidSlice(t *testing.T) {
	// Test case: Try to filter a non-slice value
	nonSlice := 42
	defer func() {
		if recover() == nil {
			t.Error("Test failed. Expected a panic for non-slice input.")
		}
	}()

	_ = sliceflow.Filter(nonSlice, func(x interface{}) bool {
		return true
	})
}

func ExampleFilter() {
	// Example usage with numbers:
	numbers := []int{1, 2, 3, 4, 5, 6}
	filteredNumbers := sliceflow.Filter(numbers, func(x interface{}) bool {
		return x.(int)%2 == 0
	})
	fmt.Println("Filtered Numbers:", filteredNumbers)

	// Example usage with structs:
	users := []User{
		{Name: "Alice", Age: 12},
		{Name: "Bob", Age: 18},
		{Name: "Charlie", Age: 9},
		{Name: "David", Age: 14},
	}
	filteredUsers := sliceflow.Filter(users, func(u interface{}) bool {
		return u.(User).Age >= 10 && u.(User).Age <= 15
	})
	fmt.Println("Filtered Users:", filteredUsers)

	// Output:
	// Filtered Numbers: [2 4 6]
	// Filtered Users: [{Alice 12} {David 14}]
}
