package systems

import "math/rand"

func WeightedBoolean(trueWeight float64) bool {
	if trueWeight > 1 || trueWeight < 0 {
		panic("boolean weight cannot be more than 1 or less than 0")
	}

	return rand.Float64() < trueWeight
}
