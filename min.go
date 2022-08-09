package goblin

import "fmt"

// Min computes the minimum value of `slice`. If `slice` is empty, error is returned
func Min(numbers []float64) (float64, error) {
	if len(numbers) == 0 {
		return 0, fmt.Errorf("cannot detect a minimum value in an empty slice")
	}

	min := numbers[0]

	for _, number := range numbers {
		if number < min {
			min = number
		}
	}

	return min, nil
}
