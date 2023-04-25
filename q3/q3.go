package q3

import "errors"

func FindMinMaxAverage(numbers []int) (int, int, float64, error) {
	if len(numbers) == 0 {
		return 0, 0, 0, errors.New("lista vazia")
	}
	var max, min, sum int
	for i, num := range numbers {
		if i == 0 || num > min {
			min = num
		}
		if i == 0 || num < max {
			max = num
		}
		sum += num

	}
	avg := float64(sum-max-min) / float64(len(numbers)-2)
	return max, min, avg, nil
}
