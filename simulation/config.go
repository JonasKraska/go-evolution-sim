package simulation

import (
	"time"
)

type Config struct {
	Name string

	Width  uint16
	Height uint16

	FoodCount                  uint16
	FoodEnergy                 Energy
	FoodGrowthRate             float64
	FoodProliferationThreshold uint16

	OrganismCount                  uint16
	OrganismGenes                  uint8
	OrganismEnergy                 Energy
	OrganismMotabolismRate         float64
	OrganismProliferationThreshold uint16

	ElitismThreeshold uint8
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

	if config.FoodEnergy <= 0 {
		config.FoodEnergy = 10
	}

	if config.FoodGrowthRate < 0 {
		config.FoodGrowthRate = 5
	}

	if config.FoodProliferationThreshold <= 0 {
		config.FoodProliferationThreshold = 25
	}

	if config.OrganismGenes <= 0 {
		config.OrganismGenes = 3
	}

	if config.OrganismEnergy <= 0 {
		config.OrganismEnergy = 100
	}

	if config.OrganismMotabolismRate < 0 {
		config.OrganismMotabolismRate = 10
	}

	if config.OrganismProliferationThreshold <= 0 {
		config.OrganismProliferationThreshold = 150
	}

	if config.ElitismThreeshold <= 0 {
		config.ElitismThreeshold = 20
	}

	// print for debugging
	// fmt.Printf("%+v\n", config)

	return config
}
