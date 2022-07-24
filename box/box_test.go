package box

import (
	"math"
	"testing"

	"rrecsulator.com/standards"
)

func sampleBoxData() BoxData {
	return BoxData{
		RegCurbBox:  10,
		RegSdwkBox:  10,
		RegOtherBox: 10,

		TotalFlats:   100,
		TotalLetters: 100,
		Bundle:       standards.ONE_BUNDLE,
	}
}

// check boxtime calc against hand calculation
func TestGetTime(t *testing.T) {
	bd := sampleBoxData()
	got := bd.GetTime()
	want := 6.765

	if math.Abs(got-want) > .001 {
		t.Errorf("%.4f New Time does not match hand calc %.4f", got, want)
	}
}

// make sure sane.  OneBundle < twoBundle < threeBundle
func TestSanity(t *testing.T) {
	bd := sampleBoxData()

	oneBundleTime := bd.GetTime()

	bd.Bundle = standards.TWO_BUNDLE
	twoBundleTime := bd.GetTime()

	bd.Bundle = standards.THREE_BUNDLE
	threeBundleTime := bd.GetTime()

	if oneBundleTime > twoBundleTime {
		t.Errorf("%.4f one bundle time greater than %.4f two Bundle Time",
			oneBundleTime, twoBundleTime)
	}
	if twoBundleTime > threeBundleTime {
		t.Errorf("%.4f two bundle time greater than %.4f three Bundle Time",
			twoBundleTime, threeBundleTime)
	}
}

func TestCurbBoxesSkipped(t *testing.T) {
	bd := BoxData{
		RegCurbBox:          5,
		RegCurbBoxesSkipped: 4,
		Bundle:              standards.THREE_BUNDLE,
	}
	got := bd.GetTime()
	want := 0.2572
	if math.Abs(got-want) > .001 {
		t.Errorf("%.4f New Time does not match hand calc %.4f", got, want)
	}
}

func TestSidewalkBoxesSkipped(t *testing.T) {
	bd := BoxData{
		RegSdwkBox:          5,
		RegSdwkBoxesSkipped: 4,
		Bundle:              standards.THREE_BUNDLE,
	}
	got := bd.GetTime()
	want := 0.2445
	if math.Abs(got-want) > .001 {
		t.Errorf("%.4f New Time does not match hand calc %.4f", got, want)
	}
}

func TestOtherBoxesSkipped(t *testing.T) {
	bd := BoxData{
		RegOtherBox:          5,
		RegOtherBoxesSkipped: 4,
		Bundle:               standards.THREE_BUNDLE,
	}
	got := bd.GetTime()
	want := 0.2434
	if math.Abs(got-want) > .001 {
		t.Errorf("%.4f New Time does not match hand calc %.4f", got, want)
	}
}
