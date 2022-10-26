package presenter

import (
	"math"

	"github.com/moriuriel/go-payments/usecase"
)

type findPayableByAccountIDPresenter struct{}

func NewFindTotalPayableByAccountIDPresenter() usecase.FindTotalPayableByAccountIDPresenter {
	return findPayableByAccountIDPresenter{}
}

func (c findPayableByAccountIDPresenter) Output(totalPaid float64, totalWaitingFunds float64) usecase.FindTotalPayableByAccountIDOutput {
	return usecase.FindTotalPayableByAccountIDOutput{
		TotalPaid:         toFixed(totalPaid, 2),
		TotalWaitingFunds: toFixed(totalWaitingFunds, 2),
	}
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}
