package main

type PaymentStatus int

const (
	Pending PaymentStatus = iota
	Paid
	Failed
	Refunded
)

func (p PaymentStatus) String() string {
	switch p {
	case Pending:
		return "Pending"
	case Paid:
		return "Paid"
	case Failed:
		return "Failed"
	case Refunded:
		return "Refunded"
	default:
		return "Unknown"
	}
}
