package main

import (
	"fmt"

	"github.com/luanlimasoueu/go-ml-kit/regression"
)

func main() {
	modelo := regression.RegressaoLinearSimples{}

	X := []float64{1, 2, 3, 4, 5}
	Y := []float64{2, 4, 5, 4, 5}

	modelo.Fit(X, Y)

	previsoes := modelo.Predict([]float64{6, 7, 8})

	fmt.Println("Previsões:")

	for i, valor := range previsoes {
		fmt.Printf("X = %.2f -> Y = %.2f\n", float64(i+6), valor)
	}
}
