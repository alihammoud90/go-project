package main

import "fmt"

type Numbers struct {
	nb1, nb2 float64
}

type Calculate interface {
	sum() float64
	subtract() float64
	product() float64
	division() float64
}

func(n Numbers) sum() float64 {
	return n.nb1 + n.nb2
}

func(n Numbers) subtract() float64 {
	return n.nb1 - n.nb2
}

func(n Numbers) product() float64 {
	return n.nb1 * n.nb2
}

func(n Numbers) division() float64 {
	return n.nb1 / n.nb2
}

func getResult(calculate Calculate, op int) float64 {
	switch op {
	case 1:
		return calculate.sum()
	case 2:
		return calculate.subtract()
	case 3:
		return calculate.product()
	case 4:
		return calculate.division()
	default:
		fmt.Println("Invalid Operation")
		return 0
	}
}

func main() {
	fmt.Print("Enter First Number: ")
	var nb1 float64
	fmt.Scanln(&nb1)
	fmt.Print("Enter Second Number: ")
	var nb2 float64
	fmt.Scanln(&nb2)
	nbs := Numbers{nb1: nb1, nb2: nb2}
	fmt.Println("Enter operator: ")
	fmt.Println("1) Sum ")
	fmt.Println("2) Subtract ")
	fmt.Println("3) Product ")
	fmt.Println("4) Division ")
	var op int
	fmt.Scanln(&op)
	result := getResult(nbs, op)
	fmt.Printf("Result: %.1f", result)
}