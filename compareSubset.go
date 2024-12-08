package learngocmp

import (
	"reflect"

	"github.com/google/go-cmp/cmp"
)

func DiffSubset(x, y any) string {
	return cmp.Diff(x, y, cmp.FilterPath(func(p cmp.Path) bool {
		expected, _ := p.Last().Values()
		return expected.Kind() == reflect.Invalid
	}, cmp.Ignore()))
}
