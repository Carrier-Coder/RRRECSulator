package pickup

import (
	"testing"
)

func samplePickupData() PickupData {
	return PickupData{

		EventForms: 1,
		Events:     2,

		PackageNum: 10,
	}
}

func TestGetTime(t *testing.T) {
	pd := samplePickupData()
	got := pd.GetTime()
	want := 4.45
	if got-want > 0.001 || got-want < -0.001 {
		t.Errorf("%v got but want %v", got, want)
	}

}
