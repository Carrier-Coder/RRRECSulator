package accountable

import (
	"testing"
)

func sampleAccountableData() AccountableData {
	return AccountableData{

		Certified:           10,
		Registered:          10,
		SignatureExpress:    10,
		NonSignatureExpress: 10,

		CODs:             10,
		PostageDueCustom: 10,

		PostageDue: 10,
	}
}

func TestTotal(t *testing.T) {
	ad := sampleAccountableData()
	got := ad.Total()
	want := 70
	if got != want {
		t.Errorf("%v got but want %v", got, want)
	}
}

func TestCalc(t *testing.T) {
	ad := sampleAccountableData()
	got := ad.GetTime()

	want := 155.1478
	if got-want > 0.001 || got-want < -0.001 {
		t.Errorf("%v got but want %v", got, want)
	}

	ad = AccountableData{
		PostageDue: 1,
	}
	got = ad.GetTime()
	want = 1.689
	if got-want > 0.001 || got-want < -0.001 {
		t.Errorf("%v got but want %v", got, want)
	}
}
