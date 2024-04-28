package controller

import "github.com/labstack/echo/v4"

type TaxController interface {
	CalculateTax(pctx echo.Context) error
}
