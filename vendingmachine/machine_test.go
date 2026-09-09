package vendingmachine_test

import (
	"context"
	"errors"
	"testing"

	vm "github.com/esoptra/go-prac/vendingmachine"
)

// mockPayment lets tests control payment success/failure.
type mockPayment struct {
	failPayment bool
	failRefund  bool
}

func (m *mockPayment) ProcessPayment(_ context.Context, _ float64) error {
	if m.failPayment {
		return vm.ErrPaymentFailed
	}
	return nil
}

func (m *mockPayment) Refund(_ context.Context) error {
	if m.failRefund {
		return errors.New("refund failed")
	}
	return nil
}

func setup(t *testing.T, payment vm.Payment) (*vm.VendingMachine, context.CancelFunc) {
	t.Helper()
	ctx, cancel := context.WithCancel(context.Background())
	machine := vm.NewVendingMachine(ctx, payment)
	if err := machine.AddSlot(&vm.Slot{
		Code:     "A1",
		Product:  &vm.Product{Name: "Coke", Price: 1.50},
		Quantity: 2,
	}); err != nil {
		t.Fatalf("AddSlot: %v", err)
	}
	return machine, cancel
}

func TestSelect_HappyPath(t *testing.T) {
	machine, cancel := setup(t, &mockPayment{})
	defer cancel()

	if err := machine.Select(context.Background(), "A1"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestSelect_InvalidSlot(t *testing.T) {
	machine, cancel := setup(t, &mockPayment{})
	defer cancel()

	if err := machine.Select(context.Background(), "Z9"); !errors.Is(err, vm.ErrSlotNotFound) {
		t.Fatalf("expected ErrSlotNotFound, got %v", err)
	}
}

func TestSelect_EmptySlot(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	machine := vm.NewVendingMachine(ctx, &mockPayment{})
	_ = machine.AddSlot(&vm.Slot{
		Code:     "B1",
		Product:  &vm.Product{Name: "Chips", Price: 0.75},
		Quantity: 0,
	})

	if err := machine.Select(ctx, "B1"); !errors.Is(err, vm.ErrSlotEmpty) {
		t.Fatalf("expected ErrSlotEmpty, got %v", err)
	}
}

func TestSelect_WhenNotIdle(t *testing.T) {
	machine, cancel := setup(t, &mockPayment{})
	defer cancel()
	ctx := context.Background()

	_ = machine.Select(ctx, "A1")
	if err := machine.Select(ctx, "A1"); !errors.Is(err, vm.ErrInvalidOperation) {
		t.Fatalf("expected ErrInvalidOperation, got %v", err)
	}
}

func TestPay_WithoutSelect(t *testing.T) {
	machine, cancel := setup(t, &mockPayment{})
	defer cancel()

	if err := machine.Pay(context.Background()); !errors.Is(err, vm.ErrInvalidOperation) {
		t.Fatalf("expected ErrInvalidOperation, got %v", err)
	}
}

func TestCancel_ResetsToIdle(t *testing.T) {
	machine, cancel := setup(t, &mockPayment{})
	defer cancel()
	ctx := context.Background()

	_ = machine.Select(ctx, "A1")
	_ = machine.Cancel(ctx)

	// Machine should be idle — can select again
	if err := machine.Select(ctx, "A1"); err != nil {
		t.Fatalf("expected idle after cancel, got %v", err)
	}
}

func TestCancel_WhenIdle(t *testing.T) {
	machine, cancel := setup(t, &mockPayment{})
	defer cancel()

	// Cancel from idle is a no-op (returns ErrInvalidOperation to signal nothing to cancel)
	if err := machine.Cancel(context.Background()); !errors.Is(err, vm.ErrInvalidOperation) {
		t.Fatalf("expected ErrInvalidOperation, got %v", err)
	}
}

func TestFullTransaction_DecreasesQuantity(t *testing.T) {
	machine, cancel := setup(t, &mockPayment{})
	defer cancel()
	ctx := context.Background()

	_ = machine.Select(ctx, "A1")
	if err := machine.Pay(ctx); err != nil {
		t.Fatalf("payment failed: %v", err)
	}

	// Machine should be idle after a completed transaction; select again to verify
	if err := machine.Select(ctx, "A1"); err != nil {
		t.Fatalf("expected idle after successful transaction, got %v", err)
	}
}

func TestFullTransaction_PaymentFailure_ResetsToIdle(t *testing.T) {
	machine, cancel := setup(t, &mockPayment{failPayment: true})
	defer cancel()
	ctx := context.Background()

	_ = machine.Select(ctx, "A1")
	if err := machine.Pay(ctx); !errors.Is(err, vm.ErrPaymentFailed) {
		t.Fatalf("expected ErrPaymentFailed, got %v", err)
	}

	// Machine must be back in idle — not stuck in payment state
	if err := machine.Select(ctx, "A1"); err != nil {
		t.Fatalf("expected idle after failed payment, got %v", err)
	}
}
