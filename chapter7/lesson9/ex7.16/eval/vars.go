package eval

func (v Var) Vars() []Var {
	return []Var{v}
}

func (l literal) Vars() []Var {
	return nil
}

func (u unary) Vars() []Var {
	return u.x.Vars()
}

func (b binary) Vars() []Var {
	return append(b.x.Vars(), b.y.Vars()...)
}

func (c call) Vars() []Var {
	var vars []Var
	for _, e := range c.args {
		vars = append(vars, e.Vars()...)
	}
	return vars
}

func (p postUnary) Vars() []Var {
	return p.x.Vars()
}
