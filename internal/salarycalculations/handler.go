package salarycalculations

import "github.com/gin-gonic/gin"

func newHandler() *salaryCalculationsHandler {
	return &salaryCalculationsHandler{service: salaryCalculationsService}
}

var SalaryCalculationsHandler = newHandler()

func (h *salaryCalculationsHandler) CalculateNetSalary(ctx *gin.Context) {
	var salaryCalculationData netSalaryCalculationBody

	err := ctx.ShouldBindJSON(&salaryCalculationData)

	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	netSalaryCalculations, err := h.service.calculateNetSalaryAndTaxes(salaryCalculationData)

	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, netSalaryCalculations)
}
