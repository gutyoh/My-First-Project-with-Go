package main

/*
[My First Project - Stage 2/3: Measuring total income](https://hyperskill.org/projects/380/stages/2267/implement)
-------------------------------------------------------------------------------
[Arithmetic operations](https://hyperskill.org/learn/step/16679)
*/

import "fmt"

func main() {
	bubblegum := 202.0
	toffee := 118.0
	iceCream := 2250.0
	milkChocolate := 1680.0
	doughnut := 1075.0
	pancake := 80.0

	totalIncome := bubblegum + toffee + iceCream + milkChocolate + doughnut + pancake

	fmt.Println("Earned amount:")
	fmt.Printf("Bubblegum: $%.1f\n", bubblegum)
	fmt.Printf("Toffee: $%.1f\n", toffee)
	fmt.Printf("Ice cream: $%.1f\n", iceCream)
	fmt.Printf("Milk chocolate: $%.1f\n", milkChocolate)
	fmt.Printf("Doughnut: $%.1f\n", doughnut)
	fmt.Printf("Pancake: $%.1f\n\n", pancake)

	fmt.Printf("Income: $%.1f\n", totalIncome)
}
