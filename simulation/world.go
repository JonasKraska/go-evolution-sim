package simulation

import (
	"github.com/JonasKraska/go-evolution-sim/engine"
	"github.com/JonasKraska/go-evolution-sim/engine/random"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

type World struct {
	engine.Game

	size engine.Size
}

type WorldConfig struct {
	Width     int
	Height    int
	Food      []FoodCohort
	Organisms []OrganismCohort
}

type OrganismCohort struct {
	Count  int
	Energy Energy
	Genome Genome
}

type FoodCohort struct {
	Count  int
	Energy Energy
}

func NewWorld(config WorldConfig) *World {
	if config.Width <= 0 {
		config.Width = 64
	}

	if config.Height <= 0 {
		config.Height = 64
	}

	w := &World{
		size: engine.Size{
			W: config.Width,
			H: config.Height,
		},
	}

	for _, cohort := range config.Food {
		for f := 1; f < cohort.Count; f++ {
			w.AddChild(NewFood(
				w.randomPosition(),
				cohort.Energy,
			))
		}
	}

	for _, cohort := range config.Organisms {
		for o := 1; o < cohort.Count; o++ {
			w.AddChild(NewOrganism(
				cohort.Genome,
				cohort.Energy,
				w.randomPosition(),
			))
		}
	}

	return w
}

// func New(config Config) *World {
// 	config = normalizeConfig(config)

// 	// organismMap := grid.New[uint32, organism.Organism](grid.Size[uint32]{
// 	// 	W: uint32(config.Size.W),
// 	// 	H: uint32(config.Size.H),
// 	// })
// 	// for o := 0; o < int(config.NumberOfOrganisms); o++ {
// 	// 	position := organismMap.RandomFreePosition(organismMap.Min(), organismMap.Max())
// 	// 	organism := organism.New(&config.OrganismConfig)
// 	// 	organismMap.Set(position, organism)
// 	// }

// 	// foodMap := grid.New[uint32, food.Food](grid.Size[uint32]{
// 	// 	W: uint32(config.Size.W),
// 	// 	H: uint32(config.Size.H),
// 	// })
// 	// for f := 0; f < int(config.NumberOfOrganisms); f++ {
// 	// 	position := foodMap.RandomFreePosition(foodMap.Min(), foodMap.Max())
// 	// 	food := food.New(d2.Point{X: 1, Y: 1}, 25.0)
// 	// 	foodMap.Set(position, food)
// 	// }

// 	return &World{
// 		Width:             config.Width,
// 		Height:            config.Height,
// 		NumberOfOrganisms: config.NumberOfOrganisms,
// 		OrganismConfig:    config.OrganismConfig,
// 	}
// }

//func (w *World) Update(delta time.Duration) {
// for position, o := range world.OrganismMap.Registry() {

// 	o.Update()

// 	// organism dies: removed from map and skipped on rebuilding
// 	// the registry in the last step of this loop
// 	if o.Energy() <= 0 {
// 		world.OrganismMap.Unset(position)

// 		existingFood, _ := world.FoodMap.Get(position)
// 		if existingFood != nil {
// 			existingFood.IncreaseEnergy(25.0)
// 		} else {
// 			newFood := food.New(d2.Point{X: 1, Y: 1}, 25.0)
// 			world.FoodMap.Set(position, newFood)
// 		}

// 		continue
// 	}

// 	newPosition := grid.Position[uint32]{
// 		X: uint32(int(position.X) + random.Between(-1, 1)),
// 		Y: uint32(int(position.Y) + random.Between(-1, 1)),
// 	}

// 	world.OrganismMap.Move(position, newPosition)
// }
//}

func (w *World) Draw() *ebiten.Image {
	background := ebiten.NewImage(w.size.W, w.size.H)
	background.Fill(color.RGBA{R: 30, G: 30, B: 30, A: 255})

	return background
}

// func translatePosition(renderer *Renderer, position grid.Position[uint32]) ebiten.GeoM {
// 	zoom := renderer.Zoom
// 	gutter := renderer.Theme.Gutter

// 	posX := int(position.X)
// 	posY := int(position.Y)

// 	geoM := ebiten.GeoM{}
// 	geoM.Scale(float64(renderer.Zoom), float64(renderer.Zoom))
// 	geoM.Translate(
// 		float64(posX*zoom)+float64(posX*int(gutter)),
// 		float64(posY*zoom)+float64(posY*int(gutter)),
// 	)

// 	return geoM
// }

func (w *World) GetDimensions() engine.Size {
	return w.size
}

func (w *World) randomPosition() engine.Position {
	return engine.Position{
		X: random.IntBetween(0, w.size.W),
		Y: random.IntBetween(0, w.size.H),
	}
}
