package hopfield

import "fmt"

func PrintMatrix(pattern []int) {
	for i := range 100 {
		if pattern[i] == 1 {
			fmt.Print("█")
		} else {
			fmt.Print(" ")
		}
		if (i+1)%10 == 0 {
			fmt.Println()
		}
	}
	fmt.Println()
}
