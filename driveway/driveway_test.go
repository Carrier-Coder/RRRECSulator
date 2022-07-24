package driveway

import (
	"math"
	"testing"
)

func sampleDrivewayData() DrivewayData {
	return DrivewayData{
		LessSwimPool:     1,
		SwimPool:         1,
		FootballField:    1,
		TwoFootballField: 1,
		QuarterMile:      1,
		HalfMile:         1,
	}
}

func TestGetTime(t *testing.T) {
	dw := sampleDrivewayData()
	got := dw.GetTime()
	want := 4.71

	if math.Abs(got-want) > .001 {
		t.Errorf("%v got but want %v", got, want)
	}
}
