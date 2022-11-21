package echspekt

import (
	"reflect"
	"testing"

	"github.com/labstack/echo/v4"
)

var (
	format = "wanted result to be \n%+v\n but got \n%+v\n"
)

func TestToHandler(t *testing.T) {
	testRouter := echo.New()
	stubEchoHandler := func(c echo.Context) error {
		return c.String(200, "hello")
	}

	t.Run("When both params are passed", func(t *testing.T) {
		when := WhenIMakeARequest(ToHandler(testRouter, stubEchoHandler))

		if !reflect.DeepEqual(testRouter, when.router) {
			t.Errorf(format, testRouter, when.router)
			return
		}

		if when.handler == nil {
			t.Errorf(format, stubEchoHandler, when.handler)
		}
	})

	t.Run("When router param has nil value", func(t *testing.T) {
		when := WhenIMakeARequest(ToHandler(nil, stubEchoHandler))

		if when.router == nil {
			t.Errorf("wanted default router but got nil")
			return
		}

		if when.handler == nil {
			t.Errorf(format, stubEchoHandler, when.handler)
		}
	})

	t.Run("When handler param has nil value", func(t *testing.T) {
		when := WhenIMakeARequest(ToHandler(testRouter, nil))

		if when.router == nil {
			t.Errorf(format, testRouter, when.router)
			return
		}

		if when.handler == nil {
			t.Errorf("wanted default router but got nil")
		}
	})
}

func TestWithMethod(t *testing.T) {
	t.Run("when the method param is legit it should return that method", func(t *testing.T) {
		want := "GET"
		when := WhenIMakeARequest(WithMethod(want))

		if when.method != want {
			t.Errorf("wanted %s but got %s\n", want, when.method)
			return
		}
	})

	t.Run("when the method param is empty it should default to GET", func(t *testing.T) {
		want := "GET"
		when := WhenIMakeARequest(WithMethod(""))

		if when.method != want {
			t.Errorf("wanted %s but got %s\n", want, when.method)
			return
		}
	})
}
