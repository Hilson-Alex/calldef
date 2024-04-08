package errhandler

type Runner func() error

func (runner Runner) OrPanicOn(logger ErrLogger) {

	if err := runner(); err != nil {
		logger.Panicln(err)
	}
}

func (runner Runner) OrPanic() {
	runner.OrPanicOn(defaultLogger)
}

func (runner Runner) OrQuitOn(logger ErrLogger) {
	if err := runner(); err != nil {
		logger.Fatalln(err)
	}
}

func (runner Runner) OrQuit() {
	runner.OrQuitOn(defaultLogger)
}

func (runner Runner) OrLogOn(logger ErrLogger) {
	if err := runner(); err != nil {
		logger.Println(err)
	}
}

func (runner Runner) OrLog() {
	runner.OrLogOn(defaultLogger)
}

func (runner Runner) OrHandle(handler func(err error)) {
	if err := runner(); err != nil {
		handler(err)
	}
}

func (runner Runner) OrIgnore() {
	var _ = runner()
}
