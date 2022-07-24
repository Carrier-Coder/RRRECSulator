package drive

import (
	"math"
	"testing"
)

func sampleDriveData() DriveData {
	return DriveData{
		Distances:     []float64{300, 250, 28, 100, 90, 2147},
		StopSigns:     5,
		YieldSigns:    5,
		TrafficLights: 5,
		Crosswalks:    5,
		RRCrossings:   5,
		AccessGates:   5,

		Creeps: 5,
	}

}

func TestFullTime(t *testing.T) {
	dd := sampleDriveData()
	got := dd.GetTime()
	want := 4.0755
	if math.Abs(got-want) > 0.001 {
		t.Errorf("%.4f old time does not match hand calc %.4f", got, want)
	}
}

func TestApplyDSM(t *testing.T) {
	dd := sampleDriveData()

	got := applyDSM(dd.Distances)
	//1.586 drive speed calced by hand using Drive Speed Matrix
	//From RRECS comprehensive guide
	want := 1.586

	if got-want > .001 {
		t.Errorf("%.4f Matrix time does not match hand calc %.4f", got, want)
	}
}

func TestStopSignTime(t *testing.T) {
	dd := sampleDriveData()

	got := dd.stopSignTime()
	want := 0.2495

	if got-want > 0.0001 {
		t.Errorf("%.4f stop sign calc does not match hand calc %.4f", got, want)
	}
}

func TestYieldSignTime(t *testing.T) {
	dd := sampleDriveData()
	got := dd.yieldSignTime()
	want := 0.0795

	if got-want > 0.0001 {
		t.Errorf("%.4f yield sign calc does not match hand calc %.4f", got, want)
	}
}

func TestTrafficLightTime(t *testing.T) {
	dd := sampleDriveData()
	got := dd.trafficLightTime()
	want := 1.02555

	if got-want > 0.0001 {
		t.Errorf("%.4f traffic light calc does not match hand calc %.4f", got, want)
	}
}

func TestCrossWalkTime(t *testing.T) {
	dd := sampleDriveData()
	got := dd.crosswalkTime()
	want := 0.046

	if got-want > 0.0001 {
		t.Errorf("%.4f crosswalk calc does not match hand calc %.4f", got, want)
	}
}

func TestrrCrossingTime(t *testing.T) {
	dd := sampleDriveData()
	got := dd.rrCrossingTime()
	want := 0.0825

	if got-want > 0.0001 {
		t.Errorf("%.4f crosswalk calc does not match hand calc %.4f", got, want)
	}
}

func TestAccessGateTime(t *testing.T) {
	dd := sampleDriveData()
	got := dd.accessGateTime()
	want := 0.9025

	if got-want > 0.0001 {
		t.Errorf("%.4f accessGate calc does not match hand calc %.4f", got, want)
	}
}

func TestCreepTime(t *testing.T) {
	dd := sampleDriveData()
	got := dd.creepTime()
	want := 0.104

	if got-want > 0.0001 {
		t.Errorf("%.4f creep time calc does not match hand calc %.4f", got, want)
	}
}
