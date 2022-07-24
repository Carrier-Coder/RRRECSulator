package parcel

import (
	"testing"

	"rrecsulator.com/standards"
)

func sampleParcelData() ParcelData {
	return ParcelData{
		MailBoxParcels: 10,
		LockerParcels:  10,
		DoorParcels:    10,
	}
}

func TestTotalParcels(t *testing.T) {
	pd := sampleParcelData()

	got := pd.TotalParcels()
	want := 30

	if got != want {
		t.Errorf("%v got but want %v", got, want)
	}
}

func TestLargeParcels(t *testing.T) {
	pd := sampleParcelData()
	got := pd.LargeParcels()
	want := 10

	if got != want {
		t.Errorf("%v got but want %v", got, want)
	}
}

func TestMediumParcels(t *testing.T) {
	pd := sampleParcelData()

	got := pd.MediumParcels()
	want := 16

	if got != want {
		t.Errorf("%v got but want %v", got, want)
	}
}

func TestSmallParcels(t *testing.T) {
	pd := sampleParcelData()

	got := pd.SmallParcels()
	want := 4

	if got != want {
		t.Errorf("%v got but want %v", got, want)
	}
}

func TestCarts(t *testing.T) {
	pd := sampleParcelData()
	got := pd.howManyCarts()
	want := 1

	// check first cart
	if got != want {
		t.Errorf("%v got but want %v", got, want)
	}

	//check multiple carts
	pd = sampleParcelData()
	pd.DoorParcels = 130
	got = pd.howManyCarts()
	want = 3

	if got != want {
		t.Errorf("%v got but want %v", got, want)
	}
}

func TestCartTime(t *testing.T) {
	pd := sampleParcelData()
	//check first cart time
	got := pd.GetCartTime()
	want := float64(standards.GATHER_1ST_CART_TIME)
	if got != want {
		t.Errorf("%v got but want %v", got, want)
	}

	pd = sampleParcelData()
	pd.DoorParcels = 130
	got = pd.GetCartTime()

	// 3 trips total
	want = float64(standards.GATHER_1ST_CART_TIME)
	want += 2 * float64(standards.GATHER_ADD_CART_TIME)

	if got != want {
		t.Errorf("%v got but want %v", got, want)
	}
}

func TestGetTime(t *testing.T) {
	pd := sampleParcelData()
	got := pd.GetTime()
	want := 17.2824
	if got != want {
		t.Errorf("%v got but want %v", got, want)
	}
}
