package main

import "testing"

func TestPaymentStatusString(t *testing.T) {
	tests := []struct {
		input PaymentStatus
		want  string
	}{
		{Pending, "Pending"},
		{Paid, "Paid"},
		{Failed, "Failed"},
		{Refunded, "Refunded"},
		{PaymentStatus(99), "Unknown"},
	}

	for _, tt := range tests {
		got := tt.input.String()
		if got != tt.want {
			t.Fatalf("got %s want %s", got, tt.want)
		}
	}
}
