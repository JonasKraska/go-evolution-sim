package engine

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
	"math"
	"time"
)

type Engine struct {
	game  Gamer
	zoom  int
	speed int64

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

func (e *Engine) SetTicksPerSecond(tps uint16) *Engine {
	e.speed = int64(tps) / 60
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
	var delta time.Duration

	for {
		delta = time.Since(e.lastUpdate) * time.Duration(e.speed)
		if delta > time.Microsecond {
			break
		}
	}

	DebugReset()
	DebugPrintln(fmt.Sprintf("Game TPS: %.2f", ebiten.ActualTPS()))
	DebugPrintln(fmt.Sprintf("Game FPS: %.2f", ebiten.ActualFPS()))

	e.updateNode(e.game, delta)
	e.moveNode(e.game, delta)

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

	// show debug output
	DebugPrint(screen)
}

func (e *Engine) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1280, 720
}

func (e *Engine) updateNode(node Noder, delta time.Duration) {
	if updater, ok := node.(Updater); ok {
		updater.Update(delta)
	}

	for _, n := range node.GetChildren() {
		e.updateNode(n, delta)
	}

	node.removeChildren()
}

func (e *Engine) moveNode(node Noder, delta time.Duration) {
	if mover, ok := node.(Mover); ok {
		mover.doMove()

		if e.game.Contains(mover.GetPosition()) == false {
			mover.cancelMove()
		}

		e.game.GetGrid().Update(mover)
	}

	for _, n := range node.GetChildren() {
		e.moveNode(n, delta)
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
		options.GeoM.Translate(math.Floor(position.X), math.Floor(position.Y))

		frame.DrawImage(sprite, options)
	}

	return frame
}
