package dismount

import (
	"math"
	"testing"
)

func sampleDismountData() DismountData {
	return DismountData{

		Other:    1,
		Sidewalk: 1,
		Cbu:      1,
		Cent:     1,
		Npu:      1,
		DetVpo:   1,

		AddTripOther:    1,
		AddTripSidewalk: 1,
		AddTripCbu:      1,
		AddTripCent:     1,
		AddTripNpu:      1,
		AddTripDetVpo:   1,
	}
}

// check boxtime calc against hand calculation
func TestGetTime(t *testing.T) {
	dd := sampleDismountData()
	got := dd.GetTime()
	want := 10.2381

	if math.Abs(got-want) > .001 {
		t.Errorf("%.4f New Time does not match hand calc %.4f", got, want)
	}
}
