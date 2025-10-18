package hopfield

import "math"

type Hopfield struct {
	Weights [][]float64
	Size    int
}

func NewHopfield(size int) *Hopfield {
	w := make([][]float64, size)
	for i := range w {
		w[i] = make([]float64, size)
	}
	return &Hopfield{Weights: w, Size: size}
}

// Entrenamos con la regla de Hebb: w_ij = Σ(p_i * p_j)
func (h *Hopfield) Train(patterns [][]int) {
	for _, p := range patterns {
		for i := 0; i < h.Size; i++ {
			for j := 0; j < h.Size; j++ {
				if i != j {
					h.Weights[i][j] += float64(p[i] * p[j])
				}
			}
		}
	}
	for i := 0; i < h.Size; i++ {
		for j := 0; j < h.Size; j++ {
			h.Weights[i][j] /= float64(h.Size)
		}
	}
}

func (h *Hopfield) Run(input []int) []int {
	state := make([]int, h.Size)
	copy(state, input)

	changed := true

	for changed {
		changed = false
		for i := 0; i < h.Size; i++ {
			sum := 0.0
			for j := 0; j < h.Size; j++ {
				sum += h.Weights[i][j] * float64(state[j])
			}
			newState := 1
			if sum < 0 {
				newState = -1
			}
			if newState != state[i] {
				state[i] = newState
				changed = true
			}
		}
	}
	return state
}

// Calculamos centro aproximado de aro "C"
func CenterOfPattern(pattern []int) (float64, float64) {
	sumX, sumY, count := 0.0, 0.0, 0.0
	for i := range 100 {
		if pattern[i] == 1 {
			x := float64(i % 10)
			y := float64(i / 10)
			sumX += x
			sumY += y
			count++
		}
	}
	if count == 0 {
		return math.NaN(), math.NaN()
	}
	return math.Round(sumX / count), math.Round(sumY / count)
}
