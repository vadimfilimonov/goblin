package goblin

import "fmt"

// Computes the mean of the values in slice. If slice is empty, error is returned.
func Mean(numbers []float64) (float64, error) {
	size := float64(len(numbers))

	if size == 0 {
		return 0, fmt.Errorf("cannot detect a mean value in an empty slice")
	}

	sum := 0.0

	for _, number := range numbers {
		sum += number
	}

	mean := sum / size

	return mean, nil
}
