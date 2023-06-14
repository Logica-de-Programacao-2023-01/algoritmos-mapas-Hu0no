package main

import "fmt"

func Divide(x map[string]float64) map[string]float64 {
	y := make(map[string]float64)
	sum := 0.0
	Number_people := 0.0
	for _, value := range x {
		sum += value
		Number_people++
	}
	average := sum / Number_people
	for person, value := range x {
		y[person] = value - average
	}
	return y
}
func main() {
	x := map[string]float64{
		"Marcos": 40,
		"João":   70,
		"Paulo":  55,
		"Jonas":  20,
	}
	fmt.Print(Divide(x))
	//Valor positivo significa que a pessoa deve receber o dinheiro e negativo que ela precisa pagar o dinheiro!
}
