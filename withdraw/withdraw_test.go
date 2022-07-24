package withdraw

import (
	"math"
	"testing"
)

func sampleWithdrawData() WithdrawData {
	return WithdrawData{

		DPSLetters: 500,
		DPSFlats:   200,

		Withdraw: true,
	}
}

func TestNoLetterNoFlat(t *testing.T) {
	wd := sampleWithdrawData()
	wd.DPSLetters = 0
	wd.DPSFlats = 0

	got := wd.GetTime()
	want := 1.0517
	if math.Abs(got-want) > 0.001 {
		t.Errorf("%v got but want %v", got, want)
	}
}

func TestNoFlat(t *testing.T) {
	wd := sampleWithdrawData()
	wd.DPSLetters = 1
	wd.DPSFlats = 0

	got := wd.GetTime()
	want := 2.0706
	if math.Abs(got-want) > 0.001 {
		t.Errorf("%v got but want %v", got, want)
	}
}

func TestNoLetter(t *testing.T) {
	wd := sampleWithdrawData()
	wd.DPSLetters = 0
	wd.DPSFlats = 1

	got := wd.GetTime()
	want := 1.422
	if math.Abs(got-want) > 0.001 {
		t.Errorf("%v got but want %v", got, want)
	}
}

func TestGetTime(t *testing.T) {
	wd := sampleWithdrawData()
	got := wd.GetTime()

	//hand calc from data above
	want := 2.8959

	if math.Abs(got-want) > 0.001 {
		t.Errorf("%v got but want %v", got, want)
	}
}
