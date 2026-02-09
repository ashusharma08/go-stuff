package components

import (
	"image/color"
	"math"
)

// DoorState represents the current state of a door
type DoorState int

const (
	DoorStateClosed DoorState = iota
	DoorStateOpening
	DoorStateOpen
	DoorStateClosing
)

// DoorMaterial represents the material of a door
type DoorMaterial string

const (
	DoorMaterialWood  DoorMaterial = "wood"
	DoorMaterialMetal DoorMaterial = "metal"
	DoorMaterialGlass DoorMaterial = "glass"
)

// Door is a component that can be opened and closed
type Door struct {
	*BaseComponent
	isOpen         bool
	openPercent    float64
	state          DoorState
	material       DoorMaterial
	openSpeed      float64
	isLocked       bool
	openDirection  float64 // Angle in radians for the direction the door opens
	maxOpenAngle   float64 // Maximum angle the door can open
	originalRotation float64 // Keep track of the original rotation
}

// NewDoor creates a new door component
func NewDoor(material DoorMaterial, width, height float64) *Door {
	var durability float64
	var doorColor color.RGBA

	switch material {
	case DoorMaterialMetal:
		durability = 150
		doorColor = color.RGBA{169, 169, 169, 255} // Dark gray
	case DoorMaterialGlass:
		durability = 30
		doorColor = color.RGBA{173, 216, 230, 200} // Light blue with transparency
	default: // Wood
		durability = 75
		doorColor = color.RGBA{139, 69, 19, 255} // Brown
	}

	return &Door{
		BaseComponent:   NewBaseComponent("door", width, height, durability, doorColor),
		isOpen:          false,
		openPercent:     0,
		state:           DoorStateClosed,
		material:        material,
		openSpeed:       0.05, // Open 5% per update
		isLocked:        false,
		openDirection:   0,   // Default opens to the right
		maxOpenAngle:    math.Pi / 2, // 90 degrees
		originalRotation: 0,
	}
}

// Update handles the door's logic for opening/closing animation
func (d *Door) Update() error {
	// Call the base component's Update method
	err := d.BaseComponent.Update()
	if err != nil {
		return err
	}

	// Update door state based on current state
	switch d.state {
	case DoorStateOpening:
		d.openPercent += d.openSpeed
		if d.openPercent >= 1.0 {
			d.openPercent = 1.0
			d.state = DoorStateOpen
			d.isOpen = true
		}
		// Update door rotation based on open percentage
		d.Rotate(d.originalRotation + d.openDirection*d.maxOpenAngle*d.openPercent)
	case DoorStateClosing:
		d.openPercent -= d.openSpeed
		if d.openPercent <= 0.0 {
			d.openPercent = 0.0
			d.state = DoorStateClosed
			d.isOpen = false
		}
		// Update door rotation based on open percentage
		d.Rotate(d.originalRotation + d.openDirection*d.maxOpenAngle*d.openPercent)
	}

	return nil
}

// Open starts opening the door
func (d *Door) Open() {
	if d.state != DoorStateOpen && d.state != DoorStateOpening && !d.isLocked {
		d.state = DoorStateOpening
	}
}

// Close starts closing the door
func (d *Door) Close() {
	if d.state != DoorStateClosed && d.state != DoorStateClosing {
		d.state = DoorStateClosing
	}
}

// Toggle toggles the door's open/closed state
func (d *Door) Toggle() {
	if d.isOpen {
		d.Close()
	} else {
		d.Open()
	}
}

// IsOpen returns whether the door is fully open
func (d *Door) IsOpen() bool {
	return d.isOpen
}

// SetLocked sets whether the door is locked
func (d *Door) SetLocked(locked bool) {
	d.isLocked = locked
}

// IsLocked returns whether the door is locked
func (d *Door) IsLocked() bool {
	return d.isLocked
}

// SetOpenDirection sets the direction the door opens
// angle is in radians, where 0 is right, PI/2 is down, PI is left, 3PI/2 is up
func (d *Door) SetOpenDirection(angle float64) {
	d.openDirection = angle
}

// SetOriginalRotation sets the door's original rotation and resets current rotation
func (d *Door) SetOriginalRotation(angle float64) {
	d.originalRotation = angle
	d.Rotate(angle)
}

// GetMaterial returns the door's material
func (d *Door) GetMaterial() DoorMaterial {
	return d.material
}