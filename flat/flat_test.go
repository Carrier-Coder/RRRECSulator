package flat

import (
	"math"
	"testing"
)

func sampleFlatData() FlatData {
	return FlatData{
		RandomFlats:       100,
		CarrierRouteFlats: 100,
		WSSFlats:          100,
		DPSFlats:          100,

		DrivePOV: true,
	}
}

// check calc against hand calculation
func TestGetTime(t *testing.T) {
	fd := sampleFlatData()
	got := fd.GetTime()
	want := 30.52
	if math.Abs(got-want) > 0.001 {
		t.Errorf("%.4f New Time does not match hand calc %.4f", got, want)
	}
}

func TestDetectCaseDPS(t *testing.T) {
	fd := sampleFlatData()
	got := fd.caseDPS()
	want := true

	if got != want {
		t.Errorf("%v got but want %v", got, want)
	}

	fd.DriveGOVLH = true
	got = fd.caseDPS()
	want = true

	if got != want {
		t.Errorf("%v got but want %v", got, want)
	}

	fd.DrivePOV = false
	fd.DriveGOVLH = false
	got = fd.caseDPS()
	want = false

	if got != want {
		t.Errorf("%v got but want %v", got, want)
	}
}

func TestTotalFlats(t *testing.T) {
	fd := sampleFlatData()

	got := fd.TotalFlats()
	want := 400

	if got != want {
		t.Errorf("%v got but wanted %v", got, want)
	}
}

func TestCasedFlats(t *testing.T) {
	fd := sampleFlatData()

	got := fd.CasedFlats()
	want := 400

	if got != want {
		t.Errorf("%v got but wanted %v", got, want)
	}

	fd.DrivePOV = false
	fd.DriveGOVLH = false
	got = fd.CasedFlats()
	want = 300
	if got != want {
		t.Errorf("%v got but wanted %v", got, want)
	}

}

//func TestGetVerifyTime(t *testing.T) {
//fd := sampleFlatData()
//
//got := fd.GetVerifyTime()
//want := 5.68
//
//if math.Abs(got-want) > 0.001 {
//t.Errorf("%.4f New Time does not match hand calc %.4f", got, want)
//}
//}
