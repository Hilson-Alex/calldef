package errhandler_test

import (
	"errors"
	"testing"

	errcall "github.com/Hilson-Alex/calldef/src/err_call"
)

func TestResolver_OrPanic(t *testing.T) {
	setupLogger()

	t.Run("Test OrPanic return", func(t *testing.T) {
		var result = errcall.
			Supply(func() (string, error) { return "value", nil }).
			OrPanic()
		if result != "value" || logger.status != none {
			t.Error("Resolver didn't returned!")
		}
	})

	t.Run("Test OrPanic failing", func(t *testing.T) {
		var _ = errcall.
			Supply(func() (interface{}, error) { return nil, errors.New("error") }).
			OrPanic()

		if logger.status != panicked {
			t.Error("Resolver didn't panicked!")
		}
	})

}

func TestResolver_OrQuit(t *testing.T) {
	setupLogger()

	t.Run("Test OrQuit return", func(t *testing.T) {
		var result = errcall.
			Supply(func() (string, error) { return "value", nil }).
			OrQuit()
		if result != "value" || logger.status != none {
			t.Error("Resolver didn't returned!")
		}
	})

	t.Run("Test OrQuit failing", func(t *testing.T) {
		var _ = errcall.
			Supply(func() (interface{}, error) { return nil, errors.New("error") }).
			OrQuit()

		if logger.status != quitted {
			t.Error("Resolver didn't quitted!")
		}
	})
}

func TestResolver_OrLog(t *testing.T) {
	setupLogger()

	t.Run("Test OrLog return", func(t *testing.T) {
		var result = errcall.
			Supply(func() (string, error) { return "value", nil }).
			OrLog()
		if result != "value" || logger.status != none {
			t.Error("Resolver didn't returned!")
		}
	})

	t.Run("Test OrLog failing", func(t *testing.T) {
		var _ = errcall.
			Supply(func() (interface{}, error) { return nil, errors.New("error") }).
			OrLog()

		if logger.status != logged {
			t.Error("Resolver didn't logged!")
		}
	})
}

func TestResolver_OrHandle(t *testing.T) {
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
		var result = errcall.
			Supply(func() (string, error) { return "value", nil }).
			OrHandle(handler)
		if result != "value" {
			t.Error("Resolver didn't returned!")
		}
		if wasHandled() {
			t.Error("The handle function has been called, when it shouldn't")
		}
	})

	t.Run("Test OrHandle failing", func(t *testing.T) {
		var handler, wasHandled = setupHandler()
		var _ = errcall.
			Supply(func() (interface{}, error) { return nil, errors.New("error") }).
			OrHandle(handler)
		if !wasHandled() {
			t.Error("The handle function wasn't called")
		}
	})
}

func TestResolver_OrIgnore(t *testing.T) {
	var value = errcall.
		Supply(func() (interface{}, error) { return nil, errors.New("error") }).
		OrIgnore()

	if value != nil {
		t.Error("The OrIgnore method returned an unexpected value")
	}
}

func TestResolver_WithDefault(t *testing.T) {
	t.Run("Test WithDefault return", func(t *testing.T) {
		var result = errcall.
			Supply(func() (string, error) { return "value", nil }).
			WithDefault("default")
		if result != "value" {
			t.Errorf("Expected %v, but %v was received instead!", "value", result)
		}
	})

	t.Run("Test WithDefault failing", func(t *testing.T) {
		var result = errcall.
			Supply(func() (string, error) { return "value", errors.New("error") }).
			WithDefault("default")
		if result != "default" {
			t.Errorf("Expected %v, but %v was received instead!", "default", result)
		}
	})
}
