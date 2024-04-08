package errcall

import handler "github.com/Hilson-Alex/calldef/src/err_handler"

func Function[P, R any](callback func(P) (R, error)) func(P) handler.Resolver[R] {
	return func(param P) handler.Resolver[R] {
		return func() (R, error) { return callback(param) }
	}
}

func Supply[R any](callback func() (R, error)) handler.Resolver[R] {
	return callback
}

func Run(callback func() error) handler.Runner {
	return callback
}

func Consume[P any](callback func(P) error) func(P) handler.Runner {
	return func(param P) handler.Runner {
		return func() error { return callback(param) }
	}
}
