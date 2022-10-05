package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to example of slices")

	var vehicles = []string{"Car", "Bike", "Truck"}
	fmt.Printf("Type of vehicles is %T\n", vehicles)

	// append adds from the last in string
	vehicles = append(vehicles, "Scooter", "Bus")
	fmt.Println(vehicles)

	vehicles = append(vehicles[1:3])
	fmt.Println(vehicles)

	highScore := make([]int, 4)

	highScore[0] = 234
	highScore[1] = 945
	highScore[2] = 465
	highScore[3] = 867
	//highScore[4] = 777

	highScore = append(highScore, 555, 666, 321)

	fmt.Println(highScore)

	sort.Ints(highScore)
	fmt.Println(highScore)
	fmt.Println(sort.IntsAreSorted(highScore))
}
