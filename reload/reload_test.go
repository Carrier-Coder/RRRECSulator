package reload

import (
	"math"
	"testing"

	"rrecsulator.com/standards"
)

func sampleReloadData() ReloadData {
	return ReloadData{

		RandomLetters:    150,
		WSSLetters:       500,
		BoxholderLetters: 500,

		//cased flats
		RandomFlats:    40,
		CRFlats:        150,
		WSSFlats:       0,
		BoxholderFlats: 500,

		//DPS
		DPSLetters: 2100,
		DPSFlats:   0,

		//parcels
		MailBoxParcels: 79,
		LockerParcels:  5,
		DoorParcels:    35,

		Bundle: standards.TWO_BUNDLE,
	}
}

func TestGetMailTrayTime(t *testing.T) {
	rd := sampleReloadData()
	got := rd.GetMailTrayTime()
	// calculated in guide
	want := 11.5751
	if math.Abs(got-want) > 0.001 {
		t.Errorf("%v got but want %v", got, want)

	}
}

func TestGetMediumParcelTime(t *testing.T) {
	rd := sampleReloadData()
	got := rd.GetMediumParcelTime()
	// calculated in guide
	want := 3.078

	if math.Abs(got-want) > 0.001 {
		t.Errorf("%v got but want %v", got, want)

	}
}

func TestGetMediumParcelTimeBiggerTable(t *testing.T) {

	// should be more trays than in the table
	rd := ReloadData{
		MailBoxParcels: 210,
	}
	got := rd.GetMediumParcelTime()

	//210 parcels should produce 17 medium trays (only paid for 15)
	want := 4.7264 + 4.2984 + 1.2172
	if math.Abs(got-want) > 0.001 {
		t.Errorf("%v got but want %v", got, want)
	}

}

func TestGetLargeParcelTime(t *testing.T) {
	rd := sampleReloadData()
	got := rd.GetLargeParcelTime()
	// calculated in guide
	want := 5.4955

	if math.Abs(got-want) > 0.001 {
		t.Errorf("%v got but want %v", got, want)

	}
}

func TestGetTime(t *testing.T) {
	rd := sampleReloadData()
	got := rd.GetTime()
	//times from tests above
	want := 11.5751 + 3.078 + 5.4955

	if math.Abs(got-want) > 0.001 {
		t.Errorf("%v got but want %v", got, want)

	}
}

func TestZeroes(t *testing.T) {
	rd := ReloadData{}
	want := 0.0

	got := rd.GetLargeParcelTime()
	if math.Abs(got-want) > 0.001 {
		t.Errorf("%v got but want %v", got, want)
	}

	got = rd.GetMediumParcelTime()
	if math.Abs(got-want) > 0.001 {
		t.Errorf("%v got but want %v", got, want)
	}

	got = rd.GetMailTrayTime()
	if math.Abs(got-want) > 0.001 {
		t.Errorf("%v got but want %v", got, want)
	}
}

func TestMax(t *testing.T) {

	// should be more trays than in the table
	rd := ReloadData{
		RandomFlats: 100000,
		Bundle:      standards.TWO_BUNDLE,
	}
	got := rd.GetMailTrayTime()
	lenTable := len(standards.MAIL_TRAY_TABLE)
	want := standards.MAIL_TRAY_TABLE[lenTable-1][rd.Bundle]
	if math.Abs(got-want) > 0.001 {
		t.Errorf("%v got but want %v", got, want)
	}

}
