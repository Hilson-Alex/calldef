package errhandler_test

import (
	errhandler "github.com/Hilson-Alex/calldef/src/err_handler"
)

type logStatus int

const (
	none logStatus = iota
	panicked
	quitted
	logged
)

type testLogger struct {
	status logStatus
}

func (log *testLogger) Panicln(_ ...any) {
	log.status = panicked
}

func (log *testLogger) Println(_ ...any) {
	log.status = logged
}

func (log *testLogger) Fatalln(_ ...any) {
	log.status = quitted
}

var logger testLogger

func setupLogger() {
	logger = testLogger{}
	errhandler.SetDefaultLogger(&logger)
}
