package vendingmachine

import "context"

// Payment is the interface for processing transactions.
// Inject a concrete implementation via NewVendingMachine.
type Payment interface {
	ProcessPayment(ctx context.Context, amount float64) error
	Refund(ctx context.Context) error
}

type Cash struct{}

func (c *Cash) ProcessPayment(_ context.Context, _ float64) error { return nil }
func (c *Cash) Refund(_ context.Context) error                    { return nil }

type CreditCard struct{}

func (c *CreditCard) ProcessPayment(_ context.Context, _ float64) error { return nil }
func (c *CreditCard) Refund(_ context.Context) error                    { return nil }

type UPI struct{}

func (u *UPI) ProcessPayment(_ context.Context, _ float64) error { return nil }
func (u *UPI) Refund(_ context.Context) error                    { return nil }
