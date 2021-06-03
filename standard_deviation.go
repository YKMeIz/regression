package regression

import "math"

// SD calculates standard deviation of values Y in dataset.
func (ds *DataSet) SD() float64 {
	mean := ds.sumY / float64(len(ds.DataPoints))

	var sigmaPow2 float64
	for _, v := range ds.DataPoints {
		sigmaPow2 += math.Pow(v.Y-mean, 2)
	}

	return math.Sqrt(sigmaPow2 / float64(len(ds.DataPoints)))
}

func sd(data []float64, mean float64) float64 {
	var sigmaPow2 float64
	for _, v := range data {
		sigmaPow2 += math.Pow(v-mean, 2)
	}

	return math.Sqrt(sigmaPow2 / float64(len(data)))
}
