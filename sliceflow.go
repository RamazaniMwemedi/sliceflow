package sliceflow

import "reflect"

/**
Push appends elements or slices to a slice and returns the updated slice.
It takes the original slice and one or more elements or slices and adds them to the original slice.
The function is variadic, allowing users to push elements from multiple sources in a single call.

Example usage:
	originalSlice := []int{1, 2, 3}
	updatedSlice := push.Push(originalSlice, []int{4, 5, 6}, 7, 8, 9)

Parameters:
	- slice: The original slice to which elements will be appended.
	- sources: One or more elements or slices containing elements to be appended to the original slice.

Returns:
	- interface{}: The updated slice with elements appended.
*/
func PushMethod(slice interface{}, sources ...interface{}) interface{} {
	// Use reflection to check the type of the slice
	sliceValue := reflect.ValueOf(slice)
	if sliceValue.Kind() != reflect.Slice {
		panic("Push function can only be used with slices")
	}

	// Iterate over each source and append its elements to the original slice
	for _, src := range sources {
		srcValue := reflect.ValueOf(src)

		if srcValue.Kind() == reflect.Slice {
			// If the source is a slice, append its elements individually to the original slice
			for i := 0; i < srcValue.Len(); i++ {
				sliceValue = reflect.Append(sliceValue, srcValue.Index(i))
			}
		} else {
			// If the source is an individual element, append it to the original slice
			sliceValue = reflect.Append(sliceValue, srcValue)
			// Add more cases for other types if needed
		}
	}

	// Return the updated slice as interface{}
	return sliceValue.Interface()
}

/**
Filter applies a filtering condition to a slice and returns a new slice containing only the elements that satisfy the condition.
The condition is defined by the provided filter function, which takes an element and returns a boolean.

Example usage:
	numbers := []int{1, 2, 3, 4, 5, 6}
	filteredNumbers := filter.Filter(numbers, func(x interface{}) bool {
		return x.(int) % 2 == 0
	})

	users := []User{
		{Name: "Alice", Age: 12},
		{Name: "Bob", Age: 18},
		{Name: "Charlie", Age: 9},
		{Name: "David", Age: 14},
	}
	filteredUsers := filter.Filter(users, func(u interface{}) bool {
		return u.(User).Age >= 10 && u.(User).Age <= 15
	})

Parameters:
	- slice: The original slice to be filtered.
	- condition: The filtering condition defined by a function that takes an element and returns a boolean.

Returns:
	- interface{}: The new slice containing only the elements that satisfy the condition.
*/
func Filter(slice interface{}, condition func(interface{}) bool) interface{} {
	// Use reflection to check the type of the slice
	sliceValue := reflect.ValueOf(slice)
	if sliceValue.Kind() != reflect.Slice {
		panic("Filter function can only be used with slices")
	}

	var filteredSlice reflect.Value

	for i := 0; i < sliceValue.Len(); i++ {
		elem := sliceValue.Index(i).Interface()
		if condition(elem) {
			if !filteredSlice.IsValid() {
				filteredSlice = reflect.MakeSlice(sliceValue.Type(), 0, 0)
			}
			filteredSlice = reflect.Append(filteredSlice, reflect.ValueOf(elem))
		}
	}

	// Return the filtered slice as interface{}
	return filteredSlice.Interface()
}

