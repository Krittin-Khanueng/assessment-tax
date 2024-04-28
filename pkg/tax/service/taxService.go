package service

import (
	_taxModel "github.com/Krittin-Khanueng/assessment-tax/pkg/tax/model"
)

type TaxService interface {
	CalculateTaxRefund(userInfo *_taxModel.UserInfo) (*_taxModel.Tax, error)
}
