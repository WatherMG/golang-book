package eval

// Expr представляет арифметическое выражение
type Expr interface {
	// Eval возвращает значение данного Expr в среде env.
	Eval(env Env) float64
	// Check сообщает об ошибках в данном Expr и добавляет свои Vars.
	Check(vars map[Var]bool) error
	String() string
	Vars() []Var
}

// Var определяет переменную, например x.
type Var string

// Представляет собой числовую константу, например 3.141.
type literal float64

// Представляет выражение с унарным оператором, например -x.
type unary struct {
	op rune // + или -
	x  Expr
}

// Представляет выражение с бинарным оператором, например x+y.
type binary struct {
	op   rune // +, -, *, /
	x, y Expr
}

// Представляет выражение вызова функции, например sin(x).
type call struct {
	fn   string // одно из pow, sin, sqrt
	args []Expr
}

// Представляет собой выражение с факториалом, например 4!
type postUnary struct {
	op rune // !
	x  Expr
}
