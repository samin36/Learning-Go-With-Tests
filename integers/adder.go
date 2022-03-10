package integers

import "fmt"

func Add(a, b int) int {
	return a + b
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
