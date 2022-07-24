package boxholder

import (
	"math"
	"testing"
)

func sampleBoxholder() BoxholderData {
	return BoxholderData{
		BoxholderFlats:   100,
		BoxholderLetters: 100,
	}
}

// check calc against hand calculation
func TestCalc(t *testing.T) {
	bhd := sampleBoxholder()
	got := bhd.GetTime()
	want := 9.49
	if math.Abs(got-want) > 0.001 {
		t.Errorf("%.4f New Time does not match hand calc %.4f", got, want)
	}
}

func TestTotalBoxholders(t *testing.T) {
	fd := sampleBoxholder()

	got := fd.TotalBoxholders()
	want := 200

	if got != want {
		t.Errorf("%v got but wanted %v", got, want)
	}
}
