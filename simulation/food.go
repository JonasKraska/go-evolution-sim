package simulation

import (
	"github.com/JonasKraska/go-evolution-sim/engine"
	"github.com/JonasKraska/go-evolution-sim/engine/random"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"time"
)

type Food struct {
	engine.Node
	engine.Placeable

	energy Energy
}

func NewFood(position engine.Position, energy Energy) *Food {
	f := &Food{
		energy: energy,
	}

	f.SetPosition(position)

	return f
}

func (f *Food) Update(delta time.Duration) {
	f.grow(delta)
	f.proliferate()
}

var foodSprite *ebiten.Image

func (f *Food) Draw() *ebiten.Image {
	if foodSprite == nil {
		foodSprite = ebiten.NewImage(1, 1)
		foodSprite.Fill(color.RGBA{R: 219, G: 93, B: 81, A: 255})
	}

	return foodSprite
}

func (f *Food) grow(delta time.Duration) {
	f.energy += simulation.config.FoodGrowthRate * delta.Seconds()
}

func (f *Food) proliferate() {
	if f.energy >= float64(simulation.config.FoodProliferationThreshold) {
		f.energy = f.energy / 2
		var position engine.Position

		for {
			diffX := random.IntBetween(-3, 3)
			diffY := random.IntBetween(-3, 3)

			if diffX == 0 && diffY == 0 {
				continue
			}

			position = engine.Position{
				X: f.GetPosition().X + float64(diffX),
				Y: f.GetPosition().Y + float64(diffY),
			}

			if world.Contains(position) {
				break
			}
		}

		world.spawnFood(position, f.energy)
	}
}
