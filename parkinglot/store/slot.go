package store

type SlotState int

const (
	Vacant SlotState = iota
	Occupied
)

type Size int

const (
	INVALID Size = 0
	MC           = iota
	LMV
	HMV
)

func GetSlotSize(t string) Size {
	switch t {
	case "motorcycle":
		return MC
	case "lmv":
		return LMV
	case "hmv":
		return HMV
	default:
		return 0
	}
}

type Slot struct {
	ID       string
	State    SlotState
	Capacity Size
	//should we add vehicle in slot to get a quick check on who is sitting there?
}

type SlotRepo interface {
	GetEmptySlot(Size) (*Slot, error)
	InitSlots(slots []*Slot) error
	UpdateSlot(slot *Slot) error
	GetSlot(slotID string) (*Slot, error)
}

type slotRepoMap struct {
	slots map[string]*Slot
}

func NewSlotRepoMap() *slotRepoMap {
	return &slotRepoMap{}
}

func (s *slotRepoMap) GetEmptySlot(sc Size) (*Slot, error) {
	for _, slot := range s.slots {
		if slot.Capacity >= sc && slot.State == Vacant {
			return slot, nil
		}
	}
	return nil, ErrNotFound
}

func (s *slotRepoMap) InitSlots(slots []*Slot) error {
	s.slots = make(map[string]*Slot, len(slots))
	for _, slot := range slots {
		s.slots[slot.ID] = slot
	}
	return nil
}

func (s *slotRepoMap) UpdateSlot(slot *Slot) error {
	_, ok := s.slots[slot.ID]
	if !ok {
		return ErrNotFound
	}
	s.slots[slot.ID].State = slot.State
	return nil
}

func (s *slotRepoMap) GetSlot(slotID string) (*Slot, error) {
	slot, ok := s.slots[slotID]
	if !ok {
		return nil, ErrNotFound
	}
	return slot, nil
}
