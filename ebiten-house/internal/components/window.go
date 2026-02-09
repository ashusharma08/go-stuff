package components

import (
	"image/color"
)

// WindowState represents the current state of a window
type WindowState int

const (
	WindowStateClosed WindowState = iota
	WindowStatePartiallyOpen
	WindowStateFullyOpen
	WindowStateBroken
)

// WindowType represents different types of windows
type WindowType string

const (
	WindowTypeRegular  WindowType = "regular"
	WindowTypeSliding  WindowType = "sliding"
	WindowTypeFixed    WindowType = "fixed"
	WindowTypeShutters WindowType = "shutters"
)

// Window is a component that represents a window
type Window struct {
	*BaseComponent
	windowType    WindowType
	state         WindowState
	openPercent   float64
	isBroken      bool
	isReinforced  bool
	transparency  float64
	glassStrength float64
}

// NewWindow creates a new window component
func NewWindow(windowType WindowType, width, height float64, isReinforced bool) *Window {
	var durability float64
	var windowColor color.RGBA
	var transparency float64
	var glassStrength float64

	// Set properties based on window type and reinforcement
	if isReinforced {
		durability = 70
		glassStrength = 2.0
		transparency = 0.7 // Reinforced glass is less transparent
		windowColor = color.RGBA{200, 220, 230, 230}
	} else {
		durability = 30
		glassStrength = 1.0
		transparency = 0.9 // Regular glass is more transparent
		windowColor = color.RGBA{173, 216, 230, 200}
	}

	// Adjust durability based on window type
	switch windowType {
	case WindowTypeFixed:
		durability *= 1.5 // Fixed windows are stronger
	case WindowTypeShutters:
		durability *= 1.2 // Windows with shutters have additional protection
	}

	return &Window{
		BaseComponent: NewBaseComponent("window", width, height, durability, windowColor),
		windowType:    windowType,
		state:         WindowStateClosed,
		openPercent:   0,
		isBroken:      false,
		isReinforced:  isReinforced,
		transparency:  transparency,
		glassStrength: glassStrength,
	}
}

// Update handles the window's logic
func (w *Window) Update() error {
	// Call the base component's Update method
	err := w.BaseComponent.Update()
	if err != nil {
		return err
	}

	// If the window is broken, reduce durability further over time
	if w.isBroken && w.durability > 0 {
		w.TakeDamage(0.1) // Gradual deterioration of broken window
	}

	return nil
}

// Open opens the window to a certain percentage
func (w *Window) Open(percent float64) {
	// Fixed windows can't be opened
	if w.windowType == WindowTypeFixed {
		return
	}

	// Can't open broken windows
	if w.isBroken {
		return
	}

	if percent <= 0 {
		w.openPercent = 0
		w.state = WindowStateClosed
	} else if percent >= 1.0 {
		w.openPercent = 1.0
		w.state = WindowStateFullyOpen
	} else {
		w.openPercent = percent
		w.state = WindowStatePartiallyOpen
	}
}

// Close closes the window
func (w *Window) Close() {
	w.Open(0)
}

// Break causes the window to break
func (w *Window) Break() {
	w.isBroken = true
	w.state = WindowStateBroken
	
	// Reduce durability significantly when broken
	w.TakeDamage(w.durability * 0.75)
	
	// Change appearance to broken glass
	w.color = color.RGBA{200, 200, 200, 150}
}

// Repair repairs the window
func (w *Window) Repair() {
	if w.isBroken && !w.IsDestroyed() {
		w.isBroken = false
		w.state = WindowStateClosed
		w.openPercent = 0
		
		// Restore durability partially
		w.durability = w.maxDurability * 0.8
		
		// Restore original color
		if w.isReinforced {
			w.color = color.RGBA{200, 220, 230, 230}
		} else {
			w.color = color.RGBA{173, 216, 230, 200}
		}
	}
}

// TakeDamage overrides the base method to handle window-specific damage behavior
func (w *Window) TakeDamage(amount float64) {
	// Apply base damage
	w.BaseComponent.TakeDamage(amount)
	
	// Check if the damage should break the window
	if !w.isBroken && w.durability < w.maxDurability*0.3 {
		w.Break()
	}
}

// GetTransparency returns the window's transparency (0-1)
func (w *Window) GetTransparency() float64 {
	if w.isBroken {
		return w.transparency * 0.5 // Broken windows have reduced transparency
	}
	return w.transparency
}

// IsOpen returns whether the window is open
func (w *Window) IsOpen() bool {
	return w.state == WindowStatePartiallyOpen || w.state == WindowStateFullyOpen
}

// IsBroken returns whether the window is broken
func (w *Window) IsBroken() bool {
	return w.isBroken
}

// GetType returns the window type
func (w *Window) GetType() WindowType {
	return w.windowType
}

// GetState returns the window state
func (w *Window) GetState() WindowState {
	return w.state
}