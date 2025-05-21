package stats

// ZScore calculates the standard score (z-score) of a value
// given the mean and standard deviation of the distribution
func ZScore(value, mean, stddev float64) float64 {
	return (value - mean) / stddev
}
