package cbu

import (
	"testing"

	"rrecsulator.com/standards"
)

func sampleCbuData() CBUData {
	return CBUData{
		CBUs:     10,
		CBUBoxes: 100,

		CENTs:     10,
		CENTBoxes: 100,

		NPUs:     10,
		NPUBoxes: 100,

		DETs:     10,
		DETBoxes: 100,

		CollectionBoxes: 10,

		Bundle: standards.TWO_BUNDLE,
	}
}

func TestCollectionTime(t *testing.T) {
	cubd := sampleCbuData()
	got := cubd.collectionTime()
	want := 0.558
	if got-want > 0.001 {
		t.Errorf("%.4f calculated want %.4f hand calc", got, want)
	}
}

func TestCbuSkip(t *testing.T) {
	cbud := sampleCbuData()
	cbud.CBUBoxesSkipped = 200
	got := cbud.cbuTime()
	want := 3.331
	if got-want > 0.001 {
		t.Errorf("%.4f calculated want %.4f hand calc", got, want)
	}

	cbud.CBUBoxesSkipped = 50
	got = cbud.cbuTime()
	want = 8.696
	if got-want > 0.001 {
		t.Errorf("%.4f calculated want %.4f hand calc", got, want)
	}
}

func TestCbuTime(t *testing.T) {
	cbud := sampleCbuData()
	got := cbud.cbuTime()
	want := 14.081

	if got-want > 0.001 {
		t.Errorf("%.4f calculated want %.4f hand calc", got, want)
	}
}

func TestCentSkip(t *testing.T) {
	cbud := sampleCbuData()
	cbud.CENTBoxesSkipped = 200
	got := cbud.centTime()
	want := 3.803
	if got-want > 0.001 {
		t.Errorf("%.4f calculated want %.4f hand calc", got, want)
	}

	cbud.CENTBoxesSkipped = 50
	got = cbud.centTime()
	want = 9.188
	if got-want > 0.001 {
		t.Errorf("%.4f calculated want %.4f hand calc", got, want)
	}
}

func TestCentTime(t *testing.T) {
	cbud := sampleCbuData()
	got := cbud.centTime()
	want := 14.573

	if got-want > 0.001 {
		t.Errorf("%.4f calculated want %.4f hand calc", got, want)
	}
}

func TestNpuSkip(t *testing.T) {
	cbud := sampleCbuData()
	cbud.NPUBoxesSkipped = 200
	got := cbud.npuTime()
	want := .983
	if got-want > 0.001 {
		t.Errorf("%.4f calculated want %.4f hand calc", got, want)
	}

	cbud.NPUBoxesSkipped = 50
	got = cbud.npuTime()
	want = 6.368
	if got-want > 0.001 {
		t.Errorf("%.4f calculated want %.4f hand calc", got, want)
	}
}
func TestNpuTime(t *testing.T) {
	cbud := sampleCbuData()
	got := cbud.npuTime()
	want := 11.753

	if got-want > 0.001 {
		t.Errorf("%.4f calculated want %.4f hand calc", got, want)
	}
}

func TestDetSkip(t *testing.T) {
	cbud := sampleCbuData()
	cbud.DETBoxesSkipped = 200
	got := cbud.detTime()
	want := 2.239
	if got-want > 0.001 {
		t.Errorf("%.4f calculated want %.4f hand calc", got, want)
	}

	cbud.DETBoxesSkipped = 50
	got = cbud.detTime()
	want = 12.659
	if got-want > 0.001 {
		t.Errorf("%.4f calculated want %.4f hand calc", got, want)
	}
}
func TestDetTime(t *testing.T) {
	cbud := sampleCbuData()
	got := cbud.detTime()
	want := 23.079

	if got-want > 0.001 {
		t.Errorf("%.4f calculated want %.4f hand calc", got, want)
	}
}

// check GetTime calc against hand calculation
func TestGetTime(t *testing.T) {
	cbud := sampleCbuData()
	got := cbud.GetTime()
	want := 64.044

	if got-want > 0.001 {
		t.Errorf("%.4f calculated want %.4f hand calc", got, want)
	}

}

// make sure sane.  OneBundle < twoBundle < threeBundle
func TestSanity(t *testing.T) {
	cbud := sampleCbuData()

	oneBundleTime := cbud.GetTime()

	cbud.Bundle = standards.TWO_BUNDLE
	twoBundleTime := cbud.GetTime()

	cbud.Bundle = standards.THREE_BUNDLE
	threeBundleTime := cbud.GetTime()

	if oneBundleTime > twoBundleTime {
		t.Errorf("%.4f one bundle time greater than %.4f two Bundle Time",
			oneBundleTime, twoBundleTime)
	}
	if twoBundleTime > threeBundleTime {
		t.Errorf("%.4f two bundle time greater than %.4f three Bundle Time",
			twoBundleTime, threeBundleTime)
	}
}
