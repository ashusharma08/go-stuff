package parkinglot

import (
	"context"
	"time"
)

type payRequest struct {
	price    float64
	ticketID string
	callback func(string) *response
}

type paymentManager interface {
	processPayment(context.Context, *payRequest)
}

type mockPaymentManager struct {
}

func (m *mockPaymentManager) processPayment(ctx context.Context, p *payRequest) {
	//whatever happens in payments
	time.Sleep(2 * time.Second)
	p.callback(p.ticketID)
}
