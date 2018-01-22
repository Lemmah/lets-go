package main

import (
	"fmt"
	"strings"
)

func main() {
	plants := []PowerPlant{
		PowerPlant{hydro, 300, active},
		PowerPlant{wind, 30, active},
		PowerPlant{wind, 25, inactive},
		PowerPlant{wind, 35, active},
		PowerPlant{solar, 45, unavailable},
		PowerPlant{solar, 40, active},
	}

	grid := PowerGrid{300, plants}

	fmt.Println("Gopher Power🙇‍♂️")
	fmt.Println("1. Generate Power Plant Report")
	fmt.Println("2. Generate Power Grid Report")

	var option int
	fmt.Scanln(&option)

	switch option {
	case 1:
		grid.generatePlantReport()
	case 2:
		grid.generateGridReport()
	default:
		fmt.Printf("Invalid Option: %d\n", option)
	}
}

// PlantType Type
type PlantType string

const (
	hydro PlantType = "Hydro"
	wind  PlantType = "Wind"
	solar PlantType = "Solar"
)

// PlantStatus Type
type PlantStatus string

const (
	active      PlantStatus = "Active"
	inactive    PlantStatus = "Inactive"
	unavailable PlantStatus = "Unavailable"
)

// PowerPlant Struct
type PowerPlant struct {
	plantType PlantType
	capacity  float64
	status    PlantStatus
}

// PowerGrid Struct
type PowerGrid struct {
	load   float64
	plants []PowerPlant
}

func (pg *PowerGrid) generatePlantReport() {
	for idx, plant := range pg.plants {
		label := fmt.Sprintf("%s%d", "Plant #", idx)
		fmt.Println(label)
		fmt.Println(strings.Repeat("-", len(label)))
		fmt.Printf("%-20s%s\n", "Type:", plant.plantType)
		fmt.Printf("%-20s%.0f\n", "Capacity:", plant.capacity)
		fmt.Printf("%-20s%s\n", "Status:", plant.status)
	}
}

func (pg *PowerGrid) generateGridReport() {
	capacity := 0.
	for _, plant := range pg.plants {
		if plant.status == active {
			capacity += plant.capacity
		}
	}
	label := "Power Grid Report"
	fmt.Println(label)
	fmt.Println(strings.Repeat("-", len(label)))
	fmt.Printf("%-20s%.0f\n", "Capacity:", capacity)
	fmt.Printf("%-20s%.0f\n", "Grid Load:", pg.load)
	fmt.Printf("%-20s%.1f%%\n", "Utilization:", pg.load/capacity*100)
}
