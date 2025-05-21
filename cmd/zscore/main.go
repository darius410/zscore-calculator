package main

import (
	"flag"
	"fmt"
	"os"
	"github.com/darius410/zScoreCalculator/pkg/stats"
)

func main() {
	// Define command-line flags
	value := flag.Float64("value", 0, "The value to calculate the z-score for")
	mean := flag.Float64("mean", 0, "The mean of the distribution")
	stddev := flag.Float64("stddev", 1, "The standard deviation of the distribution")
	flag.Parse()

	// Validate inputs
	if *stddev == 0 {
		fmt.Println("Error: Standard deviation cannot be zero")
		os.Exit(1)
	}

	// Calculate z-score
	z := stats.ZScore(*value, *mean, *stddev)
	fmt.Printf("Z-Score: %.4f\n", z)
}
