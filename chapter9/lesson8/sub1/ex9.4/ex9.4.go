package pipeline

func pipeline(stages int) (in, out chan interface{}) {
	if stages < 1 {
		return nil, nil
	}
	in = make(chan interface{})
	out = in
	for i := 0; i < stages; i++ {
		prev := out
		next := make(chan interface{})
		go func(prev, next chan interface{}) {
			for v := range prev {
				next <- v
			}
			close(next)
		}(prev, next)
		out = next
	}
	return in, out
}
