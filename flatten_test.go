package flatten

import (
	"reflect"
	"testing"
)

func TestCases(t *testing.T) {
	cases := []struct {
		input    []interface{}
		expected []interface{}
	}{
		{
			input:    []interface{}{[]interface{}{1, 2, []int{3}}, 4},
			expected: []interface{}{1, 2, 3, 4},
		},
		{
			input:    []interface{}{[4]int{1, 2, 3, 4}},
			expected: []interface{}{1, 2, 3, 4},
		},
		{
			input:    []interface{}{10, []int{20, 30}, []interface{}{40, []int{50, 60}}},
			expected: []interface{}{10, 20, 30, 40, 50, 60},
		},
		{
			input:    []interface{}{"alfa", [...]string{"beta", "gamma"}},
			expected: []interface{}{"alfa", "beta", "gamma"},
		},
	}

	for _, c := range cases {
		output := It(c.input)
		if !reflect.DeepEqual(c.expected, output) {
			t.Errorf("Expected output to be %q but it was %q", c.expected, output)
		}
	}
}
