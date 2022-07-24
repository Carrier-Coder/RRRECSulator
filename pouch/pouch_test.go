package pouch

import (
	"math"
	"testing"
)

func samplePouchData() PouchData {
	return PouchData{

		LowVolume:  2,
		HighVolume: 3,
	}
}

func TestGetTime(t *testing.T) {
	pd := samplePouchData()
	got := pd.GetTime()
	//hand calc from data above
	want := 20.2551

	if math.Abs(got-want) > 0.001 {
		t.Errorf("%v got but want %v", got, want)
	}
}
