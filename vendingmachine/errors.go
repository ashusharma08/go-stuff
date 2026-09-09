package vendingmachine

import "errors"

var (
	ErrInvalidOperation = errors.New("operation not valid in current state")
	ErrSlotNotFound     = errors.New("slot not found")
	ErrSlotEmpty        = errors.New("slot is empty")
	ErrPaymentFailed    = errors.New("payment failed")
	ErrDispenseFailed   = errors.New("dispense failed")
)
