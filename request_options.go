package echspekt

import "github.com/labstack/echo/v4"

func ToHandler(router *echo.Echo, handler echo.HandlerFunc) RequestOption {
	if router == nil {
		router = echo.New()
	}

	if handler == nil {
		handler = func(c echo.Context) error {
			return c.String(200, "OK")
		}
	}

	return func(w *When) {
		w.router = router
		w.handler = handler
	}
}

func WithMethod(method string) RequestOption {
	if method == "" {
		method = "GET"
	}

	return func(w *When) {
		w.method = method
	}
}
