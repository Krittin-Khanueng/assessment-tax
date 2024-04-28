package service

import (
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

// การคำนวนภาษีตามขั้นบันใด
// รายได้ 0 - 150,000 ได้รับการยกเว้น
// 150,001 - 500,000 อัตราภาษี 10%
// 500,001 - 1,000,000 อัตราภาษี 15%
// 1,000,001 - 2,000,000 อัตราภาษี 20%
// มากกว่า 2,000,000 อัตราภาษี 35%

func (s *TaxServiceImpl) CalculateTaxRefund(userInfo *_taxModel.UserInfo) (*_taxModel.Tax, error) {

	personalDeduction := s.convertBahtToSatang(deductionPersonal)
	txtLevel := make([]_taxModel.TaxLevel, 5)
	netAmount := s.convertBahtToSatang(userInfo.TotalIncome) - personalDeduction
	netAmount -= s.convertBahtToSatang(s.calculateAllowanceTypeIsDonation(userInfo))
	wat := s.convertBahtToSatang(userInfo.WHT)

	var textResult float64

	// Calculate tax refund for each tax level
	if netAmount > s.convertBahtToSatang(150000.0) {
		textResult += 0.0
		txtLevel[0].Tax = 0.0
		netAmount -= s.convertBahtToSatang(150000.0)
	} else {
		textResult += 0.0
		txtLevel[0].Tax = 0.0
		netAmount = 0.0
	}

	if netAmount > s.convertBahtToSatang(350000.0) {
		textResult += taxRate1 * s.convertBahtToSatang(350000.0)
		txtLevel[1].Tax = taxRate1 * s.convertBahtToSatang(350000.0)
		netAmount -= s.convertBahtToSatang(350000.0)
	} else {
		textResult += taxRate1 * netAmount
		txtLevel[1].Tax = taxRate1 * netAmount
		netAmount = 0.0
	}

	if netAmount > s.convertBahtToSatang(500000.0) {
		textResult += taxRate2 * s.convertBahtToSatang(500000.0)
		txtLevel[2].Tax = taxRate2 * s.convertBahtToSatang(500000.0)
		netAmount -= s.convertBahtToSatang(500000.0)
	} else {
		textResult += taxRate2 * netAmount
		txtLevel[2].Tax = taxRate2 * netAmount
		netAmount = 0.0
	}

	if netAmount > s.convertBahtToSatang(1000000.0) {
		textResult += taxRate2 * s.convertBahtToSatang(1000000.0)
		txtLevel[3].Tax = taxRate2 * s.convertBahtToSatang(1000000.0)
		netAmount -= s.convertBahtToSatang(1000000.0)
	} else {
		textResult += taxRate2 * netAmount
		txtLevel[3].Tax = taxRate2 * netAmount
		netAmount = 0.0
	}

	if netAmount > 0.0 {
		textResult += taxRate3 * netAmount
		txtLevel[4].Tax = taxRate3 * netAmount
		netAmount -= 0.0
	}

	for t, v := range txtLevel {
		txtLevel[t].Level = s.getTaxLevelRange(t)
		txtLevel[t].Tax = s.convertSatangToBaht(v.Tax)
	}
	textResult -= wat
	tax := &_taxModel.Tax{
		Tax:      s.convertSatangToBaht(textResult),
		TaxLevel: txtLevel,
	}

	return tax, nil
}

func (s *TaxServiceImpl) convertSatangToBaht(satang float64) float64 {
	return satang * 0.01

}

func (s *TaxServiceImpl) convertBahtToSatang(baht float64) float64 {
	return baht * 100
}

func (s *TaxServiceImpl) calculateAllowanceTypeIsDonation(userInfo *_taxModel.UserInfo) float64 {
	// check allownactType is donation
	allowanceTotal := 0.0
	for _, allowance := range userInfo.Allowances {
		if allowance.AllowanceType == "donation" && allowance.Amount > deductionDonationMax {
			allowanceTotal += math.Min(allowance.Amount, deductionDonationMax)
		}
	}
	return allowanceTotal
}

func (s *TaxServiceImpl) getTaxLevelRange(level int) string {
	switch level {
	case 0:
		return "0-150,000"
	case 1:
		return "150,001-500,000"
	case 2:
		return "500,001-1,000,000"
	case 3:
		return "1,000,001-2,000,000"
	case 4:
		return "2,000,001 ขึ้นไป"
	default:
		return ""
	}
}
