package parkinglot

import (
	"time"

	"github.com/esoptra/go-prac/parkinglot/store"
)

type pricingStrategy interface {
	calculate(entry, exit time.Time, vehicleType store.Size, ticketType store.TicketType) (float64, error)
}

var priceMap = map[store.Size]map[store.TicketType]float64{
	store.MC: {
		store.Day:  50,
		store.Hour: 10,
	},
	store.LMV: {
		store.Day:  100,
		store.Hour: 20,
	},
	store.HMV: {
		store.Day:  200,
		store.Hour: 40,
	},
}

type priceStrategist struct {
}

func (p *priceStrategist) calculate(entry, exit time.Time, vehicleType store.Size, ticketType store.TicketType) (float64, error) {
	diff := exit.Sub(entry)                    // exit.Sub(entry) *time for testing. should be seconds..
	price := priceMap[vehicleType][ticketType] //add a check for "exists"
	hours := diff.Hours()

	switch ticketType {
	case store.Day:
		return price * (hours / 24), nil
	case store.Hour:
		return price * hours, nil
	default:
		return 0.0, InvalidTicket
	}
}
