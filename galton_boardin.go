package main

import (
	"fmt"
	"strings"
	"math" 
	"math/rand" 
)

const (
	defaultRows  = 9
	defaultBalls = 256
	maxBalls     = 10000
	maxRows     = 50
)

func main() {
	var rows int
	fmt.Print("Enter the number of rows: ")
	_, err := fmt.Scanln(&rows)
	if err != nil || rows <= 0 || rows > maxRows {
		fmt.Printf("Invalid input for rows. Using default value %d.\n", defaultRows)
		rows = defaultRows
	}

	var balls int
	fmt.Print("Enter the number of balls: ")
	_, err = fmt.Scanln(&balls)
	if err != nil || balls <= 0 || balls > maxBalls {
		fmt.Printf("Invalid input for balls. Using default value %d.\n", defaultBalls)
		balls = defaultBalls
	}

	bins := make([]int, rows+1)
	for i := 0; i < balls; i++ {    
		ballPosition := 0
		for j := 0; j < rows; j++ {
			if rand.Float64() < 0.5 {
				ballPosition++
			}       
		}
		bins[ballPosition]++
	}

	fmt.Println("Distribution of balls:")
	for i := 0; i <= rows; i++ {
		fmt.Printf("Bin %d: %s\n", i, strings.Repeat("~", bins[i]))
	}

	totalBalls := 0
	for _, count := range bins {
		totalBalls += count
	}

	if totalBalls == 0 {
		fmt.Println("No balls were simulated. Cannot calculate statistics.")
		return
	}

	mean := float64(0)	
	for i, count := range bins {
		mean += float64(i) * float64(count)
	}
	mean /= float64(totalBalls)

	variance := float64(0)	
	for i, count := range bins {
		variance += float64(count) * math.Pow(float64(i)-mean, 2)
	}
	variance /= float64(totalBalls)
	stdDev := math.Sqrt(variance)

	fmt.Printf("\nMean: %.2f\n", mean)
	fmt.Printf("Variance: %.2f\n", variance)
	fmt.Printf("Standard Deviation: %.2f\n", stdDev)	
}
