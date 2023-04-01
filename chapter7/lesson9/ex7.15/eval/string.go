/*
Exercise 7.13
Добавьте метод String к Ехрг для красивого вывода синтаксического дерева.
Убедитесь, что результаты при повторном анализе дают эквивалентное дерево.
*/

package eval

import (
	"fmt"
	"strings"
)

func (v Var) String() string {
	return string(v)
}

func (l literal) String() string {
	return fmt.Sprintf("%g", l)
}

func (u unary) String() string {
	return fmt.Sprintf("%c%s", u.op, u.x)
}

func (b binary) String() string {
	return fmt.Sprintf("(%s %c %s)", b.x, b.op, b.y)
}

func (c call) String() string {
	buf := &strings.Builder{}
	buf.WriteString(c.fn)
	buf.WriteByte('(')
	for i, arg := range c.args {
		if i > 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(arg.String())
	}
	buf.WriteByte(')')
	return buf.String()
}

func (p postUnary) String() string {
	return fmt.Sprintf("%g%c", p.x, p.op)
}
