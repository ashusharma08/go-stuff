package vendingmachine

import "context"

type requestType int

const (
	reqAddSlot requestType = iota
	reqSelect
	reqPay
	reqDispensing // internal: transition to dispensing state
	reqCancel
	reqComplete
)

type request struct {
	typ      requestType
	ctx      context.Context // used by reqPay for the payment goroutine
	slotCode string
	slot     *Slot
	resp     chan error
}

type VendingMachine struct {
	requestChan  chan request
	inventory    map[string]*Slot
	selectedSlot *Slot
	state        state
	payment      Payment
}

func NewVendingMachine(ctx context.Context, payment Payment) *VendingMachine {
	vm := &VendingMachine{
		requestChan: make(chan request),
		inventory:   make(map[string]*Slot),
		state:       stateIdle,
		payment:     payment,
	}
	go vm.loop(ctx)
	return vm
}

// AddSlot stocks the machine with a slot. Safe to call concurrently.
func (v *VendingMachine) AddSlot(slot *Slot) error {
	resp := make(chan error, 1)
	v.requestChan <- request{typ: reqAddSlot, slot: slot, resp: resp}
	return <-resp
}

// Select chooses a product by slot code. Must be called from the idle state.
func (v *VendingMachine) Select(ctx context.Context, code string) error {
	resp := make(chan error, 1)
	v.requestChan <- request{typ: reqSelect, slotCode: code, resp: resp}
	return <-resp
}

// Pay initiates and waits for the full payment + dispense cycle to complete.
// Blocks until the transaction succeeds, fails, or ctx is cancelled.
func (v *VendingMachine) Pay(ctx context.Context) error {
	resp := make(chan error, 1)
	v.requestChan <- request{typ: reqPay, ctx: ctx, resp: resp}
	return <-resp
}

// Cancel aborts the current selection and returns to idle.
func (v *VendingMachine) Cancel(ctx context.Context) error {
	resp := make(chan error, 1)
	v.requestChan <- request{typ: reqCancel, resp: resp}
	return <-resp
}

// loop is the single goroutine that owns all state mutations.
// All public methods communicate with it via requestChan.
func (v *VendingMachine) loop(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case req, ok := <-v.requestChan:
			if !ok {
				return
			}
			switch req.typ {
			case reqAddSlot:
				req.resp <- v.handleAddSlot(req.slot)
			case reqSelect:
				req.resp <- v.handleSelect(req.slotCode)
			case reqPay:
				// handlePay spawns a goroutine and does NOT send to req.resp immediately.
				// The goroutine is responsible for sending the final result.
				v.handlePay(req)
			case reqDispensing:
				req.resp <- v.handleDispensing()
			case reqCancel:
				req.resp <- v.handleCancel()
			case reqComplete:
				req.resp <- v.handleComplete()
			}
		}
	}
}

func (v *VendingMachine) handleAddSlot(slot *Slot) error {
	if slot == nil || slot.Code == "" || slot.Product == nil {
		return ErrSlotNotFound
	}
	v.inventory[slot.Code] = slot
	return nil
}

func (v *VendingMachine) handleSelect(code string) error {
	if v.state != stateIdle {
		return ErrInvalidOperation
	}
	slot, ok := v.inventory[code]
	if !ok {
		return ErrSlotNotFound
	}
	if slot.Quantity == 0 {
		return ErrSlotEmpty
	}
	v.selectedSlot = slot
	v.state = stateSelected
	return nil
}

// handlePay transitions to statePayment and spawns a goroutine to do I/O.
// It does NOT write to req.resp — processPayment does that when finished.
func (v *VendingMachine) handlePay(req request) {
	if v.state != stateSelected {
		req.resp <- ErrInvalidOperation
		return
	}
	slot := v.selectedSlot
	v.state = statePayment
	go v.processPayment(req.ctx, slot, req.resp)
}

// processPayment runs outside the loop goroutine to avoid blocking it during I/O.
// State changes are routed back through requestChan to preserve single ownership.
func (v *VendingMachine) processPayment(ctx context.Context, slot *Slot, resp chan<- error) {
	if err := v.payment.ProcessPayment(ctx, slot.Product.Price); err != nil {
		v.sendInternal(reqCancel)
		resp <- err
		return
	}

	v.sendInternal(reqDispensing)
	if err := v.dispense(); err != nil {
		_ = v.payment.Refund(ctx)
		v.sendInternal(reqCancel)
		resp <- err
		return
	}

	v.sendInternal(reqComplete)
	resp <- nil
}

// sendInternal sends an event back to the loop and waits for it to be processed.
func (v *VendingMachine) sendInternal(typ requestType) {
	done := make(chan error, 1)
	v.requestChan <- request{typ: typ, resp: done}
	<-done
}

func (v *VendingMachine) handleDispensing() error {
	v.state = stateDispensing
	return nil
}

func (v *VendingMachine) handleCancel() error {
	v.reset()
	return nil
}

func (v *VendingMachine) handleComplete() error {
	if v.selectedSlot != nil {
		v.selectedSlot.Quantity--
	}
	v.reset()
	return nil
}

func (v *VendingMachine) dispense() error {
	// TODO: actuate physical dispensing mechanism
	return nil
}

func (v *VendingMachine) reset() {
	v.selectedSlot = nil
	v.state = stateIdle
}
