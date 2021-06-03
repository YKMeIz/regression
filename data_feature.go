package regression

func Mean(data []float64) float64 {
	m := 0.0
	for _, v := range data {
		m += v
	}
	return m / float64(len(data))
}

func Feature(data []float64) DataFeature {
	var df DataFeature
	for _, v := range data {
		df.Mean += v
		if v > df.Max {
			df.Max = v
		}
		if df.Min == 0 || v < df.Min {
			df.Min = v
		}
	}

	df.Mean = df.Mean / float64(len(data))
	df.SD = sd(data, df.Mean)

	return df
}
