package walk

import (
	"math"
	"testing"

	"rrecsulator.com/standards"
)

func sampleWalkData() WalkData {
	return WalkData{
		Steps:       520,
		FeetPerStep: standards.AVG_FEET_PER_STEP,
	}
}

func TestGetTime(t *testing.T) {
	ww := sampleWalkData()
	got := ww.GetTime()
	want := 5.577

	if math.Abs(got-want) > .001 {
		t.Errorf("%v got but want %v", got, want)
	}
}
