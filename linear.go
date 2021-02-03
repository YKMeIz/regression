package regression

import "math"

type Linear struct {
	DataSet
	w0, w1 float64
}

type DataSet struct {
	DataPoints []DataPoint
	sumX, sumY float64
}

type DataPoint struct {
	X, Y float64
}

func (ds *DataSet) Add(dp ...DataPoint)  {
	for _, v := range dp {
		ds.sumX += v.X
		ds.sumY += v.Y
	}
	ds.DataPoints = append(ds.DataPoints, dp...)
}

func (ds *DataSet) avgXY() (float64, float64) {
	return ds.sumX / float64(len(ds.DataPoints)), ds.sumY / float64(len(ds.DataPoints))
}

func (l *Linear) Train() {
	avgX, avgY := l.avgXY()

	var a, b float64

	for _, v := range l.DataPoints {
		a += (v.X - avgX) * (v.Y - avgY)
		b += math.Pow(v.X - avgX, 2)
	}

	l.w1 = a / b
	l.w0 = avgY - (l.w1 * avgX)
}

func (l *Linear) Predict(x float64) float64 {
	return l.w0 + (l.w1 * x)
}