package main

import (
	"fmt"
	"image/color"
	"log"
	"math"

	"github.com/esoptra/go-prac/ebiten-house/internal/components"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
)

const (
	screenWidth  = 1280
	screenHeight = 720
	gameTitle    = "Fortress Architect"

	// Grid settings
	gridSize   = 32
	gridWidth  = 100
	gridHeight = 100

	// Camera settings
	cameraSpeed = 5
)

// Game implements the ebiten.Game interface.
type Game struct {
	// Game state
	mode              GameMode
	buildMode         BuildMode
	selectedComponent ComponentType

	// Components
	components []components.Component

	// Camera
	cameraX, cameraY float64

	// Grid
	grid [][]Cell

	// Mouse state
	mouseX, mouseY int
	mousePressed   bool
}

// GameMode represents the current mode of the game
type GameMode int

const (
	ModeMainMenu GameMode = iota
	ModeBuild
	ModeCombat
)

// BuildMode represents the current building mode
type BuildMode int

const (
	BuildModeNone BuildMode = iota
	BuildModeTerrain
	BuildModeStructure
	BuildModeInterior
	BuildModeDefense
)

// ComponentType represents the type of component to place
type ComponentType int

const (
	ComponentNone ComponentType = iota
	ComponentBrick
	ComponentDoor
	ComponentWindow
)

// Cell represents a single cell in the grid
type Cell struct {
	Occupied    bool
	Component   components.Component
	TerrainType TerrainType
}

// TerrainType represents different types of terrain
type TerrainType int

const (
	TerrainNone TerrainType = iota
	TerrainGrass
	TerrainDirt
	TerrainStone
	TerrainSand
	TerrainWater
)

// NewGame creates a new game instance
func NewGame() *Game {
	g := &Game{
		mode:              ModeBuild,
		buildMode:         BuildModeStructure,
		selectedComponent: ComponentBrick,
		components:        make([]components.Component, 0),
		cameraX:           0,
		cameraY:           0,
	}

	// Initialize grid
	g.initGrid()

	return g
}

// initGrid initializes the grid
func (g *Game) initGrid() {
	g.grid = make([][]Cell, gridHeight)
	for y := 0; y < gridHeight; y++ {
		g.grid[y] = make([]Cell, gridWidth)
		for x := 0; x < gridWidth; x++ {
			// Initialize with grass terrain
			g.grid[y][x] = Cell{
				Occupied:    false,
				TerrainType: TerrainGrass,
			}
		}
	}
}

// Update updates the game state
func (g *Game) Update() error {
	// Update mouse position
	g.mouseX, g.mouseY = ebiten.CursorPosition()
	g.mousePressed = ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)

	// Handle input based on current mode
	switch g.mode {
	case ModeMainMenu:
		g.updateMainMenu()
	case ModeBuild:
		g.updateBuildMode()
	case ModeCombat:
		g.updateCombatMode()
	}

	// Update all components
	for _, c := range g.components {
		err := c.Update()
		if err != nil {
			return err
		}
	}

	return nil
}

// updateMainMenu handles input in the main menu
func (g *Game) updateMainMenu() {
	// Check for menu selections
	if inpututil.IsKeyJustPressed(ebiten.KeyB) {
		g.mode = ModeBuild
	} else if inpututil.IsKeyJustPressed(ebiten.KeyC) {
		g.mode = ModeCombat
	}
}

