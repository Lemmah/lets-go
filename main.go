package main

import "fmt"

func main() {
	var plantCapacities []float64
	var activePlants = []int{0, 1, 2}
	var gridLoad = 75.
	plantCapacities = []float64{30, 30, 30, 60, 60, 100}
	fmt.Println("Gopher Power🙇‍♂️")
	fmt.Println("1. Generate Power Plant Report")
	fmt.Println("2. Generate Power Grid Report")
	var option int
	fmt.Scanln(&option)
	switch option {
	case 1:
		for idx, cap := range plantCapacities {
			fmt.Printf("Plant %d capacity: %.0f\n", idx+1, cap)
		}
	case 2:
		capacity := 0.
		for _, plantID := range activePlants {
			capacity += plantCapacities[plantID]
		}
		utilization := gridLoad / capacity
		fmt.Printf("%-20s%.0f\n", "Capacity:", capacity)
		fmt.Printf("%-20s%.0f\n", "Grid Load:", gridLoad)
		fmt.Printf("%-20s%.1f%%\n", "Utilization:", utilization*100)
	default:
		fmt.Printf("Invalid Option: %d\n", option)
	}
}
