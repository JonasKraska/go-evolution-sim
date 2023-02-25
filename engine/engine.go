package engine

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

type Engine struct {
	game Gamer
	zoom int
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

	if err := ebiten.RunGame(e); err != nil {
		log.Fatal(err)
	}
}

func (e *Engine) Update() error {
	e.updateNode(e.game)

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

func (e *Engine) updateNode(node any) {
	if updater, ok := node.(Updater); ok {
		updater.Update()
	}

	if mover, ok := node.(Mover); ok {
		nextPosition := mover.getNextPosition()
        dimensions := e.game.GetDimensions()
        if nextPosition.X < 0 || nextPosition.Y < 0 || nextPosition.X >= dimensions.W || nextPosition.Y >= dimensions.H {
			mover.cancelMove()
		}

		mover.doMove()
	}

	if nester, ok := node.(Nester); ok {
		for _, n := range nester.GetChildren() {
			e.updateNode(n)
		}
	}
}

func (e *Engine) drawNode(node any) *ebiten.Image {
	drawer, isDrawer := node.(Drawer)

	if isDrawer == false {
		return nil
	}

	frame := drawer.Draw()

	if nester, ok := node.(Nester); ok {
		for _, child := range nester.GetChildren() {
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
	}

	return frame
}
