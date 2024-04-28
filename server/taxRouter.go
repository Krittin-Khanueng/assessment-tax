package server

import (
	_taxController "github.com/Krittin-Khanueng/assessment-tax/pkg/tax/controller"
	_taxService "github.com/Krittin-Khanueng/assessment-tax/pkg/tax/service"
)

func (s *echoServer) initTaxRouter() {
	router := s.app.Group("/v1/tax")

	taxService := _taxService.NewTaxServiceImpl()
	taxController := _taxController.NewTaxControllerImpl(taxService)

	router.POST("/calculate", taxController.CalculateTax)

}
