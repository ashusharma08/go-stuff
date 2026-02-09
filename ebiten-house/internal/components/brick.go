package components

import (
	"image/color"
)

// BrickType represents different types of bricks
type BrickType string

const (
	BrickTypeNormal BrickType = "normal"
	BrickTypeStone  BrickType = "stone"
	BrickTypeMetal  BrickType = "metal"
)

// Brick is a basic building component
type Brick struct {
	*BaseComponent
	brickType BrickType
}

// NewBrick creates a new brick component
func NewBrick(brickType BrickType, width, height float64) *Brick {
	var durability float64
	var brickColor color.RGBA

	switch brickType {
	case BrickTypeStone:
		durability = 100
		brickColor = color.RGBA{128, 128, 128, 255} // Gray
	case BrickTypeMetal:
		durability = 200
		brickColor = color.RGBA{169, 169, 169, 255} // Dark gray
	default: // Normal brick
		durability = 50
		brickColor = color.RGBA{205, 92, 92, 255} // Brown
	}

	return &Brick{
		BaseComponent: NewBaseComponent("brick", width, height, durability, brickColor),
		brickType:     brickType,
	}
}

// GetBrickType returns the type of brick
func (b *Brick) GetBrickType() BrickType {
	return b.brickType
}

// GetStrength returns the strength of the brick based on its type
func (b *Brick) GetStrength() float64 {
	switch b.brickType {
	case BrickTypeStone:
		return 2.0
	case BrickTypeMetal:
		return 4.0
	default:
		return 1.0
	}
}

// IsStructural returns whether the brick is a structural component
func (b *Brick) IsStructural() bool {
	return true
}