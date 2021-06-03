package regression

// Linear represents a basic linear regression model.
type Linear struct {
	DataSet
	w0, w1 float64
}

// DataSet represents a set of data for regression model.
type DataSet struct {
	DataPoints []DataPoint
	sumX, sumY float64
}

// DataPoint represents (X, Y).
type DataPoint struct {
	X, Y float64
}

// DataFeature represents basic features of a number array.
type DataFeature struct {
	Min, Max, Mean, SD float64
}
