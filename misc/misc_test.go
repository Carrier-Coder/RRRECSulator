package misc

import (
	"math"
	"testing"
)

// check calc against hand calculation

func sampleMiscData() MiscData {
	return MiscData{
		Form3982:    1,
		BlueBox:     1,
		OneStepScan: 10,
		TwoStepScan: 10,

		StampSale:  10,
		RuralReach: 5,
	}
}

func TestGetTime(t *testing.T) {
	md := sampleMiscData()
	got := md.GetTime()
	want := 38.2333
	if math.Abs(got-want) > 0.001 {
		t.Errorf("%v got but want %v", got, want)
	}
}

func TestGetTripReportTime(t *testing.T) {
	md := sampleMiscData()
	got := md.GetTripReportTime()
	want := 0.491
	if math.Abs(got-want) > 0.001 {
		t.Errorf("%v got but want %v", got, want)
	}
}
func TestGetTrayStorageTime(t *testing.T) {
	md := sampleMiscData()
	got := md.GetTrayStorageTime()
	want := 0.2643
	if math.Abs(got-want) > 0.001 {
		t.Errorf("%v got but want %v", got, want)
	}
}
func TestGetScannerSetupTime(t *testing.T) {
	md := sampleMiscData()
	got := md.GetScannerSetupTime()
	want := 0.1882
	if math.Abs(got-want) > 0.001 {
		t.Errorf("%v got but want %v", got, want)
	}
}
func TestGetStampSaleTime(t *testing.T) {
	md := sampleMiscData()
	got := md.GetStampSaleTime()
	want := 7.140
	if math.Abs(got-want) > 0.001 {
		t.Errorf("%v got but want %v", got, want)
	}
}
func TestGetRuralReachTime(t *testing.T) {
	md := sampleMiscData()
	got := md.GetRuralReachTime()
	want := 26.229
	if math.Abs(got-want) > 0.001 {
		t.Errorf("%v got but want %v", got, want)
	}
}
func TestGetBlueBoxTime(t *testing.T) {
	md := sampleMiscData()
	got := md.GetBlueBoxTime()
	want := 1.7928
	if math.Abs(got-want) > 0.001 {
		t.Errorf("%v got but want %v", got, want)
	}
}
func TestGetOneStepScanTime(t *testing.T) {
	md := sampleMiscData()
	got := md.GetOneStepScanTime()
	want := 0.795
	if math.Abs(got-want) > 0.001 {
		t.Errorf("%v got but want %v", got, want)
	}
}
func TestGetTwoStepScanTime(t *testing.T) {
	md := sampleMiscData()
	got := md.GetTwoStepScanTime()
	want := 0.928
	if math.Abs(got-want) > 0.001 {
		t.Errorf("%v got but want %v", got, want)
	}
}

func TestGetForm3982Time(t *testing.T) {
	md := sampleMiscData()
	got := md.GetForm3982Time()
	want := 0.405
	if math.Abs(got-want) > 0.001 {
		t.Errorf("%v got but want %v", got, want)
	}
}
