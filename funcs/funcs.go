package funcs

import "math"

func Average(ints []int) float64 {
	var res float64
	for i := 0; i < len(ints); i++ {
		res += float64(ints[i])
	}
	return res / float64(len(ints))
}

func Median(ints []int) float64 {
	ints = Sort(ints)
	if len(ints)%2 == 0 {
		return (float64(ints[len(ints)/2]) + float64(ints[len(ints)/2-1])) / 2
	}
	return float64(ints[len(ints)/2])
}

func Variance(ints []int) float64 {
	var res float64
	for i := 0; i < len(ints); i++ {
		res += (float64(ints[i]) - Average(ints)) * (float64(ints[i]) - Average(ints))
	}
	return res / float64(len(ints))
}

func StandardDeviation(ints []int) float64 {
	return math.Sqrt(float64(Variance(ints)))
}

func Sort(sl []int) []int {
	for i := 0; i < len(sl)-1; i++ {
		for j := 0; j < len(sl)-i-1; j++ {
			if sl[j] > sl[j+1] {
				sl[j], sl[j+1] = sl[j+1], sl[j]
			}
		}
	}
	return sl
}
