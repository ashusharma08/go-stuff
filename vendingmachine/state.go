package vendingmachine

type state int

const (
	stateIdle state = iota
	stateSelected
	statePayment
	stateDispensing
)

func (s state) String() string {
	switch s {
	case stateIdle:
		return "idle"
	case stateSelected:
		return "selected"
	case statePayment:
		return "payment"
	case stateDispensing:
		return "dispensing"
	default:
		return "unknown"
	}
}
