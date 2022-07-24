package averaged

import (
	"math"
	"testing"
)

func sampleAveragedData() AveragedData {
	return AveragedData{
		Boxes: 600,
	}
}

func TestGetParcelAcceptedTime(t *testing.T) {
	ad := sampleAveragedData()
	got := ad.GetParcelAcceptedTime()
	want := 0.0124
	if math.Abs(got-want) > 0.001 {
		t.Errorf("%v got but want %v", got, want)
	}
}

func TestGetRegCertAcceptedTime(t *testing.T) {
	ad := sampleAveragedData()
	got := ad.GetRegCertAcceptedTime()
	want := 0.005317
	if math.Abs(got-want) > 0.001 {
		t.Errorf("%v got but want %v", got, want)
	}
}

func TestGetMoneyOrderAcceptedTime(t *testing.T) {
	ad := sampleAveragedData()
	got := ad.GetMoneyOrderAcceptedTime()
	want := 0.011767
	if math.Abs(got-want) > 0.001 {
		t.Errorf("%v got but want %v", got, want)
	}
}

func TestGetTime(t *testing.T) {
	ad := sampleAveragedData()
	got := ad.GetTime()
	want := 0.029484
	if math.Abs(got-want) > 0.001 {
		t.Errorf("%v got but want %v", got, want)
	}
}
