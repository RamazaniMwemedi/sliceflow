package main

import (
	"fmt"

	"github.com/RamazaniMwemedi/sliceflow"
)

func main() {
	// Example 1: Filter even numbers from an integer slice
	numbers := []int{1, 2, 3, 4, 5, 6}
	filteredNumbers := sliceflow.Filter(numbers, func(x interface{}) bool {
		return x.(int)%2 == 0
	})
	fmt.Println("Filtered Even Numbers:", filteredNumbers)

	// Example 2: Filter users with age between 10 and 15 from a slice of structs
	type User struct {
		Name string
		Age  int
	}
	users := []User{
		{Name: "Alice", Age: 12},
		{Name: "Bob", Age: 18},
		{Name: "Charlie", Age: 9},
		{Name: "David", Age: 14},
	}
	filteredUsers := sliceflow.Filter(users, func(u interface{}) bool {
		return u.(User).Age >= 10 && u.(User).Age <= 15
	})
	fmt.Println("Filtered Users with Age between 10 and 15:", filteredUsers)

	// Example 3: Filter strings with length greater than 3 from a string slice
	strings := []string{"cat", "dog", "elephant", "rat"}
	filteredStrings := sliceflow.Filter(strings, func(s interface{}) bool {
		return len(s.(string)) > 3
	})
	fmt.Println("Filtered Strings with Length > 3:", filteredStrings)

	// Example 4: Filter positive numbers from a mixed-type slice
	mixedSlice := []interface{}{1, "two", -3, 4.5, -6, "seven"}
	filteredPositiveNumbers := sliceflow.Filter(mixedSlice, func(x interface{}) bool {
		switch v := x.(type) {
		case int:
			return v > 0
		case float64:
			return v > 0
		default:
			return false
		}
	})
	fmt.Println("Filtered Positive Numbers:", filteredPositiveNumbers)

	// Example 5: Filter elements equal to "apple" from a string slice
	fruits := []string{"apple", "banana", "apple", "orange"}
	filteredFruits := sliceflow.Filter(fruits, func(s interface{}) bool {
		return s.(string) == "apple"
	})
	fmt.Println("Filtered Elements Equal to 'apple':", filteredFruits)

	// Example 6: Filter elements based on a custom condition from a slice of structs
	type Item struct {
		Name  string
		Price float64
	}
	items := []Item{
		{Name: "Laptop", Price: 1200.0},
		{Name: "Phone", Price: 800.0},
		{Name: "Tablet", Price: 300.0},
	}
	filteredExpensiveItems := sliceflow.Filter(items, func(item interface{}) bool {
		return item.(Item).Price > 500.0
	})
	fmt.Println("Filtered Expensive Items:", filteredExpensiveItems)

	// Example 7: Filter elements based on a custom condition from a slice of integers
	integers := []int{5, 10, 15, 20, 25}
	filteredMultiplesOfFive := sliceflow.Filter(integers, func(x interface{}) bool {
		return x.(int)%5 == 0
	})
	fmt.Println("Filtered Multiples of Five:", filteredMultiplesOfFive)

	// Example 8: Filter elements with a specific prefix from a string slice
	words := []string{"apple", "banana", "cherry", "date"}
	prefix := "c"
	filteredWords := sliceflow.Filter(words, func(s interface{}) bool {
		return len(s.(string)) > 0 && s.(string)[:1] == prefix
	})
	fmt.Println("Filtered Words with Prefix 'c':", filteredWords)

	// Example 9: Filter elements based on a complex condition from a mixed-type slice
	complexSlice := []interface{}{1, "two", 3.0, "four", 5, "six"}
	filteredComplexSlice := sliceflow.Filter(complexSlice, func(x interface{}) bool {
		switch v := x.(type) {
		case int:
			return v%2 == 1
		case float64:
			return v > 3.0
		case string:
			return len(v) < 4
		default:
			return false
		}
	})
	fmt.Println("Filtered Complex Slice:", filteredComplexSlice)

	// Example 10: Filter elements based on a custom condition from a slice of booleans
	boolSlice := []bool{true, false, true, false, true}
	filteredTrueValues := sliceflow.Filter(boolSlice, func(x interface{}) bool {
		return x.(bool)
	})
	fmt.Println("Filtered True Values:", filteredTrueValues)
}
