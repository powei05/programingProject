package main

import (
	"math"
)

// for the purpose to detect whether prey and predator will meet
// we need to calculate the closest approach between two moving families
// we can use relative movement to simplify the calculation
// we aim to find the time t when the distance between two families is the smallest
// if the distance at that time is less than a certain threshold, we consider they will meet
// if the time  < 0 or > interval time, the closest approach happens at the beginning or the end of the interval
// to find the closest approach, we can use the differnential calculus
// let D(t) = distance^2 between two families at time t
// D(t) = (P0.x + Vel.x * t)^2 + (P0.y + Vel.y * t)^2
// dD/dt = 2*(P0.x + Vel.x * t)*Vel.x + 2*(P0.y + Vel.y * t)*Vel.y
// set dD/dt = 0, we can solve for t
// t = -(P0.x * Vel.x + P0.y * Vel.y) / (Vel.x^2 + Vel.y^2)
// then we can calculate the distance at time t
// D(t) = (P0.x + Vel.x * t)^2 + (P0.y + Vel.y * t)^2
// if D(t) < threshold^2, we consider they will meet
func calculateRelativeMovement(A, B *Family) (P0, Vel OrderedPair) {
	P0.x = A.Position.x - B.Position.x
	P0.y = A.Position.y - B.Position.y

	Vel.x = A.MovementDirection.x - B.MovementDirection.x
	Vel.y = A.MovementDirection.y - B.MovementDirection.y

	return P0, Vel
}

func ClosestTime(A, B *Family, maxTime float64) float64 {
	P0, Vel := calculateRelativeMovement(A, B)

	//  D^2(t) = a*t^2 + b*t + c
	a := Vel.x*Vel.x + Vel.y*Vel.y

	b := 2 * (P0.x*Vel.x + P0.y*Vel.y)

	//we don't need c for finding the closest time
	//c := P0.x*P0.x + P0.y*P0.y

	// if a == 0, means Vel.x == 0 and Vel.y == 0
	// means two families are not moving relative to each other
	// the distance is constant, so the closest time is 0
	if a < 1e-9 {
		return 0.0 // any time is the closest time, return 0
	}

	//the solution to dD/dt = 0 is:
	tStar := -b / (2 * a)

	// check the value of tStar to determine the closest time within [0, maxTime]

	// case A: t* is within [0, maxTime]
	if tStar >= 0 && tStar <= maxTime {
		// the minimum distance occurs at tStar
		return tStar
	}

	// case B: t* is outside the interval
	if tStar < 0 {
		// the minimum distance occurs in the past, so within [0, maxTime], the distance is increasing.
		// the minimum distance occurs at t=0
		return 0.0
	}

	// case C: t* > maxTime
	// the minimum distance occurs in the future, so within [0, maxTime], the distance is decreasing.
	// the minimum distance occurs at t=maxTime
	return maxTime
}

func DistanceAtTime(P0, Vel OrderedPair, t float64) float64 {
	relX := P0.x + Vel.x*t
	relY := P0.y + Vel.y*t

	return math.Sqrt(relX*relX + relY*relY)

}

func WillMeet(A, B *Family, intervalTime, threshold float64) bool {
	P0, Vel := calculateRelativeMovement(A, B)
	closestTime := ClosestTime(A, B, intervalTime)

	distance := DistanceAtTime(P0, Vel, closestTime)
	return distance < threshold
}

// make a map to store species types
func initalSpeciesMap() map[uint32]map[string]bool {
	speciesMap := make(map[uint32]map[string]bool)
	speciesMap[0] = make(map[string]bool)
	speciesMap[1] = make(map[string]bool)
	speciesMap[2] = make(map[string]bool)
	for _, name := range preySpeciesNames {
		speciesMap[0][name] = true // prey species
	}
	for _, name := range predatorSpeciesNames {
		speciesMap[1][name] = true // predator species
	}
	for _, name := range neutralSpeciesNames {
		speciesMap[2][name] = true // neutral species
	}
	return speciesMap
}

func NumOfSpeciesUpdated(allFamilies []*Family, allSpecies []*Species, speciesmap map[uint32]map[string]bool) {
	numofFamilies := len(allFamilies)

	for i := 0; i < numofFamilies; i++ {
		for j := 0; j < numofFamilies; j++ {
			if speciesmap[1][allFamilies[i].species.Name] == true && speciesmap[0][allFamilies[j].species.Name] == true {
				// species i is predator, species j is prey
				if WillMeet(allFamilies[i], allFamilies[j], 1.0, Eating_Threshold) {
					// update population of prey species
					// we use probability to determine how many prey are eaten
					preyEaten := math.Floor(allFamilies[j].species.ContactGrowthRate * float64(allFamilies[j].Size))
					predatorGrowth := math.Floor(allFamilies[i].species.GrowthRate * float64(allFamilies[i].Size))
					allFamilies[j].Size -= int(preyEaten)
					allFamilies[i].Size += int(predatorGrowth)

				}
			}
		}
	}
}
