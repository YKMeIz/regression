package regression

import "testing"

func TestDataFeature(t *testing.T) {
	testData := []float64{2, 4, 4, 4, 5, 5, 7, 9}

	df := Feature(testData)

	if df.SD != 2 {
		t.Error("SD() gets wrong value: expect 2, get", df.SD)
	}

	if df.Mean != 5 {
		t.Error("Mean() gets wrong value: expect 5, get", df.Mean)
	}

	if df.Min != 2 {
		t.Error("Min() gets wrong value: expect 2, get", df.Min)
	}

	if df.Max != 9 {
		t.Error("Max() gets wrong value: expect 5, get", df.Max)
	}
}
