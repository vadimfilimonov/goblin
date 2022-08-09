package goblin

import "fmt"

// Max computes the maximum value of `slice`. If `slice` is empty, error is returned.
func Max(numbers []float64) (float64, error) {
	if len(numbers) == 0 {
		return 0, fmt.Errorf("cannot detect a maximum value in an empty slice")
	}

	max := numbers[0]

	for _, number := range numbers {
		if number > max {
			max = number
		}
	}

	return max, nil
}