// updateBuildMode handles input in build mode
func (g *Game) updateBuildMode() {
	// Handle camera movement
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.cameraY -= cameraSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.cameraY += cameraSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.cameraX -= cameraSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.cameraX += cameraSpeed
	}

	// Handle component selection
	if inpututil.IsKeyJustPressed(ebiten.Key1) {
		g.selectedComponent = ComponentBrick
	} else if inpututil.IsKeyJustPressed(ebiten.Key2) {
		g.selectedComponent = ComponentDoor
	} else if inpututil.IsKeyJustPressed(ebiten.Key3) {
		g.selectedComponent = ComponentWindow
	}

	// Handle building mode selection
	if inpututil.IsKeyJustPressed(ebiten.KeyT) {
		g.buildMode = BuildModeTerrain
	} else if inpututil.IsKeyJustPressed(ebiten.KeyB) {
		g.buildMode = BuildModeStructure
	} else if inpututil.IsKeyJustPressed(ebiten.KeyI) {
		g.buildMode = BuildModeInterior
	} else if inpututil.IsKeyJustPressed(ebiten.KeyF) {
		g.buildMode = BuildModeDefense
	}

	// Handle mouse input for placing components
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		// Convert screen coordinates to grid coordinates
		gridX, gridY := g.screenToGrid(g.mouseX, g.mouseY)

		// Check if coordinates are within grid bounds
		if gridX >= 0 && gridX < gridWidth && gridY >= 0 && gridY < gridHeight {
			// Place component if the cell is not occupied
			if !g.grid[gridY][gridX].Occupied {
				g.placeComponent(gridX, gridY)
				fmt.Printf("Placed component at grid position: %d, %d\n", gridX, gridY)
			}
		}
	}

	// Switch to combat mode with Enter key
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		g.mode = ModeCombat
	}
}

// updateCombatMode handles input in combat mode
func (g *Game) updateCombatMode() {
	// Handle camera movement (same as build mode)
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.cameraY -= cameraSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.cameraY += cameraSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.cameraX -= cameraSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.cameraX += cameraSpeed
	}

	// Switch back to build mode with Escape key
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		g.mode = ModeBuild
	}
}

// placeComponent places a component at the specified grid coordinates
func (g *Game) placeComponent(gridX, gridY int) {
	// Calculate world position from grid coordinates
	worldX := float64(gridX*gridSize) + gridSize/2
	worldY := float64(gridY*gridSize) + gridSize/2

	var comp components.Component

	// Create component based on selected type
	switch g.selectedComponent {
	case ComponentBrick:
		brick := components.NewBrick(components.BrickTypeNormal, float64(gridSize), float64(gridSize))
		brick.SetPosition(worldX, worldY)
		comp = brick
	case ComponentDoor:
		door := components.NewDoor(components.DoorMaterialWood, float64(gridSize), float64(gridSize))
		door.SetPosition(worldX, worldY)
		comp = door
	case ComponentWindow:
		window := components.NewWindow(components.WindowTypeRegular, float64(gridSize), float64(gridSize), false)
		window.SetPosition(worldX, worldY)
		comp = window
	}

	// Add component to the game
	if comp != nil {
		g.components = append(g.components, comp)
		g.grid[gridY][gridX].Occupied = true
		g.grid[gridY][gridX].Component = comp
	}
}

// Draw draws the game
func (g *Game) Draw(screen *ebiten.Image) {
	// Clear the screen
	screen.Fill(color.RGBA{30, 30, 30, 255})

	// Draw based on current mode
	switch g.mode {
	case ModeMainMenu:
		g.drawMainMenu(screen)
	case ModeBuild:
		g.drawBuildMode(screen)
	case ModeCombat:
		g.drawCombatMode(screen)
	}
}

// drawMainMenu draws the main menu
func (g *Game) drawMainMenu(screen *ebiten.Image) {
	// Draw title
	ebitenutil.DebugPrintAt(screen, gameTitle, screenWidth/2-50, 100)

	// Draw menu options
	ebitenutil.DebugPrintAt(screen, "Press B to enter Build Mode", screenWidth/2-100, 200)
	ebitenutil.DebugPrintAt(screen, "Press C to enter Combat Mode", screenWidth/2-100, 230)
}

// drawBuildMode draws the build mode
func (g *Game) drawBuildMode(screen *ebiten.Image) {
	// Draw grid
	g.drawGrid(screen)

	// Draw all components
	for _, c := range g.components {
		x, y := c.Position()

		// Check if component is within view
		screenX, screenY := g.worldToScreen(x, y)
		if screenX >= -gridSize && screenX <= screenWidth+gridSize &&
			screenY >= -gridSize && screenY <= screenHeight+gridSize {
			c.Draw(screen)
		}
	}

	// Debug: Display component count and camera position
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Components: %d", len(g.components)), 10, 180)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Camera: (%.1f,%.1f)", g.cameraX, g.cameraY), 10, 200)

	// Draw UI
	g.drawBuildUI(screen)

	// Draw grid cursor
	g.drawGridCursor(screen)
}

