package components

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

// Component represents a game object in the Fortress Architect game
type Component interface {
	// Update handles the component's logic
	Update() error
	
	// Draw renders the component to the screen
	Draw(screen *ebiten.Image)
	
	// Position returns the component's position
	Position() (float64, float64)
	
	// SetPosition sets the component's position
	SetPosition(x, y float64)
	
	// Size returns the component's width and height
	Size() (float64, float64)
	
	// Rotate rotates the component by the given angle (in radians)
	Rotate(angle float64)
	
	// Rotation returns the current rotation angle (in radians)
	Rotation() float64
	
	// Type returns the component type
	Type() string
	
	// Durability returns the component's current durability
	Durability() float64
	
	// MaxDurability returns the component's maximum durability
	MaxDurability() float64
	
	// TakeDamage reduces the component's durability
	TakeDamage(amount float64)
	
	// IsDestroyed returns true if the component is destroyed
	IsDestroyed() bool
}

// BaseComponent provides a base implementation of the Component interface
type BaseComponent struct {
	componentType string
	x, y          float64
	width, height float64
	rotation      float64
	durability    float64
	maxDurability float64
	color         color.RGBA
	image         *ebiten.Image
}

// NewBaseComponent creates a new base component
func NewBaseComponent(componentType string, width, height, durability float64, color color.RGBA) *BaseComponent {
	return &BaseComponent{
		componentType: componentType,
		width:         width,
		height:        height,
		durability:    durability,
		maxDurability: durability,
		color:         color,
	}
}

// Update handles the component's logic
func (b *BaseComponent) Update() error {
	return nil
}

// Draw renders the component to the screen
func (b *BaseComponent) Draw(screen *ebiten.Image) {
	if b.image == nil {
		// Create a new image if it doesn't exist
		b.image = ebiten.NewImage(int(b.width), int(b.height))
		b.image.Fill(b.color)
	}

	// Create a new options struct for drawing the image
	op := &ebiten.DrawImageOptions{}
	
	// Set the position
	op.GeoM.Translate(-b.width/2, -b.height/2)
	op.GeoM.Rotate(b.rotation)
	op.GeoM.Translate(b.x, b.y)
	
	// Draw the image
	screen.DrawImage(b.image, op)
}

// Position returns the component's position
func (b *BaseComponent) Position() (float64, float64) {
	return b.x, b.y
}

// SetPosition sets the component's position
func (b *BaseComponent) SetPosition(x, y float64) {
	b.x, b.y = x, y
}

// Size returns the component's width and height
func (b *BaseComponent) Size() (float64, float64) {
	return b.width, b.height
}

// Rotate rotates the component by the given angle (in radians)
func (b *BaseComponent) Rotate(angle float64) {
	b.rotation = angle
}

// Rotation returns the current rotation angle (in radians)
func (b *BaseComponent) Rotation() float64 {
	return b.rotation
}

// Type returns the component type
func (b *BaseComponent) Type() string {
	return b.componentType
}

// Durability returns the component's current durability
func (b *BaseComponent) Durability() float64 {
	return b.durability
}

// MaxDurability returns the component's maximum durability
func (b *BaseComponent) MaxDurability() float64 {
	return b.maxDurability
}

// TakeDamage reduces the component's durability
func (b *BaseComponent) TakeDamage(amount float64) {
	b.durability -= amount
	if b.durability < 0 {
		b.durability = 0
	}
}

// IsDestroyed returns true if the component is destroyed
func (b *BaseComponent) IsDestroyed() bool {
	return b.durability <= 0
}