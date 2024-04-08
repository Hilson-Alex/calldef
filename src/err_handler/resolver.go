package errhandler

type Resolver[T any] func() (T, error)

func (resolver Resolver[T]) OrPanicOn(logger ErrLogger) T {
	var result, err = resolver()
	if err != nil {
		logger.Panicln(err)
	}
	return result
}

func (resolver Resolver[T]) OrPanic() T {
	return resolver.OrPanicOn(defaultLogger)
}

func (resolver Resolver[T]) OrQuitOn(logger ErrLogger) T {
	var result, err = resolver()
	if err != nil {
		logger.Fatalln(err)
	}
	return result
}

func (resolver Resolver[T]) OrQuit() T {
	return resolver.OrQuitOn(defaultLogger)
}

func (resolver Resolver[T]) OrLogOn(logger ErrLogger) T {
	var result, err = resolver()
	if err != nil {
		logger.Println(err)
	}
	return result
}

func (resolver Resolver[T]) OrLog() T {
	return resolver.OrLogOn(defaultLogger)
}

func (resolver Resolver[T]) OrHandle(handler func(err error)) T {
	var result, err = resolver()
	if err != nil {
		handler(err)
	}
	return result
}

func (resolver Resolver[T]) OrIgnore() T {
	var result, _ = resolver()
	return result
}

func (resolver Resolver[T]) WithDefault(defaultValue T) T {
	var result, err = resolver()
	if err != nil {
		return defaultValue
	}
	return result
}
