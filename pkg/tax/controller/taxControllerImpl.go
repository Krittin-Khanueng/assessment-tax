package controller

import (
	"net/http"

	_taxService "github.com/Krittin-Khanueng/assessment-tax/pkg/tax/service"
	"github.com/labstack/echo/v4"

	_taxModel "github.com/Krittin-Khanueng/assessment-tax/pkg/tax/model"
)

type TaxControllerImpl struct {
	taxService _taxService.TaxService
}

func NewTaxControllerImpl(taxService _taxService.TaxService) TaxController {
	return &TaxControllerImpl{taxService}
}

func (c *TaxControllerImpl) CalculateTax(pctx echo.Context) error {
	userInfo := new(_taxModel.UserInfo)
	if err := pctx.Bind(userInfo); err != nil {
		return err
	}
	tax, err := c.taxService.CalculateTaxRefund(userInfo)
	if err != nil {
		return err
	}
	return pctx.JSON(http.StatusOK, tax)
}
