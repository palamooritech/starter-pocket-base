package typos

import "github.com/labstack/echo/v5"

type PBPathFunc func(c echo.Context) error

type APIError struct {
	Error string `json:"error"`
}
