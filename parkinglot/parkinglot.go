package parkinglot

import (
	"context"
	"log/slog"
	"os"
	"time"

	"github.com/esoptra/go-prac/parkinglot/store"
)

type requesttype int

const (
	enter requesttype = iota
	exit
	payment
	//vacate???
)

type enterRequest struct {
	registrationNumber string
	ticketType         store.TicketType
	occupancyType      OccupancyType
}
type exitRequest struct {
	ticketNumber string
}
type request struct {
	typ  requesttype
	resp chan *response
	enterRequest
	exitRequest
}

type response struct {
	ticketID string
	err      error
}

type ParkingLot struct {
	request         chan request
	pricingStrategy pricingStrategy
	paymentManager  paymentManager
	ticketRepo      store.TicketRepo
	slotRepo        store.SlotRepo
}

type OccupancyType string

const (
	Motorcycle OccupancyType = "motorcycle"
	LMV        OccupancyType = "lmv"
	HMV        OccupancyType = "hmv"
)

type Slot struct {
	ID   string
	Type OccupancyType
}

func initLogger() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)
}

func NewParkingLot(ctx context.Context, slots []*Slot) (*ParkingLot, error) {
	initLogger()
	p := &ParkingLot{
		request:         make(chan request),
		ticketRepo:      store.NewTicketRepoMap(),
		slotRepo:        store.NewSlotRepoMap(),
		paymentManager:  &mockPaymentManager{},
		pricingStrategy: &priceStrategist{},
	}
	storeSlots := make([]*store.Slot, 0, len(slots))
	for _, slot := range slots {
		storeSlots = append(storeSlots, &store.Slot{
			ID:       slot.ID,
			State:    store.Vacant,
			Capacity: store.GetSlotSize(string(slot.Type)),
		})
	}
	err := p.slotRepo.InitSlots(storeSlots)
	if err != nil {
		return nil, err
	}

	go p.loop(ctx)
	return p, nil
}

func (p *ParkingLot) loop(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case req, ok := <-p.request:
			if !ok {
				return
			}
			switch req.typ {
			case enter:
				req.resp <- p.enterParking(req.registrationNumber, req.occupancyType, req.ticketType)
			case exit:
				req.resp <- p.exitParking(req.ticketNumber)
			}
		}
	}
}

func (p *ParkingLot) Enter(regNum string, ot OccupancyType) (string, error) {
	req := request{
		typ: enter,
		enterRequest: enterRequest{
			occupancyType:      ot,
			registrationNumber: regNum,
		},
		resp: make(chan *response, 1),
	}
	p.request <- req

	res := <-req.resp
	return res.ticketID, res.err
}

func (p *ParkingLot) Exit(ticketNumber string) error {
	req := request{
		typ: exit,
		exitRequest: exitRequest{
			ticketNumber: ticketNumber,
		},
		resp: make(chan *response, 1),
	}
	p.request <- req

	res := <-req.resp
	return res.err
}

func (p *ParkingLot) enterParking(vehicleNumber string, vehicleType OccupancyType, ticketType store.TicketType) *response {
	//find a relevant slot..
	size := store.GetSlotSize(string(vehicleType))
	slot, err := p.slotRepo.GetEmptySlot(size)
	if err != nil {
		return &response{
			err: err,
		}
	}
	// block the slot
	slot.State = store.Occupied
	err = p.slotRepo.UpdateSlot(slot)
	if err != nil {
		return &response{
			err: err,
		}
	}
	// create a ticket
	ticketID, err := p.ticketRepo.CreateTicket(store.Vehicle{
		Number: vehicleNumber,
		Size:   size,
		Typ:    ticketType,
	}, slot.ID)
	if err != nil {
		slot.State = store.Vacant
		_ = p.slotRepo.UpdateSlot(slot)
		return &response{
			err: err,
		}
	}

	return &response{
		ticketID: ticketID,
	}
}

func (p *ParkingLot) onSuccess(ticketID string) *response {
	ticket, err := p.ticketRepo.GetTicket(ticketID)
	if err != nil {
		return &response{
			err: err,
		}
	}
	slot, err := p.slotRepo.GetSlot(ticket.SlotID)
	if err != nil {
		return &response{
			err: err,
		}
	}

	slot.State = store.Vacant
	err = p.slotRepo.UpdateSlot(slot)
	if err != nil {
		return &response{
			err: err,
		}
	}
	return nil
}

func (p *ParkingLot) exitParking(ticketNumber string) *response {
	//find the ticket
	ticket, err := p.ticketRepo.GetTicket(ticketNumber)
	if err != nil {
		return &response{
			err: err,
		}
	}
	//check the diff
	exitTime := time.Now()
	price, err := p.pricingStrategy.calculate(ticket.CreatedAt, exitTime, ticket.Vehicle.Size, ticket.TicketType)
	if err != nil {
		return &response{
			err: err,
		}
	}

	done := make(chan *response, 1)
	go p.paymentManager.processPayment(context.Background(), &payRequest{
		price:    price,
		ticketID: ticket.ID,
		callback: func(s string) *response {
			res := p.onSuccess(s)
			done <- res
			return res
		},
	})
	return <-done
}
