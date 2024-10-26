package salarycalculations

import (
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/amorimluiz/work_benefits_api/internal/utils"
	"github.com/joho/godotenv"
)

func newService() *salaryCalculationsServiceStruct {
	service := salaryCalculationsServiceStruct{}
	loadBracketsAndRates(&service)
	return &service
}

func loadBracketsAndRates(s *salaryCalculationsServiceStruct) {
	_ = godotenv.Load()

	var envKeys = []string{"INSS_BRACKETS", "INSS_RATES", "IRRF_BRACKETS", "IRRF_RATES", "IRRF_DEDUCTIONS"}
	resultChannels := make([]chan []float64, len(envKeys))

	for i, envKey := range envKeys {
		resultChannels[i] = make(chan []float64)
		go s.parseBracketsOrRatesFromEnv(envKey, resultChannels[i])
	}

	s.INSSBrackets = <-resultChannels[0]
	s.INSSRates = <-resultChannels[1]
	s.IRRFBrackets = <-resultChannels[2]
	s.IRRFRates = <-resultChannels[3]
	s.IRRFDeductions = <-resultChannels[4]

	s.DeductionPerDependent, _ = strconv.ParseFloat(os.Getenv("DEDUCTION_PER_DEPENDENT"), 64)
}

var salaryCalculationsService = newService()

func (s *salaryCalculationsServiceStruct) parseBracketsOrRatesFromEnv(envKey string, succeededChannel chan []float64) {
	data := strings.Split(os.Getenv(envKey), ";")

	parsedData := make([]float64, len(data))

	for i, bracket := range data {
		parsedBracket, err := strconv.ParseFloat(bracket, 64)

		if err != nil {
			succeededChannel <- nil
		}

		parsedData[i] = parsedBracket
	}

	succeededChannel <- parsedData
}

func (s *salaryCalculationsServiceStruct) calculateINSSTax(grossSalary float64) float64 {
	inssDiscount := 0.0
	previousBracket := 0.0

	for i := 0; i < len(s.INSSBrackets); i++ {
		if grossSalary <= s.INSSBrackets[i] {
			inssDiscount += (grossSalary - previousBracket) * s.INSSRates[i]
			break
		}

		inssDiscount += (s.INSSBrackets[i] - previousBracket) * s.INSSRates[i]
		previousBracket = s.INSSBrackets[i]
	}

	return utils.EnsureFloatPrecision(inssDiscount, 2)
}

func (s *salaryCalculationsServiceStruct) calculateIRRFTax(irrfCalculationBracket float64) float64 {
	irrfBracketIndex := slices.IndexFunc(s.IRRFBrackets, func(bracket float64) bool {
		return irrfCalculationBracket <= bracket || slices.Max(s.IRRFBrackets) == bracket
	})

	irrfTax := irrfCalculationBracket*s.IRRFRates[irrfBracketIndex] - s.IRRFDeductions[irrfBracketIndex]

	return utils.EnsureFloatPrecision(irrfTax, 2)
}

func (s *salaryCalculationsServiceStruct) calculateNetSalaryAndTaxes(body netSalaryCalculationBody) (*netSalaryCalculationResponse, error) {
	inssTax := s.calculateINSSTax(body.GrossSalary)

	irrfCalculationBracket := body.GrossSalary - inssTax - s.DeductionPerDependent*float64(body.Dependents)

	irrfTax := s.calculateIRRFTax(irrfCalculationBracket)

	netSalary := body.GrossSalary - inssTax - irrfTax - body.OtherDiscounts

	netSalary = utils.EnsureFloatPrecision(netSalary, 2)

	return &netSalaryCalculationResponse{
		INSSTax:   inssTax,
		IRRFTax:   irrfTax,
		NetSalary: netSalary,
	}, nil
}
