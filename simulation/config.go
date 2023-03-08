package simulation

import "time"

type Config struct {
	Name string

	Width  uint16
	Height uint16

	FoodCount  uint16
	FoodEnergy Energy

	OrganismCount  uint16
	OrganismGenes  uint8
	OrganismEnergy Energy
}

func NewConfig(config Config) Config {
	if config.Name == "" {
		config.Name = "simulation_" + time.Now().Format("2006-01-02") + "_" + time.Now().Format("15:04:05")
	}

	if config.Width <= 0 {
		config.Width = 350
	}

	if config.Height <= 0 {
		config.Height = 225
	}

	if config.FoodCount <= 0 {
		config.FoodCount = 512
	}

	if config.FoodEnergy <= 0 {
		config.FoodEnergy = 50
	}

	if config.OrganismCount <= 0 {
		config.OrganismCount = 256
	}

	if config.OrganismGenes <= 0 {
		config.OrganismGenes = 3
	}

	if config.OrganismEnergy <= 0 {
		config.OrganismEnergy = 100
	}

	return config
}
