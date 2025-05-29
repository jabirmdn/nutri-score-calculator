package main


type ScoreType int 

const (
	Food ScoreType iota
    Beverage 
	Water
	Cheese	
)

type EnergyKG float64
type SugarGram float64
type SaturatedFattyAcids float64
type SodiumMilligram float64
type FruitsPercentage float64
type FibreGram float64
type ProteinGram float64

type NutritionalData struct {
	Energy              EnergyKG
	Sugars              SugarGram
	SaturatedFattyAcids SaturatedFattyAcids
	Sodium              SodiumMilligram
	Fruits              FruitsPercentage
	Fibre               FibreGram
	Protein             ProteinGram
	IsWater             bool
	// isWater bool
}
