package regression

import (
	"testing"
)

func TestLinear_Predict(t *testing.T) {
	testData := []DataPoint {
		{3, 30},
		{8, 57},
		{9, 64},
		{13, 72},
		{3, 36},
		{6, 43},
		{11, 59},
		{21, 90},
		{1, 20},
		{16, 83},
	}

	var l Linear
	l.Add(testData...)
	l.Train()

	avgX, avgY := l.avgXY()

	if l.w1 != 3.5374756199498467 {
		t.Error("w1 gets wrong value: expect 3.5, get", l.w1)
	}

	if l.w0 != 23.208971858456394 {
		t.Error("w0 gets wrong value: expect 23.6, get", l.w0)
	}

	if avgX != 9.1 {
		t.Error("avgXY() gets wrong X average value: expect 9.1, get", avgX)
	}
	if avgY != 55.4 {
		t.Error("avgXY() gets wrong X average value: expect 55.4, get", avgY)
	}

	p := l.Predict(10)
	if p != 58.58372805795486 {
		t.Error("Predict() gets wrong value: expect 58.6, get", p)
	}
}
