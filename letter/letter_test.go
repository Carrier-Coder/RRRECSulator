package letter

import (
	"math"
	"testing"
)

func sampleLetterData() LetterData {
	return LetterData{
		RandomLetters: 100,
		WSSLetters:    100,

		DPSLetters: 100,
	}
}

// check letter calc against hand calculation
func TestCalc(t *testing.T) {
	ld := sampleLetterData()
	got := ld.GetTime()
	want := 13.05
	if math.Abs(got-want) > 0.001 {
		t.Errorf("%.4f New Time does not match hand calc %.4f", got, want)
	}
}

func TestDetectCaseDPS(t *testing.T) {
	ld := sampleLetterData()
	got := ld.caseDPS()
	want := true
	if got != want {
		t.Errorf("%v got but want %v", got, want)
	}

	ld.DPSLetters = 401
	got = ld.caseDPS()
	want = false
	if got != want {
		t.Errorf("%v got but want %v", got, want)
	}
}

func TestTotalLetters(t *testing.T) {
	ld := sampleLetterData()
	ld.CasedDPSLetters = 100

	got := ld.TotalLetters()
	want := 400

	if got != want {
		t.Errorf("%v got but wanted %v", got, want)
	}
}

func TestCasedTime(t *testing.T) {
	ld := LetterData{}
	ld.CasedDPSLetters = 100

	got := ld.GetTime()
	want := 2.94

	if math.Abs(got-want) > 0.001 {
		t.Errorf("%.4f New Time does not match hand calc %.4f", got, want)

	}
}

func TestCasedLetters(t *testing.T) {
	ld := sampleLetterData()

	got := ld.CasedLetters()
	want := 300

	if got != want {
		t.Errorf("%v got but wanted %v", got, want)
	}

	ld.DPSLetters = 401
	got = ld.CasedLetters()
	want = 200
	if got != want {
		t.Errorf("%v got but wanted %v", got, want)
	}

	ld = sampleLetterData()
	ld.CasedDPSLetters = 100
	got = ld.CasedLetters()
	want = 400

	if got != want {
		t.Errorf("%v got but wanted %v", got, want)
	}

}
