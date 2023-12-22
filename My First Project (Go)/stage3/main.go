package main

/*
[My First Project - Stage 3/3: Calculate net income](https://hyperskill.org/projects/380/stages/2268/implement)
-------------------------------------------------------------------------------
[Arithmetic operations](https://hyperskill.org/learn/step/16860)
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

	fmt.Println("Staff expenses:")
	var staffExpenses float64
	fmt.Scanln(&staffExpenses)

	fmt.Println("Other expenses:")
	var otherExpenses float64
	fmt.Scanln(&otherExpenses)

	netIncome := totalIncome - staffExpenses - otherExpenses

	fmt.Printf("Net income: $%.1f\n", netIncome)
}
