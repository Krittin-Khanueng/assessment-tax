package service

import (
	"math"

	_taxModel "github.com/Krittin-Khanueng/assessment-tax/pkg/tax/model"
)

type TaxServiceImpl struct{}

func NewTaxServiceImpl() TaxService {
	return &TaxServiceImpl{}
}

func (s *TaxServiceImpl) CalculateTaxRefund(userInfo *_taxModel.UserInfo) (*_taxModel.Tax, error) {
	personalDeduction := 60000.0

	var taxRefund float64
	netAmount := userInfo.TotalIncome - personalDeduction
	if netAmount <= 150000.0 {
		taxRefund = 0.0
		taxRefund = math.Round(taxRefund*100) / 100
		return &_taxModel.Tax{Tax: taxRefund}, nil
	}

	if netAmount <= 500000.0 {
		taxRefund = 0.10 * (netAmount - 150000.0)
		taxRefund = math.Round(taxRefund*100) / 100
		return &_taxModel.Tax{Tax: taxRefund}, nil
	}

	if netAmount <= 1000000.0 {
		taxRefund = 35000.0 + (0.15 * (netAmount - 500000.0))
		taxRefund = math.Round(taxRefund*100) / 100
		return &_taxModel.Tax{Tax: taxRefund}, nil
	}

	if netAmount <= 2000000.0 {
		taxRefund = 35000.0 + 75000.0 + (0.20 * (netAmount - 1000000.0))
		taxRefund = math.Round(taxRefund*100) / 100
		return &_taxModel.Tax{Tax: taxRefund}, nil
	}
	taxRefund = 35000.0 + 75000.0 + 200000.0 + (0.35 * (netAmount - 2000000.0))
	taxRefund = math.Round(taxRefund*100) / 100
	return &_taxModel.Tax{Tax: taxRefund}, nil
}
