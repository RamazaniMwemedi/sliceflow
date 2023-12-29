package main

import (
	"fmt"

	"github.com/RamazaniMwemedi/sliceflow"
)

func main() {
	// Example 1: Append elements to an integer slice
	intSlice := []int{1, 2, 3}
	updatedIntSlice := sliceflow.PushMethod(intSlice, 7, 8, 9)
	fmt.Println("Updated Integer Slice:", updatedIntSlice)

	// Example 2: Append elements to a string slice
	stringSlice := []string{"apple", "banana"}
	updatedStringSlice := sliceflow.PushMethod(stringSlice, "cherry", "date")
	fmt.Println("Updated String Slice:", updatedStringSlice)

	// Example 3: Append elements to a slice of structs
	type Person struct {
		Name string
		Age  int
	}
	personSlice := []Person{{"Alice", 25}, {"Bob", 30}}
	updatedPersonSlice := sliceflow.PushMethod(personSlice, Person{"Charlie", 22}, Person{"David", 28})
	fmt.Println("Updated Person Slice:", updatedPersonSlice)

	// Example 4: Append elements to an empty slice
	emptySlice := []int{}
	updatedEmptySlice := sliceflow.PushMethod(emptySlice, 1, 2, 3)
	fmt.Println("Updated Empty Slice:", updatedEmptySlice)

	// Example 5: Append elements from multiple slices
	slice1 := []int{1, 2}
	slice2 := []int{3, 4, 5}
	slice3 := []int{6, 7}
	combinedSlice := sliceflow.PushMethod(slice1, slice2, slice3)
	fmt.Println("Combined Slice:", combinedSlice)

	// Example 6: Append elements of different types using variadic parameters
	mixedSlice := []interface{}{1, "two", 3.0}
	updatedMixedSlice := sliceflow.PushMethod(mixedSlice, "four", true)
	fmt.Println("Mixed Type Slice:", updatedMixedSlice)

	// Example 7: Append a single element to a slice
	singleElementSlice := []int{1, 2, 3}
	updatedSingleElementSlice := sliceflow.PushMethod(singleElementSlice, 4)
	fmt.Println("Updated Single Element Slice:", updatedSingleElementSlice)

	// Example 8: Append a single element to an empty slice
	emptySlice2 := []string{}
	updatedEmptySlice2 := sliceflow.PushMethod(emptySlice2, "hello")
	fmt.Println("Updated Empty String Slice:", updatedEmptySlice2)

	// Example 9: Append elements from multiple slices with different types
	multipleTypeSlice := []interface{}{1, "two"}
	updatedMultipleTypeSlice := sliceflow.PushMethod(multipleTypeSlice, 3.0, "four", true)
	fmt.Println("Combined Multiple Type Slice:", updatedMultipleTypeSlice)

	// Example 10: Append an empty slice to another slice
	targetSlice := []int{1, 2, 3}
	emptySlice3 := []int{}
	updatedTargetSlice := sliceflow.PushMethod(targetSlice, emptySlice3)
	fmt.Println("Updated Target Slice with Empty Slice:", updatedTargetSlice)
}
