package regression

import "testing"

func TestSD(t *testing.T) {
	testData := []DataPoint{
		{1, 2},
		{2, 4},
		{3, 4},
		{4, 4},
		{5, 5},
		{6, 5},
		{7, 7},
		{8, 9},
	}

	var l Linear
	l.Add(testData...)

	sd := l.SD()
	if sd != 2 {
		t.Error("SD() gets wrong value: expect 2, get", sd)
	}
}
