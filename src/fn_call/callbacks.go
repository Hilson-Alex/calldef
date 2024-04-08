package fncall

func Function[P, R any](callback func(P) R) func(P) func() R {
	return func(param P) func() R {
		return func() R { return callback(param) }
	}
}

func Supply[R any](resolver func() R) func() R {
	return resolver
}

func Run(runner func()) func() {
	return runner
}

func Consume[P any](callback func(P)) func(P) func() {
	return func(param P) func() {
		return func() { callback(param) }
	}
}
