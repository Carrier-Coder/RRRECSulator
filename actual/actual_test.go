package actual

import (
	"math"
	"testing"
)

func sampleActualData() ActualData {
	return ActualData{

		SafetyServiceTalks: 1.0,
		Loading:            2.0,
		Deviation:          3.0,
		EndOfShift:         4.0,
	}
}

func TestGetSafetyServiceTalkTime(t *testing.T) {
	ac := sampleActualData()
	got := ac.GetSafetyServiceTalkTime()
	want := 1.0323
	if math.Abs(got-want) > 0.001 {
		t.Errorf("%v got but want %v", got, want)
	}
}
func TestGetLoadingTime(t *testing.T) {
	ac := sampleActualData()
	got := ac.GetLoadingTime()
	want := 2.0646
	if math.Abs(got-want) > 0.001 {
		t.Errorf("%v got but want %v", got, want)
	}
}
func TestGetEndOfShiftTime(t *testing.T) {
	ac := sampleActualData()
	got := ac.GetEndOfShiftTime()
	want := 4.0
	if math.Abs(got-want) > 0.001 {
		t.Errorf("%v got but want %v", got, want)
	}
}
func TestGetDeviationTime(t *testing.T) {
	ac := sampleActualData()
	got := ac.GetDeviationTime()
	want := 3.0969
	if math.Abs(got-want) > 0.001 {
		t.Errorf("%v got but want %v", got, want)
	}
}

func TestGetTime(t *testing.T) {
	ac := sampleActualData()
	got := ac.GetTime()
	want := 10.1938
	if got != want {
		t.Errorf("%v got but want %v", got, want)
	}
}
