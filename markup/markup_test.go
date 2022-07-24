package markup

import (
	"math"
	"testing"
)

func sampleMarkupData() MarkupData {
	return MarkupData{
		AddressedMailPieces: 2000,
	}
}

func TestGetTime(t *testing.T) {
	md := sampleMarkupData()
	got := md.GetTime()
	want := 1.2699

	if math.Abs(got-want) > .001 {
		t.Errorf("%v got but want %v", got, want)
	}
}
