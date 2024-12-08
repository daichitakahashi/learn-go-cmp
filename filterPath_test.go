package learngocmp

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
)

func ExampleFilterPath_SimilarMap() {
	cmp.Diff(map[string]any{
		"field1": 1,
		"field2": "2",
	}, map[string]any{
		"field1": "3",
		"field2": 4,
	}, cmp.FilterPath(func(p cmp.Path) bool {
		fmt.Println("path:", len(p))
		expected, actual := p.Last().Values()
		fmt.Println("  expected:", expected)
		fmt.Println("  actual:", actual)
		return false
	}, cmp.Ignore()))

	// Output:
	// path: 1
	//   expected: map[field1:1 field2:2]
	//   actual: map[field1:3 field2:4]
	// path: 2
	//   expected: 1
	//   actual: 3
	// path: 2
	//   expected: 2
	//   actual: 4
	// path: 1
	//   expected: map[field1:1 field2:2]
	//   actual: map[field1:3 field2:4]
	// path: 2
	//   expected: 1
	//   actual: 3
	// path: 2
	//   expected: 2
	//   actual: 4
}

func ExampleFilterPath_DifferentMap() {
	cmp.Diff(map[string]any{
		"field1": 1,
		"field2": "2",
	}, map[string]any{
		"field3": "3",
		"field4": 4,
	}, cmp.FilterPath(func(p cmp.Path) bool {
		fmt.Println("path:", len(p))
		expected, actual := p.Last().Values()
		fmt.Println("  expected:", expected)
		fmt.Println("  actual:", actual)
		return false
	}, cmp.Ignore()))

	// Output:
	// path: 1
	//   expected: map[field1:1 field2:2]
	//   actual: map[field3:3 field4:4]
	// path: 2
	//   expected: 1
	//   actual: <invalid reflect.Value>
	// path: 2
	//   expected: 2
	//   actual: <invalid reflect.Value>
	// path: 2
	//   expected: <invalid reflect.Value>
	//   actual: 3
	// path: 2
	//   expected: <invalid reflect.Value>
	//   actual: 4
	// path: 1
	//   expected: map[field1:1 field2:2]
	//   actual: map[field3:3 field4:4]
	// path: 2
	//   expected: 1
	//   actual: <invalid reflect.Value>
	// path: 2
	//   expected: 2
	//   actual: <invalid reflect.Value>
	// path: 2
	//   expected: <invalid reflect.Value>
	//   actual: 3
	// path: 2
	//   expected: <invalid reflect.Value>
	//   actual: 4
}
