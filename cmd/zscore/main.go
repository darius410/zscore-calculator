package main

import (
	"flag"
	"fmt"
	"os"
	"github.com/darius410/zscore-calculator/pkg/stats"
)

// Set these during build using -ldflags
var (
	version = "1.0"    // default value
	commit  = "none"   // default value
	date    = "unknown" // default value
)

func main() {
	// Version flag
	versionFlag := flag.Bool("version", false, "Print version information")
	
	// Existing flags
	value := flag.Float64("value", 0, "The value to calculate the z-score for")
	mean := flag.Float64("mean", 0, "The mean of the distribution")
	stddev := flag.Float64("stddev", 1, "The standard deviation of the distribution")
	flag.Parse()

	// Handle version flag
	if *versionFlag {
		fmt.Printf("zscore version: %s", version,)
		os.Exit(0)
	}

	// Rest of your existing code...
	if *stddev == 0 {
		fmt.Println("Error: Standard deviation cannot be zero")
		os.Exit(1)
	}

	z := stats.ZScore(*value, *mean, *stddev)
	fmt.Printf("Z-Score: %.4f\n", z)
}
