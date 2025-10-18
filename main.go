package main

import (
	"fmt"

	"github.com/JMRodriguez-work/ia_tp3/hopfield"
	"github.com/JMRodriguez-work/ia_tp3/patterns"
)

func main() {
	size := 100

	pats := patterns.GetTrainingPatterns()

	net := hopfield.NewHopfield(size)
	net.Train(pats)

	test := make([]int, len(pats[0]))
	copy(test, pats[1])
	test[20] = -1
	test[70] = 1

	fmt.Println("\nPatrón ingresado:")
	hopfield.PrintMatrix(test)

	result := net.Run(test)

	fmt.Println("\n\nPatrón reconocido:")
	hopfield.PrintMatrix(result)

	x, y := hopfield.CenterOfPattern(result)
	fmt.Printf("Centro estimado del aro (X,Y): (%.0f, %.0f)\n", x, y)

	fmt.Println("\nPresiona Enter para salir...")
	fmt.Scanln()
}
