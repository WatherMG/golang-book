/*
Example 7.10
Пакет eval предоставляет вычислитель выражений.
*/

package eval

import (
	"fmt"
	"math"
)

// Env сопоставляет имена переменных со значениями
type Env map[Var]float64

// Eval выполняет поиск в среде, который возвращает нуль, если переменная не определена.
func (v Var) Eval(env Env) float64 {
	return env[v]
}

// Eval возвращает значение литерала
func (l literal) Eval(_ Env) float64 {
	return float64(l)
}

// Eval рекурсивно вычисляет операнды и применяют к ним операцию op
func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	}
	panic(fmt.Sprintf("неподдерживаемый унарный оператор: %q", u.op))
}

// Eval рекурсивно вычисляет операнды и применяют к ним операцию op
func (b binary) Eval(env Env) float64 {
	switch b.op {
	case '+':
		return b.x.Eval(env) + b.y.Eval(env)
	case '-':
		return b.x.Eval(env) - b.y.Eval(env)
	case '*':
		return b.x.Eval(env) * b.y.Eval(env)
	case '/':
		return b.x.Eval(env) / b.y.Eval(env)
	}
	panic(fmt.Sprintf("неподдерживаемый бинарный оператор: %q", b.op))
}

// Eval вычисляет аргументы функции pow, sin, sqrt и вызывает соответствующую функцию из пакета math
func (c call) Eval(env Env) float64 {
	switch c.fn {
	case "pow":
		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	}
	panic(fmt.Sprintf("неподдерживаемый вызов функции: %s", c.fn))
}
