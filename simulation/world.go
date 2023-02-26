package simulation

import (
	"github.com/JonasKraska/go-evolution-sim/engine"
	"github.com/JonasKraska/go-evolution-sim/engine/random"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

type World struct {
	engine.Game

	width  int
	height int
}

type WorldConfig struct {
	Width     uint32
	Height    uint32
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
		width:  int(config.Width),
		height: int(config.Height),
	}

	for _, cohort := range config.Food {
		for f := 1; f < cohort.Count; f++ {
			w.spawnFood(
				w.randomPosition(),
				cohort.Energy,
			)
		}
	}

	for _, cohort := range config.Organisms {
		for o := 1; o < cohort.Count; o++ {
			w.spawnOrganism(
				w.randomPosition(),
				cohort.Genome,
				cohort.Energy,
			)
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
	background := ebiten.NewImage(w.width, w.height)
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

func (w *World) Contains(position engine.Vector) bool {
	return position.X > 0 && position.Y > 0 && position.X <= float64(w.width) && position.Y <= float64(w.height)
}

func (w *World) spawnOrganism(position engine.Position, genome Genome, energy Energy) {
	organism := NewOrganism(position, genome, energy)

	w.AddChild(organism)

	organism.Register(OrganismDeathHook, func() {
		w.onOrganismDeath(organism)
	})
}

func (w *World) spawnFood(position engine.Position, energy Energy) {
	food := NewFood(position, energy)

	w.AddChild(food)
}

func (w *World) onOrganismDeath(organism *Organism) {
	w.spawnFood(organism.GetPosition(), Energy(5))
}

func (w *World) randomPosition() engine.Position {
	return engine.Position{
		X: random.FloatBetween(0, w.width),
		Y: random.FloatBetween(0, w.height),
	}
}
