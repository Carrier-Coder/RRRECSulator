package vehicle

import (
	"math"
	"testing"
)

func sampleVehicleData() VehicleData {
	return VehicleData{
		Miles: 90,
	}
}

func TestGetTime(t *testing.T) {
	vd := sampleVehicleData()
	got := vd.GetTime()

	//hand calc from data above
	want := 4.69949

	if math.Abs(got-want) > 0.001 {
		t.Errorf("%v got but want %v", got, want)
	}
}
