package strapout

import (
	"math"
	"testing"

	"rrecsulator.com/dataSets"
	"rrecsulator.com/standards"
)

func SampleStrapoutData() StrapoutData {
	fd := dataSets.FixedData{
		DrivePOV:    true,
		RegCurbBox:  10,
		RegSdwkBox:  10,
		RegOtherBox: 10,
		CBUBoxes:    10,
		CENTBoxes:   10,
		NPUBoxes:    10,
		DETBoxes:    10,
		Bundle:      standards.TWO_BUNDLE,
	}
	dd := dataSets.DailyData{
		RawLetters:       10,
		WSSLetters:       10,
		BoxholderLetters: 10,
		DPSLetters:       10,
		CasedDPSLetters:  10,

		//cased flats
		RawFlats:        10,
		PreBundledFlats: 10,
		WSSFlats:        10,
		BoxholderFlats:  10,
		CasedDPSFlats:   10,
	}
	sd := StrapoutData{}
	sd.Populate(dd, fd)
	return sd
}

func TestTotalSlots(t *testing.T) {
	sd := SampleStrapoutData()
	want := 70
	got := sd.TotalSlots()

	if got != want {
		t.Errorf("%v got but want %v", got, want)
	}
}

func TestTotalPieces(t *testing.T) {
	sd := SampleStrapoutData()
	want := 100
	got := sd.TotalPieces()

	if got != want {
		t.Errorf("%v got but want %v", got, want)
	}
}

func TestGetHandfulTime(t *testing.T) {
	sd := SampleStrapoutData()
	want := 0.1167
	got := sd.GetHandfulTime()

	if math.Abs(got-want) > 0.001 {
		t.Errorf("%v got but want %v", got, want)
	}
}
func TestGetTime(t *testing.T) {
	sd := SampleStrapoutData()
	want := 1.141
	got := sd.GetReachTime()

	if math.Abs(got-want) > 0.001 {
		t.Errorf("%v got but want %v", got, want)
	}
}
func TestGetRubberBandTime(t *testing.T) {
	sd := SampleStrapoutData()
	want := 0.39666
	got := sd.GetRubberBandTime()

	if math.Abs(got-want) > 0.001 {
		t.Errorf("%v got but want %v", got, want)
	}
}
func TestGetTrayTime(t *testing.T) {
	sd := SampleStrapoutData()
	want := 0.245
	got := sd.GetTrayTime()

	if math.Abs(got-want) > 0.001 {
		t.Errorf("%v got but want %v", got, want)
	}
}
func TestGetReachTime(t *testing.T) {
	sd := SampleStrapoutData()
	want := 1.14167
	got := sd.GetReachTime()

	if math.Abs(got-want) > 0.001 {
		t.Errorf("%v got but want %v", got, want)
	}
}
