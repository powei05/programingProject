package main

type Species struct {
	Name              string
	Class             string
	GrowthRate        float64
	ContactGrowthRate float64
}

type Family struct {
	Size              int
	MovementSpeed     OrderedPair
	Position          OrderedPair
	MovementDirection OrderedPair
	species           Species
}
type OrderedPair struct {
	x float64
	y float64
}

var preySpeciesNames = []string{"Rabbit", "Sheep", "Deer"}
var predatorSpeciesNames = []string{"Wolf"}
var neutralSpeciesNames = []string{"Human"}

var Population_Rabbit = 100
var Population_Sheep = 50
var Population_Deer = 20
var Population_Wolf = 10
var Population_Human = 2

const GrowthRate_Rabbit = 0.5
const GrowthRate_Sheep = 0.3
const GrowthRate_Deer = 0.2
const GrowthRate_Wolf = 0.1
const PlantCoefficient = 0.05
const Merging_Threshould = 15
const Eating_Threshold = 5.0
