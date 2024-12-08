package learngocmp

import "fmt"

func ExampleDiffSubset() {
	diff := DiffSubset(map[string]any{
		"field1": 1,
		"field2": 2,
		"field4": 4,
	}, map[string]any{
		"field1": 11,
		"field2": 22,
		"field3": 33,
	})
	fmt.Println(diff)

	// Output:
	// map[string]any{
	// - 	"field1": int(1),
	// + 	"field1": int(11),
	// - 	"field2": int(2),
	// + 	"field2": int(22),
	//   	... // 1 ignored entry
	// - 	"field4": int(4),
	//   }
}
