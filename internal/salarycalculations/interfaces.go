package salarycalculations

type netSalaryCalculationBody struct {
	GrossSalary    float64 `json:"gross_salary"`
	Dependents     int     `json:"dependents"`
	OtherDiscounts float64 `json:"other_discounts"`
}

type netSalaryCalculationResponse struct {
	INSSTax   float64 `json:"inss_tax"`
	IRRFTax   float64 `json:"irrf_tax"`
	NetSalary float64 `json:"net_salary"`
}

type salaryCalculationsServiceStruct struct {
	INSSBrackets          []float64
	INSSRates             []float64
	IRRFBrackets          []float64
	IRRFRates             []float64
	IRRFDeductions        []float64
	DeductionPerDependent float64
}

type salaryCalculationsHandler struct {
	service *salaryCalculationsServiceStruct
}
