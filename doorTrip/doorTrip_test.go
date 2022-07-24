package doorTrip

import (
	"math"
	"testing"
)

func sampleDoorData() DoorData {
	return DoorData{

		Trips:        19,
		Parcels:      20,
		Accountables: 1,
		Misc:         2,
		PickupEvents: 1,
	}
}

func TestDoorMisctime(t *testing.T) {
	dd := sampleDoorData()
	got := dd.doorMiscTime()
	want := 2 * 0.0854

	if math.Abs(got-want) > 0.001 {
		t.Errorf("%v got but want %v", got, want)
	}
}

func TestGetTime(t *testing.T) {
	dd := sampleDoorData()
	got := dd.GetTime()
	want := 20.0448

	if math.Abs(got-want) > 0.001 {
		t.Errorf("%v got but want %v", got, want)
	}
}