// drawCombatMode draws the combat mode
func (g *Game) drawCombatMode(screen *ebiten.Image) {
	// Draw grid
	g.drawGrid(screen)

	// Draw all components
	for _, c := range g.components {
		c.Draw(screen)
	}

	// Draw UI
	ebitenutil.DebugPrintAt(screen, "COMBAT MODE", 10, 10)
	ebitenutil.DebugPrintAt(screen, "Press ESC to return to Build Mode", 10, 30)
}

// drawGrid draws the grid
func (g *Game) drawGrid(screen *ebiten.Image) {
	// Calculate visible grid range based on camera position
	startX := int(math.Max(0, math.Floor(g.cameraX/gridSize)))
	endX := int(math.Min(float64(gridWidth), math.Ceil((g.cameraX+screenWidth)/gridSize)))
	startY := int(math.Max(0, math.Floor(g.cameraY/gridSize)))
	endY := int(math.Min(float64(gridHeight), math.Ceil((g.cameraY+screenHeight)/gridSize)))

	// Draw grid cells
	for y := startY; y < endY; y++ {
		for x := startX; x < endX; x++ {
			// Calculate screen position
			screenX, screenY := g.worldToScreen(float64(x*gridSize), float64(y*gridSize))

			// Draw cell based on terrain type
			var cellColor color.RGBA
			switch g.grid[y][x].TerrainType {
			case TerrainGrass:
				cellColor = color.RGBA{50, 150, 50, 255}
			case TerrainDirt:
				cellColor = color.RGBA{150, 100, 50, 255}
			case TerrainStone:
				cellColor = color.RGBA{120, 120, 120, 255}
			case TerrainSand:
				cellColor = color.RGBA{200, 180, 100, 255}
			case TerrainWater:
				cellColor = color.RGBA{50, 100, 200, 255}
			default:
				cellColor = color.RGBA{50, 50, 50, 255}
			}

			// Draw cell background
			ebitenutil.DrawRect(screen, float64(screenX), float64(screenY), float64(gridSize), float64(gridSize), cellColor)

			// Draw grid lines
			ebitenutil.DrawLine(screen, float64(screenX), float64(screenY), float64(screenX+gridSize), float64(screenY), color.RGBA{0, 0, 0, 100})
			ebitenutil.DrawLine(screen, float64(screenX), float64(screenY), float64(screenX), float64(screenY+gridSize), color.RGBA{0, 0, 0, 100})
		}
	}
}

// drawBuildUI draws the build mode UI
func (g *Game) drawBuildUI(screen *ebiten.Image) {
	// Draw mode and selected component
	var modeText string
	switch g.buildMode {
	case BuildModeTerrain:
		modeText = "Terrain Mode"
	case BuildModeStructure:
		modeText = "Structure Mode"
	case BuildModeInterior:
		modeText = "Interior Mode"
	case BuildModeDefense:
		modeText = "Defense Mode"
	}

	var componentText string
	switch g.selectedComponent {
	case ComponentBrick:
		componentText = "Brick"
	case ComponentDoor:
		componentText = "Door"
	case ComponentWindow:
		componentText = "Window"
	}

	// Draw UI text
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Mode: %s", modeText), 10, 10)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Selected: %s", componentText), 10, 30)
	ebitenutil.DebugPrintAt(screen, "Controls:", 10, 60)
	ebitenutil.DebugPrintAt(screen, "WASD: Move Camera", 10, 80)
	ebitenutil.DebugPrintAt(screen, "1-3: Select Component", 10, 100)
	ebitenutil.DebugPrintAt(screen, "T,B,I,F: Change Build Mode", 10, 120)
	ebitenutil.DebugPrintAt(screen, "Click: Place Component", 10, 140)
	ebitenutil.DebugPrintAt(screen, "Enter: Switch to Combat Mode", 10, 160)

	// Draw component palette
	g.drawComponentPalette(screen)
}

