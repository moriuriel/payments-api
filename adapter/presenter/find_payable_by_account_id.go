package presenter

import "github.com/moriuriel/go-payments/usecase"

type findPayableByAccountIDPresenter struct{}

func NewFindTotalPayableByAccountIDPresenter() usecase.FindTotalPayableByAccountIDPresenter {
	return findPayableByAccountIDPresenter{}
}

func (c findPayableByAccountIDPresenter) Output(totalPaid float64, totalWaitingFunds float64) usecase.FindTotalPayableByAccountIDOutput {
	return usecase.FindTotalPayableByAccountIDOutput{
		TotalPaid:         totalPaid,
		TotalWaitingFunds: totalWaitingFunds,
	}
}
