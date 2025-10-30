package main

func Check(A, B Family) (float64, float64) {
	if distance(A.Position, B.Position) < Eating_Threshold && (A.species.Type == "predator" && B.species.Type == "prey") {
		return A.species.ContactGrowthRate, B.species.ContactGrowthRate
	}

	return 0.0, 0.0
}
