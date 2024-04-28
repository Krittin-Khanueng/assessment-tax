package service

import (
	"fmt"
	"math"

	_taxModel "github.com/Krittin-Khanueng/assessment-tax/pkg/tax/model"
)

const (
	taxRate1             = 0.10
	taxRate2             = 0.15
	taxRate3             = 0.20
	taxRate4             = 0.35
	deductionDonationMax = 100000.0
	deductionPersonal    = 60000.0
	deductionKReceiptMax = 50000.0
)

type TaxServiceImpl struct{}

func NewTaxServiceImpl() TaxService {
	return &TaxServiceImpl{}
}

func (s *TaxServiceImpl) CalculateTaxRefund(userInfo *_taxModel.UserInfo) (*_taxModel.Tax, error) {
	personalDeduction := s.convertBahtToSatang(deductionPersonal)

	var taxRefund float64
	netAmount := s.convertBahtToSatang(userInfo.TotalIncome) - personalDeduction
	netAmount = netAmount - s.convertBahtToSatang(s.checkAllowanceTypeIsDonation(userInfo))
	fmt.Printf("netAmount: %f\n", netAmount)
	wat := s.convertBahtToSatang(userInfo.WHT)
	if netAmount <= s.convertBahtToSatang(150000.0) {
		taxRefund = s.calculateWht(netAmount, wat)
		taxRefund = s.convertSatangToBaht(taxRefund)
		return &_taxModel.Tax{Tax: taxRefund}, nil
	}

	if netAmount <= s.convertBahtToSatang(500000.0) {
		taxRefund = s.calculateWht(netAmount, wat)
		taxRefund = s.convertSatangToBaht(taxRefund)
		return &_taxModel.Tax{Tax: taxRefund}, nil
	}

	if netAmount <= s.convertBahtToSatang(1000000.0) {
		taxRefund = s.calculateWht(netAmount, wat)
		taxRefund = s.convertSatangToBaht(taxRefund)
		return &_taxModel.Tax{Tax: taxRefund}, nil
	}

	if netAmount <= s.convertBahtToSatang(2000000.0) {
		taxRefund = s.calculateWht(netAmount, wat)
		taxRefund = s.convertSatangToBaht(taxRefund)
		return &_taxModel.Tax{Tax: taxRefund}, nil
	}
	taxRefund = 35000.0 + 75000.0 + 200000.0 + (0.35 * (netAmount - 2000000.0))
	taxRefund = s.convertSatangToBaht(taxRefund)
	return &_taxModel.Tax{Tax: taxRefund}, nil
}

func (s *TaxServiceImpl) calculateWht(netAmount, WHT float64) float64 {
	fmt.Printf("netAmount: %f, WHT: %f\n", netAmount, WHT)
	if netAmount <= s.convertBahtToSatang(150000.0) {
		return WHT
	}
	if netAmount <= s.convertBahtToSatang(500000.0) {
		return 0.10*(netAmount-s.convertBahtToSatang(150000.0)) - WHT
	}
	if netAmount <= 10000000.0 {
		return 0.10*(5000000.0-1500000.0) + 0.15*(netAmount-5000000.0) - WHT
	}

	if netAmount <= 20000000.0 {
		return 0.10*(5000000.0-1500000.0) + 0.15*(10000000.0-5000000.0) + 0.20*(netAmount-10000000.0) - WHT
	}

	return 350000.0 + 750000.0 + 2000000.0 + (0.35 * (netAmount - 20000000.0)) - WHT
}

func (s *TaxServiceImpl) convertSatangToBaht(satang float64) float64 {
	return satang * 0.01

}

func (s *TaxServiceImpl) convertBahtToSatang(baht float64) float64 {
	return baht * 100
}

func (s *TaxServiceImpl) checkAllowanceTypeIsDonation(userInfo *_taxModel.UserInfo) float64 {
	// check allownactType is donation
	allowanceTotal := 0.0
	for _, allowance := range userInfo.Allowances {
		if allowance.AllowanceType == "donation" && allowance.Amount > deductionDonationMax {
			allowanceTotal += math.Min(allowance.Amount, deductionDonationMax)
		}
	}
	return allowanceTotal
}
