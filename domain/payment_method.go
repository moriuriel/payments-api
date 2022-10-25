package domain

import "time"

type PaymentMethod string

type Method struct {
	paymentMethod PaymentMethod
}

func NewMethod(paymentMethod PaymentMethod) Method {
	return Method{
		paymentMethod: paymentMethod,
	}
}

func (Payment PaymentMethod) String() string {
	return string(Payment)
}

func (m Method) CalculateAmountPaid(amount float64) (float64, int64) {
	switch m.paymentMethod {
	case "credit_card":
		return amount - (amount / 100 * 5), 5
	case "debit_card":
		return amount - (amount / 100 * 3), 3
	default:
		return amount, 0
	}
}

func (m Method) PaymentStatus() string {
	switch m.paymentMethod {
	case "credit_card":
		return "waiting_funds"
	case "debit_card":
		return "paid"
	default:
		return ""
	}
}

func (m Method) PaymentDate() time.Time {
	switch m.paymentMethod {
	case "credit_card":
		return time.Now().AddDate(0, 1, 0)
	case "debit_card":
		return time.Now()
	default:
		return time.Now()
	}
}
