# SliceFlow

SliceFlow is a Go package that provides utility functions for working with slices. It includes methods for pushing elements to a slice and filtering a slice based on a custom condition.

## Installation

To use SliceFlow in your Go project, you can install it with the following command:

```bash
go get -u github.com/RamazaniMwemedi/sliceflow
```

## Usage

### Push Method

```go
package main

import (
	"fmt"
	"github.com/RamazaniMwemedi/sliceflow"
)

func main() {
	originalSlice := []int{1, 2, 3}
	updatedSlice := sliceflow.PushMethod(originalSlice, []int{4, 5, 6}, 7, 8, 9)
	fmt.Println("Updated Slice:", updatedSlice)
}
```

### Filter Method

```go
package main

import (
	"fmt"
	"github.com/RamazaniMwemedi/sliceflow"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6}
	filteredNumbers := sliceflow.Filter(numbers, func(x interface{}) bool {
		return x.(int) % 2 == 0
	})
	fmt.Println("Filtered Numbers:", filteredNumbers)
}
```

## API Documentation

### `PushMethod`

Push appends elements or slices to a slice and returns the updated slice.

#### Parameters:

- `slice`: The original slice to which elements will be appended.
- `sources`: One or more elements or slices containing elements to be appended to the original slice.

#### Returns:

- `interface{}`: The updated slice with elements appended.

### `Filter`

Filter applies a filtering condition to a slice and returns a new slice containing only the elements that satisfy the condition.

#### Parameters:

- `slice`: The original slice to be filtered.
- `condition`: The filtering condition defined by a function that takes an element and returns a boolean.

#### Returns:

- `interface{}`: The new slice containing only the elements that satisfy the condition.

## Contributing

Contributions to SliceFlow are welcome! If you find any issues or have ideas for improvements, please open an issue or submit a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
