/*
Exercise 7.13
Добавьте метод String к Ехрг для красивого вывода синтаксического дерева.
Убедитесь, что результаты при повторном анализе дают эквивалентное дерево.
*/

package eval

import "testing"

func TestString(t *testing.T) {
	tcs := []struct {
		expr string
		want string
	}{
		{"-1 + -x", "(-1 + -x)"},
		{"-1 - x", "(-1 - x)"},
		{"sqrt(A / pi)", "sqrt((A / pi))"},
		{"pow(x, 3) + pow(y, 3)", "(pow(x, 3) + pow(y, 3))"},
		{"5 / 9 * (F - 32)", "((5 / 9) * (F - 32))"},
	}

	for i, tc := range tcs {
		expr, err := Parse(tc.expr)
		if err != nil {
			t.Error(err)
			continue
		}
		got := expr.String()
		if got != tc.want {
			t.Fatalf("%d. got %v, expr %v", i, got, tc.want)
		}
	}
}
