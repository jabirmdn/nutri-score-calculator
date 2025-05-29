package main

import "fmt"

func main() {
	ns := CalculateNutritionalScore(NutritionalData{}, Food)
	fmt.Println(ns)

}
