package simple

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func ExceptionMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := next(c); err != nil {
			switch e := err.(type) {
			case *Error:
				return c.JSON(e.HttpCode, NewErrorResponse(e, e.Message, e.ErrorCode))
			case *echo.HTTPError:
				return c.JSON(e.Code, NewErrorResponse(e, e.Message.(string), e.Code))
			default:
				return c.JSON(http.StatusInternalServerError, NewErrorResponse(e, e.Error(), 1000))
			}
		}

		return nil
	}
}
