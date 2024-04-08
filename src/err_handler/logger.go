package errhandler

import "log"

type ErrLogger interface {
	Panicln(...any)
	Println(...any)
	Fatalln(...any)
}

var defaultLogger ErrLogger = log.Default()

func SetDefaultLogger(logger ErrLogger) {
	defaultLogger = logger
}
