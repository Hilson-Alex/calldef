package errhandler_test

import (
	"errors"
	"testing"

	errcall "github.com/Hilson-Alex/calldef/src/err_call"
)

func TestRunner_OrPanic(t *testing.T) {
	setupLogger()

	t.Run("Test OrPanic run", func(t *testing.T) {
		errcall.
			Run(func() error { return nil }).
			OrPanic()
		if logger.status != none {
			t.Error("Runner didn't runned as expected!")
		}
	})

	t.Run("Test OrPanic failing", func(t *testing.T) {
		errcall.
			Run(func() error { return errors.New("error") }).
			OrPanic()

		if logger.status != panicked {
			t.Error("Runner didn't panicked!")
		}
	})

}

func TestRunner_OrQuit(t *testing.T) {
	setupLogger()

	t.Run("Test OrQuit run", func(t *testing.T) {
		errcall.
			Run(func() error { return nil }).
			OrQuit()
		if logger.status != none {
			t.Error("Runner didn't runned as expected!")
		}
	})

	t.Run("Test OrQuit failing", func(t *testing.T) {
		errcall.
			Run(func() error { return errors.New("error") }).
			OrQuit()

		if logger.status != quitted {
			t.Error("Runner didn't quitted!")
		}
	})
}

func TestRunner_OrLog(t *testing.T) {
	setupLogger()

	t.Run("Test OrLog run", func(t *testing.T) {
		errcall.
			Run(func() error { return nil }).
			OrLog()
		if logger.status != none {
			t.Error("Runner didn't runned as expected!")
		}
	})

	t.Run("Test OrLog failing", func(t *testing.T) {
		errcall.
			Run(func() error { return errors.New("error") }).
			OrLog()

		if logger.status != logged {
			t.Error("Runner didn't logged!")
		}
	})
}

func TestRunner_OrHandle(t *testing.T) {
	setupLogger()
	var setupHandler = func() (func(error), func() bool) {
		var handled = false
		return func(err error) {
				if err != nil {
					handled = true
				}
			}, func() bool {
				return handled
			}
	}
	t.Run("Test OrHandle return", func(t *testing.T) {
		var handler, wasHandled = setupHandler()
		errcall.
			Run(func() error { return nil }).
			OrHandle(handler)
		if wasHandled() {
			t.Error("The handle function has been called, when it shouldn't")
		}
	})

	t.Run("Test OrHandle failing", func(t *testing.T) {
		var handler, wasHandled = setupHandler()
		errcall.
			Run(func() error { return errors.New("error") }).
			OrHandle(handler)
		if !wasHandled() {
			t.Error("The handle function wasn't called")
		}
	})
}

func TestRunner_OrIgnore(t *testing.T) {
	setupLogger()
	errcall.
		Run(func() error { return errors.New("error") }).
		OrIgnore()

	if logger.status != none {
		t.Error("The OrIgnore method returned an unexpected value")
	}
}
