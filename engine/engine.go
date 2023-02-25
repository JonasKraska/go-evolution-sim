package engine

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
    "time"
)

type Engine struct {
	game Gamer
	zoom int

    lastUpdate time.Time
}

func New() *Engine {
	return &Engine{
		zoom: 1,
	}
}

func (e *Engine) SetZoom(factor uint8) *Engine {
    e.zoom = int(factor)
	return e
}

func (e *Engine) SetTicksPerSecond(tps uint8) *Engine {
    ebiten.SetTPS(int(tps))
	return e
}

func (e *Engine) Run(game Gamer) {
	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("Genetic Simulation")

	e.game = game
    e.lastUpdate = time.Now()

	if err := ebiten.RunGame(e); err != nil {
		log.Fatal(err)
	}
}

func (e *Engine) Update() error {
    delta := time.Since(e.lastUpdate)
    e.updateNode(e.game, delta)
    e.lastUpdate = time.Now()

	return nil
}

func (e *Engine) Draw(screen *ebiten.Image) {
	frame := e.drawNode(e.game)
	options := &ebiten.DrawImageOptions{}

	// scale frame according to current zoom level
	options.GeoM.Scale(float64(e.zoom), float64(e.zoom))

	// center frame on screen
	options.GeoM.Translate(
		float64((screen.Bounds().Dx()-(frame.Bounds().Dx()*e.zoom))/2),
		float64((screen.Bounds().Dy()-(frame.Bounds().Dy()*e.zoom))/2),
	)

	// draw frame on screen
	screen.DrawImage(frame, options)

	// show debug labels
	ebitenutil.DebugPrint(screen, fmt.Sprintf("Game TPS: %.2f", ebiten.ActualTPS()))
	ebitenutil.DebugPrint(screen, fmt.Sprintf("\nGame FPS: %.2f", ebiten.ActualFPS()))
}

func (e *Engine) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1280, 720
}

func (e *Engine) updateNode(node Noder, delta time.Duration) {
	if updater, ok := node.(Updater); ok {
        updater.Update(delta)
	}

	if mover, ok := node.(Mover); ok {
		nextPosition := mover.getNextPosition()
        dimensions := e.game.GetDimensions()
        if nextPosition.X < 0 || nextPosition.Y < 0 || nextPosition.X >= dimensions.W || nextPosition.Y >= dimensions.H {
			mover.cancelMove()
		}

		mover.doMove()
	}

    for _, n := range node.GetChildren() {
            e.updateNode(n, delta)
    }
}

func (e *Engine) drawNode(node Noder) *ebiten.Image {
	drawer, isDrawer := node.(Drawer)

	if isDrawer == false {
		return nil
	}

	frame := drawer.Draw()

    for _, child := range node.GetChildren() {
    placer, isPlacer := child.(Placer)
    sprite := e.drawNode(child)

    if sprite == nil || isPlacer == false {
    continue
    }

    position := placer.GetPosition()

    options := &ebiten.DrawImageOptions{}
    options.GeoM.Translate(float64(position.X), float64(position.Y))

    frame.DrawImage(sprite, options)
    }

	return frame
}
