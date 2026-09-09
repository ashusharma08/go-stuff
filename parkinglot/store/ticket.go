package store

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type TicketType int

const (
	Hour TicketType = iota
	Day
)

type ticketStatus int

const (
	PendingPayment ticketStatus = iota
	Paid
	Failed
)

var (
	ErrNotFound = errors.New("Not found")
)

type Vehicle struct {
	Number string
	Size   Size
	Typ    TicketType
}

type Ticket struct {
	TicketType TicketType
	Vehicle    Vehicle
	ID         string
	SlotID     string
	CreatedAt  time.Time
}

type TicketRepo interface {
	GetTicket(ticketID string) (*Ticket, error)
	CreateTicket(veh Vehicle, slotID string) (string, error)
}

type ticketRepoMap struct {
	tickets map[string]*Ticket
}

func NewTicketRepoMap() *ticketRepoMap {
	return &ticketRepoMap{
		tickets: make(map[string]*Ticket),
	}
}

func (t *ticketRepoMap) GetTicket(ticketID string) (*Ticket, error) {
	v, ok := t.tickets[ticketID]
	if !ok {
		return nil, ErrNotFound
	}
	return v, nil
}
func (t *ticketRepoMap) CreateTicket(veh Vehicle, slotID string) (string, error) {
	tid := uuid.New()
	tc := &Ticket{
		ID:        tid.String(),
		SlotID:    slotID,
		CreatedAt: time.Now(),
		Vehicle:   veh,
	}
	t.tickets[tid.String()] = tc
	return tc.ID, nil
}