// drawComponentPalette draws the component selection palette
func (g *Game) drawComponentPalette(screen *ebiten.Image) {
	// Draw palette background
	ebitenutil.DrawRect(screen, float64(screenWidth-200), 10, 190, 100, color.RGBA{50, 50, 50, 200})

	// Draw component buttons
	buttonWidth := 50
	buttonHeight := 50
	buttonSpacing := 10
	startX := screenWidth - 190
	startY := 20

	// Brick button
	buttonColor := color.RGBA{205, 92, 92, 255}
	if g.selectedComponent == ComponentBrick {
		buttonColor = color.RGBA{255, 120, 120, 255}
	}
	ebitenutil.DrawRect(screen, float64(startX), float64(startY), float64(buttonWidth), float64(buttonHeight), buttonColor)
	text.Draw(screen, "1", basicfont.Face7x13, startX+buttonWidth/2-5, startY+buttonHeight/2+5, color.White)

	// Door button
	buttonColor = color.RGBA{139, 69, 19, 255}
	if g.selectedComponent == ComponentDoor {
		buttonColor = color.RGBA{160, 90, 40, 255}
	}
	ebitenutil.DrawRect(screen, float64(startX+buttonWidth+buttonSpacing), float64(startY), float64(buttonWidth), float64(buttonHeight), buttonColor)
	text.Draw(screen, "2", basicfont.Face7x13, startX+buttonWidth+buttonSpacing+buttonWidth/2-5, startY+buttonHeight/2+5, color.White)

	// Window button
	buttonColor = color.RGBA{173, 216, 230, 200}
	if g.selectedComponent == ComponentWindow {
		buttonColor = color.RGBA{200, 240, 255, 220}
	}
	ebitenutil.DrawRect(screen, float64(startX+2*(buttonWidth+buttonSpacing)), float64(startY), float64(buttonWidth), float64(buttonHeight), buttonColor)
	text.Draw(screen, "3", basicfont.Face7x13, startX+2*(buttonWidth+buttonSpacing)+buttonWidth/2-5, startY+buttonHeight/2+5, color.White)
}

// drawGridCursor draws the cursor on the grid
func (g *Game) drawGridCursor(screen *ebiten.Image) {
	// Convert mouse position to grid coordinates
	gridX, gridY := g.screenToGrid(g.mouseX, g.mouseY)

	// Check if coordinates are within grid bounds
	if gridX >= 0 && gridX < gridWidth && gridY >= 0 && gridY < gridHeight {
		// Convert grid coordinates back to screen coordinates
		screenX, screenY := g.worldToScreen(float64(gridX*gridSize), float64(gridY*gridSize))

		// Draw cursor rectangle
		var cursorColor color.RGBA
		if g.grid[gridY][gridX].Occupied {
			cursorColor = color.RGBA{255, 0, 0, 100} // Red for occupied
		} else {
			cursorColor = color.RGBA{0, 255, 0, 100} // Green for available
		}

		ebitenutil.DrawRect(screen, float64(screenX), float64(screenY), float64(gridSize), float64(gridSize), cursorColor)
	}
}

// screenToGrid converts screen coordinates to grid coordinates
func (g *Game) screenToGrid(screenX, screenY int) (int, int) {
	worldX := float64(screenX) + g.cameraX
	worldY := float64(screenY) + g.cameraY
	return int(worldX) / gridSize, int(worldY) / gridSize
}

// worldToScreen converts world coordinates to screen coordinates
func (g *Game) worldToScreen(worldX, worldY float64) (int, int) {
	screenX := int(worldX - g.cameraX)
	screenY := int(worldY - g.cameraY)
	return screenX, screenY
}

// Layout returns the game's logical screen size
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle(gameTitle)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	game := NewGame()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
