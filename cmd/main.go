package main

import (
	"fmt"
	"github.com/OakAnderson/academia/services"
)

func main() {
	fmt.Printf("Média de peso: %.2f\n", services.MediaPeso())
	fmt.Printf("Cliente que mais perdeu peso:\n%v\n", services.ClienteQueMaisPerdeuPeso())
}
